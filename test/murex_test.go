package test

import (
	"testing"
)

// TestMurex proves the murex's scripting wrapper for Go's test framework works.
// Please note this shouldn't be confused with the murex scripting language's inbuilt testing framework!
func TestMurex(t *testing.T) {
	tests := []MurexTest{
		{
			Block:  "true",
			Stdout: "true\n",
		},
		{
			Block:   "false",
			Stdout:  "false\n",
			ExitNum: 1,
		},
	}

	RunMurexTests(tests, t)
}
