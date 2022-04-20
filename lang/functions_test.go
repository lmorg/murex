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

	for i, test := range tests {
		actual, err := lang.ParseMxFunctionParameters(test.Parameters)
		if (err != nil) == test.Error {
			t.Errorf("Unexpected error raised in test %d", i)
			t.Logf("Parameters: %s", test.Parameters)
			t.Logf("Expected:   %s", json.LazyLogging(test.Expected))
			t.Logf("Actual:     %s", json.LazyLogging(actual))
			t.Logf("exp err:    %v", test.Error)
			t.Logf("act err:    %s", err)
		}

		if json.LazyLogging(test.Expected) == json.LazyLogging(actual) {
			t.Errorf("Unexpected error raised in test %d", i)
			t.Logf("Parameters: %s", test.Parameters)
			t.Logf("Expected:   %s", json.LazyLogging(test.Expected))
			t.Logf("Actual:     %s", json.LazyLogging(actual))
			t.Logf("exp err:    %v", test.Error)
			t.Logf("act err:    %s", err)
		}
	}
}

func TestFuncParseDataTypes(t *testing.T) {
	tests := []testFuncParseDataTypesT{
		{
			Parameters: `name: str, age: int`,
			Error:      false,
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
			Error:      false,
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
	}

	testFuncParseDataTypes(t, tests)
}
