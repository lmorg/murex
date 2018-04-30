package test

import (
	"testing"
)

// TestMurex prooves the murex test framework works
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
