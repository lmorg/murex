package lang_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

var (
	testsAnd = []test.MurexTest{
		{
			Block:  `out "foo" && out "bar"`,
			Stdout: "foo\nbar\n",
		},
		{
			Block:   `err "foo" && out "bar"`,
			Stderr:  "foo\n",
			ExitNum: 1,
		},
		{
			Block:   `out "foo" && err "bar"`,
			Stdout:  "foo\n",
			Stderr:  "bar\n",
			ExitNum: 1,
		},
		{
			Block:   `err "foo" && err "bar"`,
			Stderr:  "foo\n",
			ExitNum: 1,
		},
		/////
		{
			Block:  `out "foo"&&out "bar"`,
			Stdout: "foo\nbar\n",
		},
		{
			Block:   `err "foo"&&out "bar"`,
			Stderr:  "foo\n",
			ExitNum: 1,
		},
		{
			Block:   `out "foo"&&err "bar"`,
			Stdout:  "foo\n",
			Stderr:  "bar\n",
			ExitNum: 1,
		},
		{
			Block:   `err "foo"&&err "bar"`,
			Stderr:  "foo\n",
			ExitNum: 1,
		},
		/////
		{
			Block:  `out "foo" \&& out "bar"`,
			Stdout: "foo && out bar\n",
		},
		{
			Block:  `out "foo" &\& out "bar"`,
			Stdout: "foo && out bar\n",
		},
		{
			Block:  `out "foo" \&\& out "bar"`,
			Stdout: "foo && out bar\n",
		},
		/////
		{
			Block:  `out "foo" \&&& out "bar"`,
			Stdout: "foo &\nbar\n",
		},
		{
			Block:  `out "foo" &\&& out "bar"`,
			Stdout: "foo &&& out bar\n",
		},
	}

	testsOr = []test.MurexTest{
		{
			Block:  `out "foo" || out "bar"`,
			Stdout: "foo\n",
		},
		{
			Block:  `err "foo" || out "bar"`,
			Stdout: "bar\n",
			Stderr: "foo\n",
		},
		{
			Block:  `out "foo" || err "bar"`,
			Stdout: "foo\n",
		},
		{
			Block:   `err "foo" || err "bar"`,
			Stderr:  "foo\nbar\n",
			ExitNum: 1,
		},
		/////
		{
			Block:  `out "foo"||out "bar"`,
			Stdout: "foo\n",
		},
		{
			Block:  `err "foo"||out "bar"`,
			Stdout: "bar\n",
			Stderr: "foo\n",
		},
		{
			Block:  `out "foo"||err "bar"`,
			Stdout: "foo\n",
		},
		{
			Block:   `err "foo"||err "bar"`,
			Stderr:  "foo\nbar\n",
			ExitNum: 1,
		},
		/////
		{
			Block:  `out "foo" \| | regexp s/f/g/`,
			Stdout: "goo |\n",
		},
	}
)

func TestOperatorLogicAndNormal(t *testing.T) {
	test.RunMurexTests(testsAnd, t)
}

func TestOperatorLogicOrNormal(t *testing.T) {
	test.RunMurexTests(testsOr, t)
}

func TestOperatorLogicAndTry(t *testing.T) {
	var tests []test.MurexTest
	for _, src := range testsAnd {
		newTest := src
		src.Block = "try {" + src.Block + "}"
		tests = append(tests, newTest)
	}

	test.RunMurexTests(tests, t)
}

func TestOperatorLogicOrTry(t *testing.T) {
	var tests []test.MurexTest
	for _, src := range testsOr {
		newTest := src
		src.Block = "try {" + src.Block + "}"
		tests = append(tests, newTest)
	}
	test.RunMurexTests(tests, t)
}

func TestOperatorLogicAndTryPipe(t *testing.T) {
	var tests []test.MurexTest
	for _, src := range testsAnd {
		newTest := src
		src.Block = "trypipe {" + src.Block + "}"
		tests = append(tests, newTest)
	}
	test.RunMurexTests(tests, t)
}

func TestOperatorLogicOrTryPipe(t *testing.T) {
	var tests []test.MurexTest
	for _, src := range testsOr {
		newTest := src
		src.Block = "trypipe {" + src.Block + "}"
		tests = append(tests, newTest)
	}
	test.RunMurexTests(tests, t)
}

func TestOperatorsTry(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `out: 0; false; echo foobar`,
			Stdout: "0\nfalsefoobar\n",
		},
		{
			Block:  `out: 1; true || echo foobar`,
			Stdout: "1\ntrue",
		},
		{
			Block:  `out: 1.1; true || true || echo foobar`,
			Stdout: "1.1\ntrue",
		},
		{
			Block:  `out: 1.2; true || false || echo foobar`,
			Stdout: "1.2\ntrue",
		},
		{
			Block:  `out: 1.3; false || true || echo foobar`,
			Stdout: "1.3\nfalsetrue",
		},
		{
			Block:  `out: 2; true && echo foobar`,
			Stdout: "2\ntruefoobar\n",
		},
		{
			Block:   `out: 2.1; false && false && echo foobar`,
			Stdout:  "2.1\nfalse",
			ExitNum: 1,
		},
		{
			Block:   `out: 2.2; true && false && echo foobar`,
			Stdout:  "2.2\ntruefalse",
			ExitNum: 1,
		},
		{
			Block:   `out: 2.3; false && true && echo foobar`,
			Stdout:  "2.3\nfalse",
			ExitNum: 1,
		},
		{
			Block:  `out: 3; false || echo foobar`,
			Stdout: "3\nfalsefoobar\n",
		},
		{
			Block:   `out: 4; false && echo foobar`,
			Stdout:  "4\nfalse",
			ExitNum: 1,
		},
		{
			Block:  `out: 4.1; true || true && echo foobar`,
			Stdout: "4.1\ntrue",
		},
		///
		{
			Block:   `out: 5; try { false; echo foobar }`,
			Stdout:  "5\nfalse",
			ExitNum: 1,
		},
		{
			Block:  `out: 6; try { true || echo foobar }`,
			Stdout: "6\ntrue",
		},
		{
			Block:  `out: 7; try { true && echo foobar }`,
			Stdout: "7\ntruefoobar\n",
		},
		{
			Block:  `out: 8; try { false || echo foobar }`,
			Stdout: "8\nfalsefoobar\n",
		},
		{
			Block:   `out: 9; try { false && echo foobar }`,
			Stdout:  "9\nfalse",
			ExitNum: 1,
		},
		{
			Block:  `out: 9.1; try { true || true; echo foobar }`,
			Stdout: "9.1\ntruefoobar\n",
		},
		{
			Block:  `out: 9.1; try { true || true && echo foobar }`,
			Stdout: "9.1\ntruefoobar\n",
		},
		///
		{
			Block:   `out: 10; trypipe { false; echo foobar }`,
			Stdout:  "10\nfalse",
			ExitNum: 1,
		},
		{
			Block:  `out: 11; trypipe { true || echo foobar }`,
			Stdout: "11\ntrue",
		},
		{
			Block:  `out: 12; trypipe { true && echo foobar }`,
			Stdout: "12\ntruefoobar\n",
		},
		{
			Block:  `out: 13; trypipe { false || echo foobar }`,
			Stdout: "13\nfalsefoobar\n",
		},
		{
			Block:   `out: 14; trypipe { false && echo foobar }`,
			Stdout:  "14\nfalse",
			ExitNum: 1,
		},
		{
			Block:  `out: 14.1; trypipe { true || true ; echo foobar }`,
			Stdout: "14.1\ntruefoobar\n",
		},
		{
			Block:  `out: 14.1; trypipe { true || true && echo foobar }`,
			Stdout: "14.1\ntruefoobar\n",
		},
	}

	test.RunMurexTests(tests, t)
}
