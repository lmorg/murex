package alter_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/alter"
)

func sumTest(t *testing.T, test *plan) {
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

	v, err := alter.Sum(context.TODO(), old, pathS, test.change)
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

func TestSumMapInt(t *testing.T) {
	test := plan{
		original: `
			{
				"a": 1,
				"b": 2,
				"c": 3
			}`,
		path: "/",
		change: `
			{
				"a": 9,
				"b": 7,
				"c": 5
			}`,
		expected: `
			{
				"a": 10,
				"b": 9,
				"c": 8
			}`,
	}

	sumTest(t, &test)
}

func TestSumMapFloat64(t *testing.T) {
	test := plan{
		original: `
			{
				"a": 1.1,
				"b": 2.2,
				"c": 3.3
			}`,
		path: "/",
		change: `
			{
				"a": 9.9,
				"b": 7.7,
				"c": 5.5
			}`,
		expected: `
			{
				"a": 11,
				"b": 9.9,
				"c": 8.8
			}`,
	}

	sumTest(t, &test)
}

func TestSumMapInterface(t *testing.T) {
	test := plan{
		original: `
			{
				"a": "1",
				"b": "2",
				"c": "3"
			}`,
		path: "/",
		change: `
			{
				"a": "9",
				"b": "7.5",
				"c": "5"
			}`,
		expected: `
			{
				"a": 10,
				"b": 9.5,
				"c": 8
			}`,
	}

	sumTest(t, &test)
}

/*
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
*/
