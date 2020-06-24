package typemgmt

import (
	"encoding/json"
	"os"
	"testing"

	_ "github.com/lmorg/murex/builtins/core/io"
	"github.com/lmorg/murex/config/defaults"
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

//var varTestMutex sync.Mutex

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

	defaults.Defaults(lang.InitConf, false)
	lang.InitEnv()

	for i := range tests {
		fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_NEW_VARTABLE | lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_CREATE_STDERR)
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

	defaults.Defaults(lang.InitConf, false)
	lang.InitEnv()

	for i := range tests {
		fork := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_CREATE_STDERR)
		fork.Name = "UnSetTests()"
		block := unsetter + ": " + tests[i]
		_, err := fork.Execute([]rune(block))
		if err != nil {
			t.Error(err.Error())
		}

		b, err := fork.Stderr.ReadAll()
		if err != nil {
			t.Error("unable to read from stderr: " + err.Error())
		}

		value := fork.Variables.GetString(tests[i])
		dataType := fork.Variables.GetDataType(tests[i])

		if value != "" || dataType != "" || len(b) > 0 {
			t.Errorf("Test %d failed:", i)
			t.Logf("  unsetter block: %s", block)
			t.Logf("  variable name:  %s", tests[i])
			t.Logf("  variable value: %s", value)
			t.Logf("  variable type:  %s", dataType)
			t.Logf("  stderr output:  %s", b)
		}
	}
}

////

func TestSplitVarString(t *testing.T) {
	splitTest := []struct {
		Input    []string
		DataType string
		Name     string
		Value    string
		Error    bool
	}{
		{
			Input:    []string{"foo=bar"},
			DataType: "",
			Name:     "foo",
			Value:    "bar",
		},
		{
			Input: []string{"-=bar"},
			Error: true,
		},
		{
			Input: []string{"-", "=bar"},
			Error: true,
		},
		{
			Input: []string{"=bar"},
			Error: true,
		},
		{
			Input: []string{"", "=bar"},
			Error: true,
		},
		{
			Input:    []string{"foo =bar"},
			DataType: "",
			Name:     "foo",
			Value:    "bar",
		},
		{
			Input:    []string{"foo= bar"},
			DataType: "",
			Name:     "foo",
			Value:    " bar",
		},
		{
			Input:    []string{"foo = bar"},
			DataType: "",
			Name:     "foo",
			Value:    " bar",
		},
		{
			Input:    []string{"foo  =bar"},
			DataType: "",
			Name:     "foo",
			Value:    "bar",
		},
		{
			Input:    []string{"foo=  bar"},
			DataType: "",
			Name:     "foo",
			Value:    "  bar",
		},
		{
			Input:    []string{"foo  =  bar"},
			DataType: "",
			Name:     "foo",
			Value:    "  bar",
		},
		/////
		{
			Input: []string{" foo=bar"},
			Error: true,
		},
		{
			Input: []string{" foo =bar"},
			Error: true,
		},
		{
			Input: []string{" foo= bar"},
			Error: true,
		},
		{
			Input: []string{" foo = bar"},
			Error: true,
		},
		{
			Input:    []string{"foo  =bar "},
			DataType: "",
			Name:     "foo",
			Value:    "bar ",
		},
		{
			Input:    []string{"foo=  bar "},
			DataType: "",
			Name:     "foo",
			Value:    "  bar ",
		},
		{
			Input:    []string{"foo  =  bar "},
			DataType: "",
			Name:     "foo",
			Value:    "  bar ",
		},

		/////
		{
			Input: []string{"dt foo=bar"},
			Error: true,
		},
		{
			Input: []string{"dt foo bar"},
			Error: true,
		},
		{
			Input: []string{"dt foo bar=value"},
			Error: true,
		},
		{
			Input: []string{"dt foo bar =value"},
			Error: true,
		},
		{
			Input: []string{"dt foo bar= value"},
			Error: true,
		},
		{
			Input: []string{"dt foo bar = value"},
			Error: true,
		},
		{
			Input: []string{"d t foo bar=value"},
			Error: true,
		},
		{
			Input: []string{"d t foo bar =value"},
			Error: true,
		},
		{
			Input: []string{"d t foo bar= value"},
			Error: true,
		},
		{
			Input: []string{"d t foo bar = value"},
			Error: true,
		},
		{
			Input:    []string{"dt", "foo=bar"},
			DataType: "dt",
			Name:     "foo",
			Value:    "bar",
		},
		{
			Input: []string{"dt", "foo", "bar"},
			Error: true,
		},
		{
			Input: []string{"dt", "foo", "bar=value"},
			Error: true,
		},
		{
			Input: []string{"dt", "foo", "bar =value"},
			Error: true,
		},
		{
			Input: []string{"dt", "foo", "bar= value"},
			Error: true,
		},
		{
			Input: []string{"dt", "foo", "bar = value"},
			Error: true,
		},
		{
			Input: []string{"d", "t", "foo", "bar=value"},
			Error: true,
		},
		{
			Input: []string{"d", "t", "foo", "bar =value"},
			Error: true,
		},
		{
			Input: []string{"d", "t", "foo", "bar= value"},
			Error: true,
		},
		{
			Input: []string{"d", "t", "foo", "bar = value"},
			Error: true,
		},
		/////
		{
			Input:    []string{"foo", "=", "bar"},
			DataType: "",
			Name:     "foo",
			Value:    "bar",
		},
		{
			Input:    []string{"foo=", "bar"},
			DataType: "",
			Name:     "foo",
			Value:    "bar",
		},
		{
			Input:    []string{"foo", "=bar"},
			DataType: "",
			Name:     "foo",
			Value:    "bar",
		},
		{
			Input:    []string{"dt", "foo", "=", "bar"},
			DataType: "dt",
			Name:     "foo",
			Value:    "bar",
		},
		{
			Input:    []string{"dt", "foo=", "bar"},
			DataType: "dt",
			Name:     "foo",
			Value:    "bar",
		},
		{
			Input:    []string{"dt", "foo", "=bar"},
			DataType: "dt",
			Name:     "foo",
			Value:    "bar",
		},
		{
			Input:    []string{"dt", "test", "=", "foo", "bar"},
			DataType: "dt",
			Name:     "test",
			Value:    "foo bar",
		},
		{
			Input:    []string{"dt", "test", "=", "foo ", "bar"},
			DataType: "dt",
			Name:     "test",
			Value:    "foo  bar",
		},
		{
			Input:    []string{"dt", "test", "=", "foo", "bar "},
			DataType: "dt",
			Name:     "test",
			Value:    "foo bar ",
		},
		{
			Input:    []string{"dt", "test", "=", "foo=bar"},
			DataType: "dt",
			Name:     "test",
			Value:    "foo=bar",
		},
		{
			Input:    []string{"dt", "test", "=", "foo=", "bar"},
			DataType: "dt",
			Name:     "test",
			Value:    "foo= bar",
		},
		{
			Input:    []string{"dt", "test", "=", "foo", "=bar"},
			DataType: "dt",
			Name:     "test",
			Value:    "foo =bar",
		},
		{
			Input:    []string{"dt", "test", "=", "foo", "=", "bar"},
			DataType: "dt",
			Name:     "test",
			Value:    "foo = bar",
		},
	}

	for i := range splitTest {
		name, value, dataType, err := splitVarString(splitTest[i].Input)

		switch {
		case (err != nil) != splitTest[i].Error:
			t.Error("Variable block incorrectly parsed: error condition doesn't match expected")
		case name != splitTest[i].Name && !splitTest[i].Error:
			t.Error("Variable block incorrectly parsed: name doesn't match expected")
		case value != splitTest[i].Value && !splitTest[i].Error:
			t.Error("Variable block incorrectly parsed: value doesn't match expected")
		case dataType != splitTest[i].DataType && !splitTest[i].Error:
			t.Error("Variable block incorrectly parsed: data type doesn't match expected")
		default:
			continue
		}

		b, jErr := json.Marshal(splitTest[i].Input)
		if jErr != nil {
			panic(jErr)
		}

		t.Logf("  Test #:    %d", i)
		t.Logf("  Input:     %s", string(b))
		t.Logf("  Exp Name: '%s'", splitTest[i].Name)
		t.Logf("  Act Name: '%s'", name)
		t.Logf("  Exp Val:  '%s'", splitTest[i].Value)
		t.Logf("  Act Val:  '%s'", value)
		t.Logf("  Exp DT:   '%s'", splitTest[i].DataType)
		t.Logf("  Act DT:   '%s'", dataType)
		t.Logf("  Exp Err:   %v (boolean state)", splitTest[i].Error)
		t.Logf("  Act Err:   %s (error message)", err)
	}
}
