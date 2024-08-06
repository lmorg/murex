package alter_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/alter"
)

type plan struct {
	original string
	path     string
	change   string
	expected string
	error    bool
}

func alterTest(t *testing.T, test *plan) {
	t.Helper()

	count.Tests(t, 1)

	pathS, err := alter.SplitPath(test.path)
	if err != nil {
		panic(err)
	}

	var expV interface{}
	if err := json.Unmarshal([]byte(test.expected), &expV); err != nil {
		panic(err)
	}
	b, err := json.Marshal(expV)
	if err != nil {
		panic(err)
	}
	test.expected = string(b)

	var old interface{}
	err = json.Unmarshal([]byte(test.original), &old)
	if err != nil {
		t.Error("Error unmarshalling original for alter.Alter()")
		t.Logf("  original: %s", test.original)
		t.Logf("  path:     %s: %v", test.path, pathS)
		t.Logf("  change:   %s", test.change)
		t.Logf("  expected: %s.(%T)", test.expected, expV)
		t.Logf("  actual:   %s", "n/a")
		t.Logf("  error:    %s", err)
		return
	}

	new := alter.StrToInterface(test.change)

	v, err := alter.Alter(context.TODO(), old, pathS, new)
	if err != nil {
		t.Error("Error received from alter.Alter()")
		t.Logf("  original: %s", test.original)
		t.Logf("  path:     %s: %v", test.path, pathS)
		t.Logf("  change:   %s", test.change)
		t.Logf("  expected: %s.(%T)", test.expected, expV)
		t.Logf("  actual:   %s", "n/a")
		t.Logf("  error:    %s", err)
		return
	}

	actual, err := json.Marshal(v)
	if err != nil {
		t.Error("Error marshalling v from alter.Alter()")
		t.Logf("  original: %s", test.original)
		t.Logf("  path:     %s: %v", test.path, pathS)
		t.Logf("  change:   %s", test.change)
		t.Logf("  expected: %s.(%T)", test.expected, expV)
		t.Logf("  actual:   %s", "n/a")
		t.Logf("  error:    %s", err)
		return
	}

	if string(actual) != test.expected {
		t.Error("Expected does not match actual")
		t.Logf("  original: %s.(%T)", test.original, old)
		t.Logf("  path:     %s: %v", test.path, pathS)
		t.Logf("  change:   %s", test.change)
		t.Logf("  expected: %s.(%T)", test.expected, expV)
		t.Logf("  actual:   %s.(%T)", string(actual), v)
		t.Logf("  error:    %s", "nil")
	}
}

func TestUpdateMap(t *testing.T) {
	test := plan{
		original: `{"1": "foo", "2": "bar"}`,
		path:     "/2",
		change:   `test`,
		expected: `{
						"1": "foo",
						"2": "test"
					}`,
	}

	alterTest(t, &test)
}

func TestUpdateNestedMap(t *testing.T) {
	test := plan{
		original: `
			{
				"1": "foo",
				"2": "bar",
				"3": {
					"a": "aye",
					"b": "bee",
					"c": "cee"
				}
			}`,
		path: "/3",
		change: `
			{
				"d": "dee",
				"e": "ee",
				"f": "eff"
			}`,
		expected: `
		{
			"1": "foo",
			"2": "bar",
			"3": {
				"d": "dee",
				"e": "ee",
				"f": "eff"
			}
		}`,
	}

	alterTest(t, &test)
}

func TestNewMap(t *testing.T) {
	test := plan{
		original: `{"1": "foo", "2": "bar"}`,
		path:     "/3",
		change:   `{"3": "test"}`,
		expected: `{
						"1": "foo",
						"2": "bar",
						"3": {
							"3": "test"
						}
					}`,
	}

	alterTest(t, &test)
}

func TestUpdateArrayAlpha(t *testing.T) {
	test := plan{
		original: `["foo", "bar"]`,
		path:     "/1",
		change:   `test`,
		expected: `[
						"foo",
						"test"
					]`,
	}

	alterTest(t, &test)
}

func TestUpdateArrayNumeric(t *testing.T) {
	test := plan{
		original: `[1, 2]`,
		path:     "/1",
		change:   `3`,
		expected: `[
						1,
						3
					]`,
	}

	alterTest(t, &test)
}

func TestNewArray(t *testing.T) {
	test := plan{
		original: `{"1": "foo", "2": "bar"}`,
		path:     "/3",
		change:   `[4, 5, 6]`,
		expected: `{
						"1": "foo",
						"2": "bar",
						"3": [
							4,
							5,
							6
						]
					}`,
	}

	alterTest(t, &test)
}

func TestUpdateNestedArrayAlpha(t *testing.T) {
	test := plan{
		original: `
			{
				"1": "foo",
				"2": "bar",
				"3": [
					"aye",
					"bee",
					"cee"
				]
			}`,
		path: "/3",
		change: `
			[
				"dee",
				"ee",
				"eff"
			]`,
		expected: `
		{
			"1": "foo",
			"2": "bar",
			"3": [
				"dee",
				"ee",
				"eff"
			]
		}`,
	}

	alterTest(t, &test)
}

func TestUpdateNestedArrayNumeric(t *testing.T) {
	test := plan{
		original: `
			{
				"1": "foo",
				"2": "bar",
				"3": [
					4,
					5,
					6
				]
			}`,
		path: "/3",
		change: `
			[
				7,
				8,
				9
			]`,
		expected: `
		{
			"1": "foo",
			"2": "bar",
			"3": [
				7,
				8,
				9
			]
		}`,
	}

	alterTest(t, &test)
}

func TestUpdateArrayDiffDataTypesInt(t *testing.T) {
	test := plan{
		original: `[ 1, 2, 3 ]`,
		path:     "/1",
		change:   `10`,
		expected: `[ 1, 10, 3 ]`,
	}

	alterTest(t, &test)
}

func TestUpdateArrayDiffDataTypesFloat(t *testing.T) {
	test := plan{
		original: `[ 1.1, 2.2, 3.3 ]`,
		path:     "/1",
		change:   `10.01`,
		expected: `[ 1.1, 10.01, 3.3 ]`,
	}

	alterTest(t, &test)
}

func TestUpdateArrayDiffDataTypesBoolTrue(t *testing.T) {
	test := plan{
		original: `[ true, true, true ]`,
		path:     "/1",
		change:   `false`,
		expected: `[ true, false, true ]`,
	}

	alterTest(t, &test)
}

func TestUpdateArrayDiffDataTypesBoolFalse(t *testing.T) {
	test := plan{
		original: `[ false, false, false ]`,
		path:     "/1",
		change:   `true`,
		expected: `[ false, true, false ]`,
	}

	alterTest(t, &test)
}

func TestUpdateArrayDiffDataTypesString(t *testing.T) {
	test := plan{
		original: `[ "a", "b", "c" ]`,
		path:     "/1",
		change:   `z`,
		expected: `[ "a", "z", "c" ]`,
	}

	alterTest(t, &test)
}

func TestUpdateArrayFloat(t *testing.T) {
	test := plan{
		original: `[ 1.1, 2.2, 3.3 ]`,
		path:     "/1",
		change:   `4.4`,
		expected: `[ 1.1, 4.4, 3.3 ]`,
	}

	alterTest(t, &test)
}

// https://github.com/lmorg/murex/issues/850
func TestUpdateIssue850(t *testing.T) {
	test := plan{
		original: `{}`,
		path:     "/",
		change:   `{ "key": [{"hello": "world"}] }`,
		expected: `{"":{"key":[{"hello":"world"}]}}`,
	}

	alterTest(t, &test)
}

