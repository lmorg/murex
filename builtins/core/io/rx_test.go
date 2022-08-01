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
			Block:   "rx: README$",
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
	}
	test.RunMurexTestsRx(tests, t)
}
