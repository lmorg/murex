package system_test

import (
	"runtime"
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestOs(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `os`,
			Stdout: runtime.GOOS,
		},
		{
			Block:   `os bob`,
			Stdout:  "false",
			ExitNum: 1,
		},
		{
			Block:  `os ` + runtime.GOOS,
			Stdout: "true",
		},
	}

	test.RunMurexTests(tests, t)
}
