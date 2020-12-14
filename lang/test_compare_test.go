package lang

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
)

type testIsArrayMap struct {
	Input    string
	DataType string
	Expected TestStatus
}

func TestIsArray(t *testing.T) {
	tests := []testIsArrayMap{
		{
			Input:    `a\nb\nc\n`,
			DataType: types.Generic,
			Expected: TestFailed,
		},
		{
			Input:    `a\nb\nc\n`,
			DataType: types.String,
			Expected: TestPassed,
		},
		{
			Input:    `["a","b","c"]`,
			DataType: types.Json,
			Expected: TestPassed,
		},
		{
			Input:    `{"a": "a", "b": "b", "c": "c"}`,
			DataType: types.Json,
			Expected: TestFailed,
		},
		{
			Input:    `- a\n- b\n- c\n`,
			DataType: "yaml",
			Expected: TestPassed,
		},
	}

	testIsArrayTest(t, tests)
}

func testIsArrayTest(t *testing.T, testArray []testIsArrayMap) {
	InitEnv()
	count.Tests(t, len(testArray))
	t.Helper()

	for i := range testArray {
		status, msg := testIsArray([]byte(testArray[i].Input), testArray[i].DataType, "TestIsArray")
		if status != testArray[i].Expected {
			t.Error("testIsArray() returned unexpected result:")
			t.Logf("  Test Num: %d", i)
			t.Logf("  Input:    %s", string(testArray[i].Input))
			t.Logf("  DataType: %s", testArray[i].DataType)
			t.Logf("  Expected: %s", testArray[i].Expected)
			t.Logf("  Actual:   %s", status)
			t.Logf("  Error:    %s", msg)
		}
	}
}

func TestIsMap(t *testing.T) {
	tests := []testIsArrayMap{
		{
			Input:    `a\nb\nc\n`,
			DataType: types.Generic,
			Expected: TestFailed,
		},
		{
			Input:    `a\nb\nc\n`,
			DataType: types.String,
			Expected: TestFailed,
		},
		{
			Input:    `["a","b","c"]`,
			DataType: types.Json,
			Expected: TestFailed,
		},
		{
			Input:    `{"a": "a", "b": "b", "c": "c"}`,
			DataType: types.Json,
			Expected: TestPassed,
		},
		{
			Input:    `- a\n- b\n- c\n`,
			DataType: "yaml",
			Expected: TestFailed,
		},
	}

	testIsMapTest(t, tests)
}

func testIsMapTest(t *testing.T, testArray []testIsArrayMap) {
	InitEnv()
	count.Tests(t, len(testArray))
	t.Helper()

	for i := range testArray {
		status, msg := testIsMap([]byte(testArray[i].Input), testArray[i].DataType, "TestIsMap")
		if status != testArray[i].Expected {
			t.Error("testIsMap() returned unexpected result:")
			t.Logf("  Test Num: %d", i)
			t.Logf("  Input:    %s", string(testArray[i].Input))
			t.Logf("  DataType: %s", testArray[i].DataType)
			t.Logf("  Expected: %s", testArray[i].Expected)
			t.Logf("  Actual:   %s", status)
			t.Logf("  Error:    %s", msg)
		}
	}
}
