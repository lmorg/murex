package test

import (
	"testing"

	"github.com/lmorg/murex/test/count"

	_ "github.com/lmorg/murex/builtins/core/typemgmt" // import boolean builtins
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

// BooleanTest is a basic framework for each boolean test of murex code.
// Please note this shouldn't be confused with the murex scripting language's inbuilt testing framework!
type BooleanTest struct {
	Block  string
	Result bool
}

// RunBooleanTests runs through all the test cases for BooleanTest.
// Please note this shouldn't be confused with the murex scripting language's inbuilt testing framework!
func RunBooleanTests(tests []BooleanTest, t *testing.T) {
	t.Helper()
	count.Tests(t, len(tests), "RunBooleanTests")

	defaults.Defaults(lang.InitConf, false)
	lang.InitEnv()

	for i := range tests {

		fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_CREATE_STDERR)
		fork.Name = "RunBooleanTests()"
		exitNum, err := fork.Execute([]rune(tests[i].Block))
		if err != nil {
			t.Error(err.Error())
		}

		b, err := fork.Stderr.ReadAll()
		if err != nil {
			t.Error("unable to read from stderr: " + err.Error())
		}

		if len(b) > 0 {
			t.Error("stderr returned: " + string(b))
		}

		b, err = fork.Stdout.ReadAll()
		if err != nil {
			t.Error("unable to read from stdout: " + err.Error())
		}

		boolean := types.IsTrue(b, exitNum)
		if boolean != tests[i].Result {
			t.Error(tests[i].Block, "returned", boolean)
		}
	}
}
