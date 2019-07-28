package typemgmt

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/io"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
)

type Test struct {
	Block    string
	Name     string
	Value    string
	DataType string
	Fail     bool
}

func VariableTests(tests []Test, t *testing.T) {
	defaults.Defaults(lang.InitConf, false)
	lang.InitEnv()

	for i := range tests {
		fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_CREATE_STDERR)
		fork.Name = "VariableTests()"
		_, err := fork.Execute([]rune(tests[i].Block))
		if err != nil {
			t.Error(err.Error())
		}

		b, err := fork.Stderr.ReadAll()
		if err != nil {
			t.Error("unable to read from stderr: " + err.Error())
		}

		value := fork.Variables.GetString(tests[i].Name)
		dataType := fork.Variables.GetDataType(tests[i].Name)

		if value != tests[i].Value ||
			dataType != tests[i].DataType ||
			(len(b) > 0) != tests[i].Fail {

			t.Errorf("Test %d failed:", i)
			t.Logf("  code block:     %s", tests[i].Block)
			t.Logf("  variable name:  %s", tests[i].Name)
			t.Logf("  expected value: %s", tests[i].Value)
			t.Logf("  actual value:   %s", value)
			t.Log("  expected bytes: ", []byte(tests[i].Value))
			t.Log("  actual bytes:   ", []byte(value))
			t.Logf("  expected type:  %s", tests[i].DataType)
			t.Logf("  actual type:    %s", dataType)
			t.Logf("  stderr output:  %s", b)
			t.Logf("  error expected: %t", tests[i].Fail)
		}
	}
}
