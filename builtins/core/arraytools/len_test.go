package arraytools

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/types/generic"
	_ "github.com/lmorg/murex/builtins/types/json"
	_ "github.com/lmorg/murex/builtins/types/jsonlines"
	_ "github.com/lmorg/murex/builtins/types/string"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
)

func TestLenJsonLines(t *testing.T) {
	tests := []string{
		"{\"a\":1,\"b\":2}\n{\"a\":1,\"b\":2}\n{\"a\":1,\"b\":2}\n",
		"{\"a\":1,\"b\":2}\n{\"a\":1,\"b\":2}\n{\"a\":1,\"b\":2}",
		"{\"a\":1,\"b\":2}\r\n{\"a\":1,\"b\":2}\r\n{\"a\":1,\"b\":2}\r\n",
		"{\"a\":1,\"b\":2}\r\n{\"a\":1,\"b\":2}\r\n{\"a\":1,\"b\":2}",

		"[1,2,3,4,5]\n[1,2,3,4,5]\n[1,2,3,4,5]\n",
		"[1,2,3,4,5]\r\n[1,2,3,4,5]\r\n[1,2,3,4,5]\r\n",
		"[1,2,3,4,5]\n[1,2,3,4,5]\n[1,2,3,4,5]",
		"[1,2,3,4,5]\r\n[1,2,3,4,5]\r\n[1,2,3,4,5]",
	}

	for _, s := range tests {
		test.RunMethodTest(t,
			cmdLen, "len",
			s,
			types.JsonLines,
			[]string{},
			`3`,
			nil,
		)
	}
}

func TestLenJson(t *testing.T) {
	tests := []string{
		`[1,2,3,4,5]`,
		`["1","2","3","4","5"]`,
		`["a","b","c","d","e"]`,
		`["abc","def","ghi","jkl","mno"]`,
		`{"1":[], "2":[], "3":[], "4":[], "5":[]}`,
		`{"1":[1,2], "2":[3,4], "3":[5,6], "4":[7,8], "5":[9,0]}`,
		`{"1":{}, "2":{}, "3":{}, "4":{}, "5":{}}`,
		`{"1":{"a":1,"b":2}, "2":{"c":3,"d":4}, "3":{"e":5,"f":6}, "4":{"g":7,"h":8}, "5":{"i":9,"j":0}}`,
	}

	for _, s := range tests {
		test.RunMethodTest(t,
			cmdLen, "len",
			s,
			types.JsonLines,
			[]string{},
			`5`,
			nil,
		)
	}
}

func TestLenString1(t *testing.T) {
	tests := []string{
		"1 2 3 4 5",
		"a b c d e",
		"abc def ghi jkl mno",
		"1\t2\t3\t4\t5",
		"a\tb\tc\td\te",
		"abc\tdef\tghi\tjkl\tmno",
	}

	for _, s := range tests {
		test.RunMethodTest(t,
			cmdLen, "len",
			s,
			types.String,
			[]string{},
			`1`,
			nil,
		)
	}
}

func TestLenString5(t *testing.T) {
	tests := []string{
		"1\n2\n3\n4\n5\n",
		"a\nb\nc\nd\ne\n",
		"abc\ndef\nghi\njkl\nmno\n",
		"1\r\n2\r\n3\r\n4\r\n5\r\n",
		"a\r\nb\r\nc\r\nd\r\ne\r\n",
		"abc\r\ndef\r\nghi\r\njkl\r\nmno\r\n",
		"1\n2\n3\n4\n5",
		"a\nb\nc\nd\ne",
		"abc\ndef\nghi\njkl\nmno",
		"1\r\n2\r\n3\r\n4\r\n5",
		"a\r\nb\r\nc\r\nd\r\ne",
		"abc\r\ndef\r\nghi\r\njkl\r\nmno",
	}

	for _, s := range tests {
		test.RunMethodTest(t,
			cmdLen, "len",
			s,
			types.String,
			[]string{},
			`5`,
			nil,
		)
	}
}

func TestLenString6(t *testing.T) {
	tests := []string{
		"\n1\n2\n3\n4\n5\n",
		"\na\nb\nc\nd\ne\n",
		"\nabc\ndef\nghi\njkl\nmno\n",
		"\r\n1\r\n2\r\n3\r\n4\r\n5\r\n",
		"\r\na\r\nb\r\nc\r\nd\r\ne\r\n",
		"\r\nabc\r\ndef\r\nghi\r\njkl\r\nmno\r\n",
		"\r\n1\n2\n3\n4\n5",
		"\r\na\nb\nc\nd\ne",
		"\r\nabc\ndef\nghi\njkl\nmno",
		"\r\n1\r\n2\r\n3\r\n4\r\n5",
		"\r\na\r\nb\r\nc\r\nd\r\ne",
		"\r\nabc\r\ndef\r\nghi\r\njkl\r\nmno",
	}

	for _, s := range tests {
		test.RunMethodTest(t,
			cmdLen, "len",
			s,
			types.String,
			[]string{},
			`6`,
			nil,
		)
	}
}

func TestLenGeneric1(t *testing.T) {
	tests := []string{
		"1 2 3 4 5",
		"a b c d e",
		"abc def ghi jkl mno",
		"1\t2\t3\t4\t5",
		"a\tb\tc\td\te",
		"abc\tdef\tghi\tjkl\tmno",
	}

	for _, s := range tests {
		test.RunMethodTest(t,
			cmdLen, "len",
			s,
			types.Generic,
			[]string{},
			`1`,
			nil,
		)
	}
}

func TestLenGeneric5(t *testing.T) {
	tests := []string{
		"1\n2\n3\n4\n5\n",
		"a\nb\nc\nd\ne\n",
		"abc\ndef\nghi\njkl\nmno\n",
		"1\r\n2\r\n3\r\n4\r\n5\r\n",
		"a\r\nb\r\nc\r\nd\r\ne\r\n",
		"abc\r\ndef\r\nghi\r\njkl\r\nmno\r\n",
		"1\n2\n3\n4\n5",
		"a\nb\nc\nd\ne",
		"abc\ndef\nghi\njkl\nmno",
		"1\r\n2\r\n3\r\n4\r\n5",
		"a\r\nb\r\nc\r\nd\r\ne",
		"abc\r\ndef\r\nghi\r\njkl\r\nmno",
	}

	for _, s := range tests {
		test.RunMethodTest(t,
			cmdLen, "len",
			s,
			types.Generic,
			[]string{},
			`5`,
			nil,
		)
	}
}

func TestLenGeneric6(t *testing.T) {
	tests := []string{
		"\n1\n2\n3\n4\n5\n",
		"\na\nb\nc\nd\ne\n",
		"\nabc\ndef\nghi\njkl\nmno\n",
		"\r\n1\r\n2\r\n3\r\n4\r\n5\r\n",
		"\r\na\r\nb\r\nc\r\nd\r\ne\r\n",
		"\r\nabc\r\ndef\r\nghi\r\njkl\r\nmno\r\n",
		"\r\n1\n2\n3\n4\n5",
		"\r\na\nb\nc\nd\ne",
		"\r\nabc\ndef\nghi\njkl\nmno",
		"\r\n1\r\n2\r\n3\r\n4\r\n5",
		"\r\na\r\nb\r\nc\r\nd\r\ne",
		"\r\nabc\r\ndef\r\nghi\r\njkl\r\nmno",
	}

	for _, s := range tests {
		test.RunMethodTest(t,
			cmdLen, "len",
			s,
			types.Generic,
			[]string{},
			`6`,
			nil,
		)
	}
}
