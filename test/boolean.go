package test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/typemgmt"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
)

// BooleanTest is a basic framework for each boolean test of murex code
type BooleanTest struct {
	Block  string
	Result bool
}

// RunBooleanTests runs through all the test cases for BooleanTest
func RunBooleanTests(tests []BooleanTest, t *testing.T) {
	defaults.Defaults(proc.InitConf, false)
	proc.InitEnv()

	for i := range tests {
		stdout := streams.NewStdin()
		stderr := streams.NewStdin()

		exitNum, err := lang.RunBlockShellConfigSpace([]rune(tests[i].Block), nil, stdout, stderr)
		if err != nil {
			t.Error(err.Error())
		}

		b, err := stderr.ReadAll()
		if err != nil {
			t.Error("unable to read from stderr: " + err.Error())
		}

		if len(b) > 0 {
			t.Error("stderr returned: " + string(b))
		}

		b, err = stdout.ReadAll()
		if err != nil {
			t.Error("unable to read from stdout: " + err.Error())
		}

		boolean := types.IsTrue(b, exitNum)
		if boolean != tests[i].Result {
			t.Error(tests[i].Block, "returned", boolean)
		}
	}
}
