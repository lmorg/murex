package io

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestLsF(t *testing.T) {
	tests := []test.MurexTest{
		// f
		{
			Block:  "f: +f",
			Stdout: "README.md",
		},
	}
	test.RunMurexTestsRx(tests, t)
}
