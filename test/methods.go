package test

import (
	"strings"
	"testing"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang"
)

// RunMethodTest is a template function for testing builtins run as methods.RunMethodTest.
// dataType should preloaded in calling functions, eg
//
//     _ "github.com/lmorg/murex/builtins/types/generic"
//	   _ "github.com/lmorg/murex/builtins/types/json"
func RunMethodTest(t *testing.T, cmd func(*lang.Process) error, methodName string, input string, dataType string, params []string, output string, expectedError error) {
	p := lang.NewTestProcess()
	p.IsMethod = true
	p.Parameters.Params = params

	p.Stdin = streams.NewStdin()
	p.Stdin.SetDataType(dataType)
	p.Stdin.Write([]byte(input))

	p.Stdout = streams.NewStdin()

	err := cmd(p)
	if err != nil && err != expectedError {
		t.Error(err)
		return
	}

	b, err := p.Stdout.ReadAll()
	if err != nil {
		t.Error("Error reading from STDOUT")
		t.Log(err)
		//return
	}

	if string(b) != output || expectedError != nil {
		t.Errorf("Unexpected `%s` return", methodName)
		t.Log("  input:          ", strings.Replace(input, "\n", `\n`, -1))
		t.Log("  input data type:", dataType)
		t.Log("  parameters      ", params)
		t.Log("  expected output:", strings.Replace(output, "\n", `\n`, -1))
		t.Log("  actual output:  ", strings.Replace(string(b), "\n", `\n`, -1))
		t.Log("  expected error: ", expectedError)
		t.Log("  actual error:   ", err)
		t.Log("  eo bytes:       ", []byte(output))
		t.Log("  ao bytes:       ", b)
	}
}
