package io

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestLsG(t *testing.T) {
	tests := []test.MurexTest{
		// g
		{
			Block:  "g: README*",
			Stdout: "README.md",
		},
		{
			Block:   "g: README",
			Stderr:  "Error",
			ExitNum: 1,
		},
		// !g
		{
			Block:   "!g: *",
			Stderr:  "Error",
			ExitNum: 1,
		},
		{
			Block:  "!g: README",
			Stdout: "README.md",
		},
		// ->g
		{
			Block:  "g: R* -> g: *.md",
			Stdout: "README.md",
		},
		{
			Block:   "g: R* -> g: *.doesntexist",
			Stderr:  "Error",
			ExitNum: 1,
		},
		{
			Block:   "g: *doesntexist -> g: *.md",
			Stderr:  "Error",
			ExitNum: 1,
		},
		// ->!g
		{
			Block:   "g: R* -> !g: *.md",
			Stderr:  "Error",
			ExitNum: 1,
		},
		{
			Block:  "g: R* -> !g: *.doesntexist",
			Stdout: "README.md",
		},
		{
			Block:   "g: *doesntexist -> !g: *.md",
			Stderr:  "Error",
			ExitNum: 1,
		},
	}
	test.RunMurexTestsRx(tests, t)
}
