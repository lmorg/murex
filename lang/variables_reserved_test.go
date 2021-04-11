package lang_test

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestVarSelf(t *testing.T) {
	tests := []test.MurexTest{
		// TTY
		{
			Block: `
				function TestVarSelf {
					$SELF -> [TTY]
				}
				TestVarSelf
			`,
			Stdout: "false\n",
		},

		// Method
		{
			Block: `
				function TestVarSelf {
					$SELF -> [Method]
				}
				TestVarSelf
			`,
			Stdout: "false\n",
		},
		{
			Block: `
				function TestVarSelf {
					$SELF -> [Method]
				}
				out foobar -> TestVarSelf
			`,
			Stdout: "true\n",
		},

		// Not
		{
			Block: `
				function TestVarSelf {
					$SELF -> [Not]
				}
				TestVarSelf
			`,
			Stdout: "false\n",
		},
		{
			Block: `
				function !TestVarSelf {
					$SELF -> [Not]
				}
				!TestVarSelf
			`,
			Stdout: "true\n",
		},

		// Background
		{
			Block: `
				function TestVarSelf {
					$SELF -> [Background]
				}
				TestVarSelf
			`,
			Stdout: "false\n",
		},
		{
			Block: `
				function TestVarSelf {
					$SELF -> [Background]
				}
				bg { TestVarSelf }
				sleep 1
			`,
			Stdout: "true\n",
		},
	}

	test.RunMurexTests(tests, t)
}
