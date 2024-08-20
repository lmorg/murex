//go:build windows
// +build windows

package system_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestOsUnix(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `os windows`,
			Stdout: "true",
		},
		{
			Block:   `os posix`,
			Stdout:  "false",
			ExitNum: 1,
		},
	}

	test.RunMurexTests(tests, t)
}
