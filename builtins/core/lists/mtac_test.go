package lists

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/types/json"
	_ "github.com/lmorg/murex/builtins/types/string"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
)

type methodTest struct {
	Stdin  string
	Stdout string
	Error  error
}

func TestMtacString(t *testing.T) {
	tests := []methodTest{
		{
			Stdin:  "a\nb\nc\nd\ne\n",
			Stdout: "e\nd\nc\nb\na\n",
		},
		{
			Stdin:  "a\nb\nc\nd\ne",
			Stdout: "e\nd\nc\nb\na\n",
		},
		{
			Stdin:  "e\nd\nc\nb\na\n",
			Stdout: "a\nb\nc\nd\ne\n",
		},
		{
			Stdin:  "e\nd\nc\nb\na",
			Stdout: "a\nb\nc\nd\ne\n",
		},
		{
			Stdin:  "aaa\nbbb\nccc\nddd\neee\n",
			Stdout: "eee\nddd\nccc\nbbb\naaa\n",
		},
		{
			Stdin:  "aaa\nbbb\nccc\nddd\neee",
			Stdout: "eee\nddd\nccc\nbbb\naaa\n",
		},
		{
			Stdin:  "eee\nddd\nccc\nbbb\naaa\n",
			Stdout: "aaa\nbbb\nccc\nddd\neee\n",
		},
		{
			Stdin:  "eee\nddd\nccc\nbbb\naaa",
			Stdout: "aaa\nbbb\nccc\nddd\neee\n",
		},
		{
			Stdin:  "1\n2\n3\n4\n5\n",
			Stdout: "5\n4\n3\n2\n1\n",
		},
		{
			Stdin:  "1\n2\n3\n4\n5",
			Stdout: "5\n4\n3\n2\n1\n",
		},
		{
			Stdin:  "5\n4\n3\n2\n1\n",
			Stdout: "1\n2\n3\n4\n5\n",
		},
		{
			Stdin:  "5\n4\n3\n2\n1",
			Stdout: "1\n2\n3\n4\n5\n",
		},
	}

	for i := range tests {
		test.RunMethodTest(t, cmdMtac, "mtac", tests[i].Stdin, types.String, []string{}, tests[i].Stdout, nil)
	}
}

func TestMtacGeneric(t *testing.T) {
	tests := []methodTest{
		{
			Stdin:  "a\nb\nc\nd\ne\n",
			Stdout: "e\nd\nc\nb\na\n",
		},
		{
			Stdin:  "a\nb\nc\nd\ne",
			Stdout: "e\nd\nc\nb\na\n",
		},
		{
			Stdin:  "e\nd\nc\nb\na\n",
			Stdout: "a\nb\nc\nd\ne\n",
		},
		{
			Stdin:  "e\nd\nc\nb\na",
			Stdout: "a\nb\nc\nd\ne\n",
		},
		{
			Stdin:  "aaa\nbbb\nccc\nddd\neee\n",
			Stdout: "eee\nddd\nccc\nbbb\naaa\n",
		},
		{
			Stdin:  "aaa\nbbb\nccc\nddd\neee",
			Stdout: "eee\nddd\nccc\nbbb\naaa\n",
		},
		{
			Stdin:  "eee\nddd\nccc\nbbb\naaa\n",
			Stdout: "aaa\nbbb\nccc\nddd\neee\n",
		},
		{
			Stdin:  "eee\nddd\nccc\nbbb\naaa",
			Stdout: "aaa\nbbb\nccc\nddd\neee\n",
		},
		{
			Stdin:  "1\n2\n3\n4\n5\n",
			Stdout: "5\n4\n3\n2\n1\n",
		},
		{
			Stdin:  "1\n2\n3\n4\n5",
			Stdout: "5\n4\n3\n2\n1\n",
		},
		{
			Stdin:  "5\n4\n3\n2\n1\n",
			Stdout: "1\n2\n3\n4\n5\n",
		},
		{
			Stdin:  "5\n4\n3\n2\n1",
			Stdout: "1\n2\n3\n4\n5\n",
		},
	}

	for i := range tests {
		test.RunMethodTest(t, cmdMtac, "mtac", tests[i].Stdin, types.Generic, []string{}, tests[i].Stdout, nil)
	}
}

func TestMtacJsonString(t *testing.T) {
	tests := []methodTest{
		{
			Stdin:  `["a","b","c","d","e"]`,
			Stdout: `["e","d","c","b","a"]`,
		},
		{
			Stdin:  `["e","d","c","b","a"]`,
			Stdout: `["a","b","c","d","e"]`,
		},
		{
			Stdin:  `["aaa","bbb","ccc","ddd","eee"]`,
			Stdout: `["eee","ddd","ccc","bbb","aaa"]`,
		},
		{
			Stdin:  `["eee","ddd","ccc","bbb","aaa"]`,
			Stdout: `["aaa","bbb","ccc","ddd","eee"]`,
		},
		{
			Stdin:  `["1","2","3","4","5"]`,
			Stdout: `["5","4","3","2","1"]`,
		},
		{
			Stdin:  `["5","4","3","2","1"]`,
			Stdout: `["1","2","3","4","5"]`,
		},
		{
			Stdin:  `["111","222","333","444","555"]`,
			Stdout: `["555","444","333","222","111"]`,
		},
		{
			Stdin:  `["555","444","333","222","111"]`,
			Stdout: `["111","222","333","444","555"]`,
		},
	}

	for i := range tests {
		test.RunMethodTest(t, cmdMtac, "mtac", tests[i].Stdin, types.Json, []string{}, tests[i].Stdout, nil)
	}
}

func TestMtacJsonInt(t *testing.T) {
	tests := []methodTest{
		{
			Stdin:  `[1,2,3,4,5]`,
			Stdout: `[5,4,3,2,1]`,
		},
		{
			Stdin:  `[5,4,3,2,1]`,
			Stdout: `[1,2,3,4,5]`,
		},
		{
			Stdin:  `[111,222,333,444,555]`,
			Stdout: `[555,444,333,222,111]`,
		},
		{
			Stdin:  `[555,444,333,222,111]`,
			Stdout: `[111,222,333,444,555]`,
		},
	}

	for i := range tests {
		test.RunMethodTest(t, cmdMtac, "mtac", tests[i].Stdin, types.Json, []string{}, tests[i].Stdout, nil)
	}
}
