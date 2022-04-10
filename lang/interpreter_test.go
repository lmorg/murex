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
			ExitNum: 1,
		},
		{
			Block:  `err "foo" || out "bar"`,
			Stdout: "bar\n",
			Stderr: "foo\n",
		},
		{
			Block:   `out "foo" || err "bar"`,
			Stdout:  "foo\n",
			ExitNum: 1,
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
			ExitNum: 1,
		},
		{
			Block:  `err "foo"||out "bar"`,
			Stdout: "bar\n",
			Stderr: "foo\n",
		},
		{
			Block:   `out "foo"||err "bar"`,
			Stdout:  "foo\n",
			ExitNum: 1,
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
	test.RunMurexTests(testsOr, t)
}

func TestOperatorLogicAndTryPipe(t *testing.T) {
	var tests []test.MurexTest
	for _, src := range testsAnd {
		newTest := src
		src.Block = "trypipe {" + src.Block + "}"
		tests = append(tests, newTest)
	}
	test.RunMurexTests(testsAnd, t)
}

func TestOperatorLogicOrTryPipe(t *testing.T) {
	var tests []test.MurexTest
	for _, src := range testsOr {
		newTest := src
		src.Block = "trypipe {" + src.Block + "}"
		tests = append(tests, newTest)
	}
	test.RunMurexTests(testsOr, t)
}
