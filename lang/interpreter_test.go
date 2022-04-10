package lang_test

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestOperatorLogicAndNormal(t *testing.T) {
	tests := []test.MurexTest{
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
	}

	test.RunMurexTests(tests, t)
}

func TestOperatorLogicOrNormal(t *testing.T) {
	tests := []test.MurexTest{
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
	}

	test.RunMurexTests(tests, t)
}
