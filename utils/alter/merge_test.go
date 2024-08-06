package alter_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/alter"
)

func mergeTest(t *testing.T, test *plan) {
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
		t.Error("Error unmarshalling original for alter.Merge()")
		t.Logf("  original: %s", test.original)
		t.Logf("  path:     %s: %v", test.path, pathS)
		t.Logf("  change:   %s", test.change)
		t.Logf("  expected: %s.(%T)", test.expected, expV)
		t.Logf("  actual:   %s", "n/a")
		t.Logf("  error:    %s", err)
		return
	}

	new := alter.StrToInterface(test.change)

	v, err := alter.Merge(context.TODO(), old, pathS, new)
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
		t.Error("Error marshalling v from alter.Merge()")
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

func TestMergeMap(t *testing.T) {
	test := plan{
		original: `
			{
				"a": "aye",
				"b": "bee",
				"c": "cee"
			}`,
		path: "/",
		change: `
			{
				"d": "dee",
				"e": "ee",
				"f": "eff"
			}`,
		expected: `
			{
				"a": "aye",
				"b": "bee",
				"c": "cee",
				"d": "dee",
				"e": "ee",
				"f": "eff"
			}`,
	}

	mergeTest(t, &test)
}

func TestMergeAndUpdateMap(t *testing.T) {
	test := plan{
		original: `
			{
				"a": "aye",
				"b": "bee",
				"c": "cee"
			}`,
		path: "/",
		change: `
			{
				"c": "update",
				"e": "ee",
				"f": "eff"
			}`,
		expected: `
			{
				"a": "aye",
				"b": "bee",
				"c": "update",
				"e": "ee",
				"f": "eff"
			}`,
	}

	mergeTest(t, &test)
}

func TestMergeArrayAlpha(t *testing.T) {
	test := plan{
		original: `
			[
				"aye",
				"bee",
				"cee"
			]`,
		path: "/",
		change: `
			[
				"dee",
				"ee",
				"eff"
			]`,
		expected: `
			[
				"aye",
				"bee",
				"cee",
				"dee",
				"ee",
				"eff"
			]`,
	}

	mergeTest(t, &test)
}

func TestMergeArrayNumeric(t *testing.T) {
	test := plan{
		original: `
			[
				1,
				2,
				3
			]`,
		path: "/",
		change: `
			[
				5,
				6,
				7
			]`,
		expected: `
			[
				1,
				2,
				3,
				5,
				6,
				7
			]`,
	}

	mergeTest(t, &test)
}

func TestMergeNestedArrayAlpha(t *testing.T) {
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
				"aye",
				"bee",
				"cee",
				"dee",
				"ee",
				"eff"
			]
		}`,
	}

	mergeTest(t, &test)
}

func TestMergeNestedArrayNumeric(t *testing.T) {
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
				4,
				5,
				6,
				7,
				8,
				9
			]
		}`,
	}

	mergeTest(t, &test)
}

func TestMergeNestedArrayMap(t *testing.T) {
	test := plan{
		original: `
			[
				{
					"foo": 1
				}
			]`,
		path: "/",
		change: `
			[
				{
					"bar": 2
				}
			]`,
		expected: `
		[
			{
				"foo": 1
			},
			{
				"bar": 2
			}
		]`,
	}

	mergeTest(t, &test)
}

func TestMergeNestedArrayArray(t *testing.T) {
	test := plan{
		original: `
			[
				[ 1, 2, 3 ]
			]`,
		path: "/",
		change: `
			[
				[ 6, 7, 8 ]
			]`,
		expected: `
			[
				[ 1, 2, 3 ],
				[ 6, 7, 8 ]
			]
		`,
	}

	mergeTest(t, &test)
}

func TestMergeNestedMapArray(t *testing.T) {
	test := plan{
		original: `
			{
				"foo": [ 1, 2, 3 ]
			}`,
		path: "/",
		change: `
			{
				"foo": [ 6, 7, 8 ]
			}`,
		expected: `
		{
			"foo": [
				1,
				2,
				3,
				6,
				7,
				8
			]
		}`,
	}

	mergeTest(t, &test)
}

func TestMergeNestedMapMap(t *testing.T) {
	test := plan{
		original: `
			{
				"foo": { "a": 1, "b": 2 }
			}`,
		path: "/",
		change: `
			{
				"foo": { "c": 3, "d": 4 }
			}`,
		expected: `
		{
			"foo": {
				"a": 1,
				"b": 2,
				"c": 3,
				"d": 4
			}
		}`,
	}

	mergeTest(t, &test)
}

// https://github.com/lmorg/murex/issues/850
func TestMergeIssue850n1(t *testing.T) {
	test := plan{
		original: `{}`,
		path:     "/",
		change:   `{ "key": [{"hello": "world"}] }`,
		expected: `{"key":[{"hello":"world"}]}`,
	}

	mergeTest(t, &test)
}

// https://github.com/lmorg/murex/issues/850
func TestMergeIssue850n2(t *testing.T) {
	test := plan{
		original: `{"key":[{"hello":"world"}]}`,
		path:     "/",
		change:   `{ "key": [{"hello": "earth"}] }`,
		expected: `{"key":[{"hello":"world"},{"hello":"earth"}]}`,
	}

	mergeTest(t, &test)
}

// https://github.com/lmorg/murex/issues/850
func TestMergeIssue850n3(t *testing.T) {
	test := plan{
		original: `{"key":[{"hello":"world"}]}`,
		path:     "/",
		change:   `{ "key": [] }`,
		expected: `{"key":[{"hello":"world"}]}`,
	}

	mergeTest(t, &test)
}

// https://github.com/lmorg/murex/issues/850
func TestMergeIssue850n4(t *testing.T) {
	test := plan{
		original: `{"key":[{"hello":"world"}]}`,
		path:     "/",
		change:   `{ "key2": [{"hello":"earth"}] }`,
		expected: `{"key":[{"hello":"world"}],"key2":[{"hello":"earth"}]}`,
	}

	mergeTest(t, &test)
}
