//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package system_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestOsWindows(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `os windows`,
			Stdout: "true",
		},
		{
			Block:  `os posix`,
			Stdout: "false",
		},
	}

	test.RunMurexTests(tests, t)
}
