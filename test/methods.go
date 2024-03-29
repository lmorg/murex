package test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
)

func errString(err error) string {
	if err != nil {
		return err.Error()
	}

	return ""
}

// RunMethodTest is a template function for testing builtins run as methods.
// dataType should be preloaded in calling functions, eg
//
//	    _ "github.com/lmorg/murex/builtins/types/generic"
//		   _ "github.com/lmorg/murex/builtins/types/json"
func RunMethodTest(t *testing.T, cmd func(*lang.Process) error, methodName string, input string, dataType string, params []string, output string, expectedError error) {
	t.Helper()
	count.Tests(t, 1)

	p := lang.NewTestProcess()
	p.IsMethod = true
	p.Parameters.DefineParsed(params)

	p.Stdin = streams.NewStdin()
	p.Stdin.SetDataType(dataType)
	p.Stdin.Write([]byte(input))

	p.Stdout = streams.NewStdin()

	err := cmd(p)

	b, readErr := p.Stdout.ReadAll()
	if readErr != nil {
		t.Error("Error reading from STDOUT")
		t.Log(readErr)
		return
	}

	if string(b) != output || errString(expectedError) != errString(err) {
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

// RunMethodRegexTest is a template function for using regex to test builtins
// run as methods.
// dataType should be preloaded in calling functions, eg
//
//	    _ "github.com/lmorg/murex/builtins/types/generic"
//		   _ "github.com/lmorg/murex/builtins/types/json"
func RunMethodRegexTest(t *testing.T, cmd func(*lang.Process) error, methodName string, input string, dataType string, params []string, outputMatchRx string) {
	t.Helper()
	count.Tests(t, 1)

	p := lang.NewTestProcess()
	p.IsMethod = true
	p.Parameters.DefineParsed(params)

	p.Stdin = streams.NewStdin()
	p.Stdin.SetDataType(dataType)
	p.Stdin.Write([]byte(input))

	p.Stdout = streams.NewStdin()

	rx, err := regexp.Compile(outputMatchRx)
	if err != nil {
		t.Error(err)
		return
	}

	err = cmd(p)
	if err != nil {
		t.Error(err)
		return
	}

	b, err := p.Stdout.ReadAll()
	if err != nil {
		t.Error("Error reading from STDOUT")
		t.Log(err)
		return
	}

	if !rx.MatchString(string(b)) {
		t.Errorf("Unexpected `%s` return", methodName)
		t.Log("  input:          ", strings.Replace(input, "\n", `\n`, -1))
		t.Log("  input data type:", dataType)
		t.Log("  parameters      ", params)
		t.Log("  expected match: ", outputMatchRx)
		t.Log("  actual output:  ", strings.Replace(string(b), "\n", `\n`, -1))
		t.Log("  actual error:   ", err)
	}
}
