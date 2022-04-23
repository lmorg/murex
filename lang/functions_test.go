package lang_test

import (
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

type testFuncParseDataTypesT struct {
	Parameters string
	Error      bool
	Expected   []lang.MxFunctionParams
}

func testFuncParseDataTypes(t *testing.T, tests []testFuncParseDataTypesT) {
	t.Helper()
	count.Tests(t, len(tests))

	for i := range tests {
		actual, err := lang.ParseMxFunctionParameters(tests[i].Parameters)
		if (err == nil) == tests[i].Error {
			t.Errorf("Unexpected error raised in test %d", i)
			t.Logf("Parameters: %s", tests[i].Parameters)
			t.Logf("Expected:   %s", json.LazyLogging(tests[i].Expected))
			t.Logf("Actual:     %s", json.LazyLogging(actual))
			t.Logf("exp err:    %v", tests[i].Error)
			t.Logf("act err:    %s", err)
		}

		if json.LazyLogging(tests[i].Expected) != json.LazyLogging(actual) {
			t.Errorf("Unexpected error raised in test %d", i)
			t.Logf("Parameters: %s", tests[i].Parameters)
			t.Logf("Expected:   %s", json.LazyLogging(tests[i].Expected))
			t.Logf("Actual:     %s", json.LazyLogging(actual))
			t.Logf("exp err:    %v", tests[i].Error)
			t.Logf("act err:    %s", err)
		}
	}
}

func TestFuncParseDataTypes(t *testing.T) {
	tests := []testFuncParseDataTypesT{
		{
			Parameters: `name: str, age: int`,
			Expected: []lang.MxFunctionParams{{
				Name:     "name",
				DataType: "str",
			}, {
				Name:     "age",
				DataType: "int",
			}},
		},
		{
			Parameters: `name: str "What is your name?", age: int "How old are you?"`,
			Expected: []lang.MxFunctionParams{{
				Name:        "name",
				DataType:    "str",
				Description: "What is your name?",
			}, {
				Name:        "age",
				DataType:    "int",
				Description: "How old are you?",
			}},
		},
		{
			Parameters: `name: str [Bob], age: int [100]`,
			Expected: []lang.MxFunctionParams{{
				Name:     "name",
				DataType: "str",
				Default:  "Bob",
			}, {
				Name:     "age",
				DataType: "int",
				Default:  "100",
			}},
		},
		{
			Parameters: `name: str "What is your name?" [Bob], age: int "How old are you?" [100]`,
			Expected: []lang.MxFunctionParams{{
				Name:        "name",
				DataType:    "str",
				Description: "What is your name?",
				Default:     "Bob",
			}, {
				Name:        "age",
				DataType:    "int",
				Description: "How old are you?",
				Default:     "100",
			}},
		},
		{
			Parameters: `name: str [Bob]`,
			Expected: []lang.MxFunctionParams{{
				Name:     "name",
				DataType: "str",
				Default:  "Bob",
			}},
		},
	}

	testFuncParseDataTypes(t, tests)
}

func FuzzFuncParseDataTypes(f *testing.F) {
	tests := []string{"name: str, age: int", "", "!12345"}
	for _, tc := range tests {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		lang.ParseMxFunctionParameters(orig)
		// we are just testing we can't cause an unhandled panic
	})
}
