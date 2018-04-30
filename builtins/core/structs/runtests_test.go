package structs

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/typemgmt"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
)

type test struct {
	block  string
	result bool
}

func runTests(tests []test, t *testing.T) {
	defaults.Defaults(proc.InitConf, false)
	proc.InitEnv()

	for i := range tests {
		stdout := streams.NewStdin()
		stderr := streams.NewStdin()

		exitNum, err := lang.RunBlockShellNamespace([]rune(tests[i].block), nil, stdout, stderr)
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
		if boolean != tests[i].result {
			t.Error(tests[i].block, "returned", boolean)
		}
	}
}
