package test_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

// TestMurex proves the murex's scripting wrapper for Go's test framework works.
// Please note this shouldn't be confused with the murex scripting language's inbuilt testing framework!
func TestMurex(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  "true",
			Stdout: "true",
		},
		{
			Block:   "false",
			Stdout:  "false",
			ExitNum: 1,
		},
	}

	test.RunMurexTests(tests, t)
}
