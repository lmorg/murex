package typemgmt

import (
	"os"
	"testing"

	_ "github.com/lmorg/murex/builtins/core/io"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
)

type Test struct {
	Block    string
	Name     string
	Value    string
	DataType string
	Fail     bool
}

const envVarPrefix = "MUREX_TEST_VAR_"

func VariableTests(tests []Test, t *testing.T) {
	// these tests don't support multiple counts
	if os.Getenv(envVarPrefix+t.Name()) == "1" {
		return
	}

	err := os.Setenv(envVarPrefix+t.Name(), "1")
	if err != nil {
		t.Fatalf("Aborting test because unable to set env: %s", err)
	}

	count.Tests(t, len(tests)*2)

	for i := range tests {
		fork := lang.ShellProcess.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_CREATE_STDERR)
		fork.Name.Set("VariableTests()")
		_, err := fork.Execute([]rune(tests[i].Block))
		if err != nil {
			t.Error(err.Error())
		}

		b, err := fork.Stderr.ReadAll()
		if err != nil {
			t.Error("unable to read from stderr: " + err.Error())
		}

		value, err := fork.Variables.GetString(tests[i].Name)
		dataType := fork.Variables.GetDataType(tests[i].Name)

		if err != nil {
			t.Errorf("Test %d failed on %s:", i, t.Name())
			t.Logf("  code block:    %s", tests[i].Block)
			t.Logf("  variable name: %s", tests[i].Name)
			t.Logf("  error value:   %s", err.Error())
			continue
		}

		if value != tests[i].Value ||
			dataType != tests[i].DataType ||
			(len(b) > 0) != tests[i].Fail {

			t.Errorf("Test %d failed on %s:", i, t.Name())
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

func UnSetTests(unsetter string, tests []string, t *testing.T) {
	// these tests don't support multiple counts
	if os.Getenv(envVarPrefix+t.Name()) == "1" {
		return
	}

	err := os.Setenv(envVarPrefix+t.Name(), "1")
	if err != nil {
		t.Fatalf("Aborting test because unable to set env: %s", err)
	}

	count.Tests(t, len(tests)*2)

	for i := range tests {
		fork := lang.ShellProcess.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_CREATE_STDERR)
		fork.Name.Set("UnSetTests()")

		old, err := fork.Variables.GetString(tests[i])
		if err != nil {
			t.Errorf("Test %d failed:", i)
			t.Logf("  variable name: %s", tests[i])
			t.Logf("  error:         %s", err.Error())
			continue
		}

		block := unsetter + ": " + tests[i]
		_, err = fork.Execute([]rune(block))
		if err != nil {
			t.Error(err.Error())
		}

		b, err := fork.Stderr.ReadAll()
		if err != nil {
			t.Error("unable to read from stderr: " + err.Error())
		}

		value, err := fork.Variables.GetString(tests[i])
		dataType := fork.Variables.GetDataType(tests[i])

		if err != nil {
			t.Errorf("Test %d failed:", i)
			t.Logf("  unsetter block: %s", block)
			t.Logf("  variable name:  %s", tests[i])
			t.Logf("  variable value: %s", value)
			t.Logf("  variable type:  %s", dataType)
			t.Logf("  error:          %s", err.Error())
			continue
		}

		if value == old || dataType != "" || len(b) > 0 {
			t.Errorf("Test %d failed:", i)
			t.Logf("  unsetter block: %s", block)
			t.Logf("  variable name:  %s", tests[i])
			t.Logf("  variable value: %s", value)
			t.Logf("  variable type:  %s", dataType)
			t.Logf("  stderr output:  %s", b)
		}
	}
}
