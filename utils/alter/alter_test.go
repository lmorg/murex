package alter_test

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/alter"
)

type plan struct {
	original string
	path     string
	change   string
	expected string
}

func alterTest(t *testing.T, test *plan) {
	count.Tests(t, 1)

	pathS := strings.Split(test.path, "/")[1:]

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

	v, err := alter.Alter(context.TODO(), old, pathS, test.change)
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

func TestUpdateNest(t *testing.T) {
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

func TestUpdateArray(t *testing.T) {
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
