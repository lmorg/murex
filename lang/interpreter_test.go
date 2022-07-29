package lang_test

import (
	"testing"

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
			Block:   `out "foo" || out "bar"`,
			Stdout:  "foo\n",
		},
		{
			Block:  `err "foo" || out "bar"`,
			Stdout: "bar\n",
			Stderr: "foo\n",
		},
		{
			Block:   `out "foo" || err "bar"`,
			Stdout:  "foo\n",
		},
		{
			Block:   `err "foo" || err "bar"`,
			Stderr:  "foo\nbar\n",
			ExitNum: 1,
		},
		/////
		{
			Block:   `out "foo"||out "bar"`,
			Stdout:  "foo\n",
		},
		{
			Block:  `err "foo"||out "bar"`,
			Stdout: "bar\n",
			Stderr: "foo\n",
		},
		{
			Block:   `out "foo"||err "bar"`,
			Stdout:  "foo\n",
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
			Stdout: "0\nfalse\nfoobar\n",
		},
		{
			Block:  `out: 1; true || echo foobar`,
			Stdout: "1\ntrue\n",
		},
		{
			Block:  `out: 1.1; true || true || echo foobar`,
			Stdout: "1.1\ntrue\n",
		},
		{
			Block:  `out: 1.2; true || false || echo foobar`,
			Stdout: "1.2\ntrue\n",
		},
		{
			Block:  `out: 1.3; false || true || echo foobar`,
			Stdout: "1.3\nfalse\ntrue\n",
		},
		{
			Block:  `out: 2; true && echo foobar`,
			Stdout: "2\ntrue\nfoobar\n",
		},
		{
			Block:   `out: 2.1; false && false && echo foobar`,
			Stdout:  "2.1\nfalse\n",
			ExitNum: 1,
		},
		{
			Block:   `out: 2.2; true && false && echo foobar`,
			Stdout:  "2.2\ntrue\nfalse\n",
			ExitNum: 1,
		},
		{
			Block:   `out: 2.3; false && true && echo foobar`,
			Stdout:  "2.3\nfalse\n",
			ExitNum: 1,
		},
		{
			Block:  `out: 3; false || echo foobar`,
			Stdout: "3\nfalse\nfoobar\n",
		},
		{
			Block:   `out: 4; false && echo foobar`,
			Stdout:  "4\nfalse\n",
			ExitNum: 1,
		},
		{
			Block:  `out: 4.1; true || true && echo foobar`,
			Stdout: "4.1\ntrue\nfoobar\n", // this seems wrong but if falls in line with how bash works
		},
		///
		{
			Block:   `out: 5; try { false; echo foobar }`,
			Stdout:  "5\nfalse\n",
			ExitNum: 1,
		},
		{
			Block:  `out: 6; try { true || echo foobar }`,
			Stdout: "6\ntrue\n",
		},
		{
			Block:  `out: 7; try { true && echo foobar }`,
			Stdout: "7\ntrue\nfoobar\n",
		},
		{
			Block:  `out: 8; try { false || echo foobar }`,
			Stdout: "8\nfalse\nfoobar\n",
		},
		{
			Block:   `out: 9; try { false && echo foobar }`,
			Stdout:  "9\nfalse\n",
			ExitNum: 1,
		},
		///
		{
			Block:   `out: 10; trypipe { false; echo foobar }`,
			Stdout:  "10\nfalse\n",
			ExitNum: 1,
		},
		{
			Block:  `out: 11; trypipe { true || echo foobar }`,
			Stdout: "11\ntrue\n",
		},
		{
			Block:  `out: 12; trypipe { true && echo foobar }`,
			Stdout: "12\ntrue\nfoobar\n",
		},
		{
			Block:  `out: 13; trypipe { false || echo foobar }`,
			Stdout: "13\nfalse\nfoobar\n",
		},
		{
			Block:   `out: 14; trypipe { false && echo foobar }`,
			Stdout:  "14\nfalse\n",
			ExitNum: 1,
		},
	}

	test.RunMurexTests(tests, t)
}
