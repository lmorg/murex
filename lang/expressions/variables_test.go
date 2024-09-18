package expressions

import (
	"testing"

	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/test/count"
)

type createIndexBlockTestT struct {
	Name     string
	Index    string
	Flags    string
	Expected string
}

func TestCreateIndexBlock(t *testing.T) {
	tests := []createIndexBlockTestT{
		{
			Name:  "foobar",
			Index: "hello world",
		},
		{
			Name:  "",
			Index: "",
		},
		{
			Name:  "a",
			Index: "b",
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		if test.Expected == "" {
			test.Expected = "$" + test.Name + "-> [" + test.Index + "]"
		}

		getIorE := &getVarIndexOrElementT{varName: []rune(test.Name), key: []rune(test.Index)}
		block := createIndexBlock(getIorE)
		if string(block) != test.Expected {
			t.Errorf("block does not match expected in test %d", i)
			t.Logf("  Name:     '%s'", test.Name)
			t.Logf("  Index:    '%s'", test.Index)
			t.Logf("  Expected: '%s'", test.Expected)
			t.Logf("  Actual:   '%s'", string(block))
		}
	}
}

func TestCreateElementBlock(t *testing.T) {
	tests := []createIndexBlockTestT{
		{
			Name:  "foobar",
			Index: "hello world",
		},
		{
			Name:  "",
			Index: "",
		},
		{
			Name:  "a",
			Index: "b",
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		if test.Expected == "" {
			test.Expected = "$" + test.Name + "-> [[" + test.Index + "]]"
		}

		getIorE := &getVarIndexOrElementT{varName: []rune(test.Name), key: []rune(test.Index)}
		block := createElementBlock(getIorE)
		if string(block) != test.Expected {
			t.Errorf("block does not match expected in test %d", i)
			t.Logf("  Name:     '%s'", test.Name)
			t.Logf("  Index:    '%s'", test.Index)
			t.Logf("  Expected: '%s'", test.Expected)
			t.Logf("  Actual:   '%s'", string(block))
		}
	}
}

func TestCreateRangeBlock(t *testing.T) {
	tests := []createIndexBlockTestT{
		{
			Name:  "foobar",
			Index: "hello world",
			Flags: "abc",
		},
		{
			Name:  "",
			Index: "",
			Flags: "",
		},
		{
			Name:  "a",
			Index: "b",
			Flags: "c",
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		if test.Expected == "" {
			test.Expected = "$" + test.Name + "-> @[" + test.Index + "]" + test.Flags
		}

		block := createRangeBlock([]rune(test.Name), []rune(test.Index), []rune(test.Flags))
		if string(block) != test.Expected {
			t.Errorf("block does not match expected in test %d", i)
			t.Logf("  Name:     '%s'", test.Name)
			t.Logf("  Index:    '%s'", test.Index)
			t.Logf("  Expected: '%s'", test.Expected)
			t.Logf("  Actual:   '%s'", string(block))
		}
	}
}

func TestFormatBytesArrayBuilder(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `function test0 { %{a:1, b:2, c:3} }; %[ ${test0} ]`,
			Stdout: `[{"a":1,"b":2,"c":3}]`,
		},
		/////
		{
			Block:  `function test8 { tout null "" }; %[ ${test8} ]`,
			Stdout: `[[null]]`, // this is potentially unexpected behaviour
		},
		{
			Block:  `function test9 { null };  %[ ${test9} ]`,
			Stdout: `[[null]]`, // this is potentially unexpected behaviour
		},
	}

	test.RunMurexTests(tests, t)
}

func TestFormatBytesObjectBuilder(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `function test0 { %{a:1, b:2, c:3} }; %{${test0}: ${test0}}`,
			Stdout: `{"{\"a\":1,\"b\":2,\"c\":3}":{"a":1,"b":2,"c":3}}`,
		},
		{
			Block:  `function test1 { tout str "test1" }; %{${test1}: ${test1}}`,
			Stdout: `{"test1":"test1"}`,
		},
		{
			Block:  `function test2 { tout str 2 }; %{${test2}: ${test2}}`,
			Stdout: `{"2":"2"}`,
		},
		{
			Block:  `function test3 { tout int 3 }; %{${test3}: ${test3}}`,
			Stdout: `{"3":3}`,
		},
		{
			Block:  `function test4 { tout float 4.1 }; %{${test4}: ${test4}}`,
			Stdout: `{"4.1":4.1}`,
		},
		{
			Block:  `function test5 { tout num 5.2 }; %{${test5}: ${test5}}`,
			Stdout: `{"5.2":5.2}`,
		},
		{
			Block:  `function test6 { true }; %{${test6}: ${test6}}`,
			Stdout: `{"true":true}`,
		},
		{
			Block:  `function test7 { false }; %{${test7}: ${test7}}`,
			Stdout: `{"false":false}`,
		},
		/////
		{
			Block:  `function test8 { tout null "" }; %{${test8}: ${test8}}`,
			Stdout: `{"":[null]}`,
		},
		{
			Block:  `function test9 { null }; %{${test9}: ${test9}}`,
			Stdout: `{"":[null]}`,
		},
	}

	test.RunMurexTests(tests, t)
}
