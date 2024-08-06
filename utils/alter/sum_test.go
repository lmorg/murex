package alter_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/alter"
)

func sumTest(t *testing.T, test *plan) {
	t.Helper()
	count.Tests(t, 1)

	pathS, err := alter.SplitPath(test.path)
	if err != nil {
		panic(err)
	}

	var expV interface{}
	if !test.error {
		if err := json.Unmarshal([]byte(test.expected), &expV); err != nil {
			panic(err)
		}
		b, err := json.Marshal(expV)
		if err != nil {
			panic(err)
		}
		test.expected = string(b)
	}

	var old interface{}
	err = json.Unmarshal([]byte(test.original), &old)
	if err != nil {
		t.Error("Error unmarshalling original for alter.Sum()")
		t.Logf("  original: %s", test.original)
		t.Logf("  path:     %s: %v", test.path, pathS)
		t.Logf("  change:   %s", test.change)
		t.Logf("  expected: %s.(%T)", test.expected, expV)
		t.Logf("  actual:   %s", "n/a")
		t.Logf("  error:    %s", err)
		return
	}

	new := alter.StrToInterface(test.change)
	v, err := alter.Sum(context.TODO(), old, pathS, new)
	if (err != nil) != test.error {
		t.Error("Error received from alter.Sum()")
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
		t.Error("Error marshalling v from alter.Sum()")
		t.Logf("  original: %s", test.original)
		t.Logf("  path:     %s: %v", test.path, pathS)
		t.Logf("  change:   %s", test.change)
		t.Logf("  expected: %s.(%T)", test.expected, expV)
		t.Logf("  actual:   %s", "n/a")
		t.Logf("  error:    %s", err)
		return
	}

	if test.error {
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

// https://github.com/lmorg/murex/issues/850
func TestSumIssue850(t *testing.T) {
	test := plan{
		original: `{}`,
		path:     "/",
		change:   `{ "key": [{"hello": "world"}] }`,
		error:    true,
	}

	sumTest(t, &test)
}
