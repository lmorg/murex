package expressions

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/mkarray"
	_ "github.com/lmorg/murex/builtins/types/generic"
	_ "github.com/lmorg/murex/builtins/types/json"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/test"
)

func TestParseArray(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.ArrayBegin,
		tests: []expTestT{
			{
				input:    `%[1 2 3]`,
				expected: `[1,2,3]`,
				pos:      6,
			},
			{
				input: `%[1 2 3`,
				error: true,
			},
			{
				input:    `%[1  2  3]`,
				expected: `[1,2,3]`,
				pos:      8,
			},
			{
				input:    `%[1,2,3]`,
				expected: `[1,2,3]`,
				pos:      6,
			},
			{
				input:    `%[1, 2, 3]`,
				expected: `[1,2,3]`,
				pos:      8,
			},
			{
				input:    `%[  1  ,  2  ,  3  ]`,
				expected: `[1,2,3]`,
				pos:      18,
			},
			/////
			{
				input:    `%[  foo  ,  bar  ,  baz  ]`,
				expected: `["foo","bar","baz"]`,
				pos:      24,
			},
			/////
			{
				input: `%[  [1 2 3]  ,  [ "foo" "bar" ]`,
				error: true,
			},
			{
				input: `%[%]`,
				error: true,
			},
			{
				input:    `%[%(%)]`,
				expected: `["%"]`,
				pos:      5,
			},
			{
				input:    `%[  [1 2 3]  ,  [ "foo" "bar" ] ]`,
				expected: `[[1,2,3],["foo","bar"]]`,
				pos:      31,
			},
			{
				input:    `%[  %[1 2 3]  ,  %[ "foo" "bar" ] ]`,
				expected: `[[1,2,3],["foo","bar"]]`,
				pos:      33,
			},
			/////
			{
				input:    "%[\n\t1\n\t2\n\t3\n]",
				expected: `[1,2,3]`,
				pos:      11,
			},
			/////
			{
				input:    "%[$TestParseArray]",
				expected: `[null]`,
				pos:      16,
			},
			{
				input:    "%[1,2,$TestParseArray]",
				expected: `[1,2,null]`,
				pos:      20,
			},
			{
				input:    "%[@TestParseArray]",
				expected: `[]`,
				pos:      16,
			},
			{
				input:    "%[[@TestParseArray]]",
				expected: `[[]]`,
				pos:      18,
			},
			/////
			{
				input:    "%[[mon..wed]]",
				expected: `["mon","tue","wed"]`,
				pos:      11,
			},
			/////
			{
				input:    "%[-2,1,0,3.4]",
				expected: `[-2,1,0,3.4]`,
				pos:      11,
			},
			{
				input:    "%[-2 1 0 3.4]",
				expected: `[-2,1,0,3.4]`,
				pos:      11,
			},
			{
				input:    "%[-]",
				expected: `["-"]`,
				pos:      2,
			},
			{
				input:    "%[-one]",
				expected: `["-one"]`,
				pos:      5,
			},
			{
				input:    "%[{a:1} {b:2}]",
				expected: `[{"a":1},{"b":2}]`,
				pos:      12,
			},
			/////
			{
				input:    "%[1..3]",
				expected: `[1,2,3]`,
				pos:      5,
			},
		},
	}

	testParserObject(t, tests)
}

func TestParseArrayBarewords(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.ArrayBegin,
		tests: []expTestT{
			{
				input:    `%[false,null,true]`,
				expected: `[false,null,true]`,
				pos:      16,
			},
			{
				input:    `%[false null true]`,
				expected: `[false,null,true]`,
				pos:      16,
			},
			{
				input:    "%[\nfalse\nnull\ntrue\n]",
				expected: `[false,null,true]`,
				pos:      18,
			},
			{
				input:    `%["false" "null" "true"]`,
				expected: `["false","null","true"]`,
				pos:      22,
			},
		},
	}

	testParserObject(t, tests)
}

func TestParseArrayWithArrayVar(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `bob = %{a:1, b:2, c:[a b c]}; %[x y z @bob[c] ]`,
			Stdout: `["x","y","z","a","b","c"]`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseArrayObjects(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.ArrayBegin,
		tests: []expTestT{
			{
				input:    `%[{a:1},{b:2},{c:3}]`,
				expected: `[{"a":1},{"b":2},{"c":3}]`,
				pos:      18,
			},
			{
				input: `%[
					{
						a:1
					},
					{
						b:2
					},
					{
						c:3
					}
				]`,
				expected: `[{"a":1},{"b":2},{"c":3}]`,
				pos:      80,
			},
		},
	}

	testParserObject(t, tests)
}

func TestParseArrayBugfix832(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `%[]`,
			Stdout: `^\[\]$`,
		},
		{
			Block:  `out %[]`,
			Stdout: `^\[\]\n$`,
		},
		{
			Block:  `TestParseArrayBugfix832_3 = %[]; out $TestParseArrayBugfix832_3`,
			Stdout: `^\[\]\n$`,
		},
		{
			Block:   `TestParseArrayBugfix832_4 = %[]; out @TestParseArrayBugfix832_4`,
			Stderr:  `Error`,
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}

func TestParseArrayNestedExpr(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.ArrayBegin,
		tests: []expTestT{
			{
				input:    "%[(1+1),(2+2)]",
				expected: `[2,4]`,
				pos:      12,
			},
			{
				input: "%[(1+-),(2+2)]",
				error: true,
			},
			{
				input: "%[(1+1),(2+-)]",
				error: true,
			},
		},
	}

	testParserObject(t, tests)
}
