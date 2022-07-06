package datatools_test

import (
	"fmt"
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/utils/json"
)

type plan struct {
	original string
	path     string
	change   string
	expected string
}

func reMarshal(s string) string {
	var v interface{}
	err := json.Unmarshal([]byte(s), &v)
	if err != nil {
		panic(err.Error())
	}
	b, err := json.Marshal(&v, false)
	if err != nil {
		panic(err.Error())
	}
	return string(b)
}

func alterTest(t *testing.T, p *plan, flag string) {
	test.RunMurexTests([]test.MurexTest{
		{
			Block: fmt.Sprintf("tout json (%s) -> alter: %s %s (%s)",
				p.original, flag, p.path, p.change),
			Stdout: reMarshal(p.expected),
		},
	}, t)
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

	alterTest(t, &test, "")
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

	alterTest(t, &test, "")
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

	alterTest(t, &test, "")
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

	alterTest(t, &test, "")
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

	alterTest(t, &test, "")
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

	alterTest(t, &test, "")
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

	alterTest(t, &test, "")
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

	alterTest(t, &test, "")
}

func TestUpdateArrayDiffDataTypesInt(t *testing.T) {
	test := plan{
		original: `[ 1, 2, 3 ]`,
		path:     "/1",
		change:   `10`,
		expected: `[ 1, 10, 3 ]`,
	}

	alterTest(t, &test, "")
}

func TestUpdateArrayDiffDataTypesFloat(t *testing.T) {
	test := plan{
		original: `[ 1.1, 2.2, 3.3 ]`,
		path:     "/1",
		change:   `10.01`,
		expected: `[ 1.1, 10.01, 3.3 ]`,
	}

	alterTest(t, &test, "")
}

func TestUpdateArrayDiffDataTypesBoolTrue(t *testing.T) {
	test := plan{
		original: `[ true, true, true ]`,
		path:     "/1",
		change:   `false`,
		expected: `[ true, false, true ]`,
	}

	alterTest(t, &test, "")
}

func TestUpdateArrayDiffDataTypesBoolFalse(t *testing.T) {
	test := plan{
		original: `[ false, false, false ]`,
		path:     "/1",
		change:   `true`,
		expected: `[ false, true, false ]`,
	}

	alterTest(t, &test, "")
}

func TestUpdateArrayDiffDataTypesString(t *testing.T) {
	test := plan{
		original: `[ "a", "b", "c" ]`,
		path:     "/1",
		change:   `z`,
		expected: `[ "a", "z", "c" ]`,
	}

	alterTest(t, &test, "")
}

func TestUpdateArrayFloat(t *testing.T) {
	test := plan{
		original: `[ 1.1, 2.2, 3.3 ]`,
		path:     "/1",
		change:   `4.4`,
		expected: `[ 1.1, 4.4, 3.3 ]`,
	}

	alterTest(t, &test, "")
}
