package io

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestLsRx(t *testing.T) {
	tests := []test.MurexTest{
		// rx
		{
			Block:  "rx: R*ME",
			Stdout: "README.md",
		},
		{
			Block:   "rx: 'README$'",
			Stderr:  "Error",
			ExitNum: 1,
		},
		// !rx
		{
			Block:   "!rx: .*",
			Stderr:  "Error",
			ExitNum: 1,
		},
		{
			Block:  `!rx: (go|yaml)`,
			Stdout: "README.md",
		},
		// -> rx
		{
			Block:  "rx: R*ME -> rx: .*md",
			Stdout: "README.md",
		},
		{
			Block:   "rx: R*ME -> rx: .*doesntexist",
			Stderr:  "Error",
			ExitNum: 1,
		},
		{
			Block:   "rx: 'README$' ->  rx: .*md",
			Stderr:  "Error",
			ExitNum: 1,
		},
		// -> !rx
		{
			Block:   "rx: R*ME -> !rx: .*md",
			Stderr:  "Error",
			ExitNum: 1,
		},
		{
			Block:  "rx: R*ME -> !rx: .*doesntexist",
			Stdout: "README.md",
		},
		{
			Block:   "rx: 'README$' -> !rx: .*md",
			Stderr:  "Error",
			ExitNum: 1,
		},
	}
	test.RunMurexTestsRx(tests, t)
}
