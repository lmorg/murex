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
			Block:   `os windows`,
			Stdout:  "false",
			ExitNum: 1,
		},
		{
			Block:  `os posix`,
			Stdout: "true",
		},
	}

	test.RunMurexTests(tests, t)
}
