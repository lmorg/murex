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
			Stdout: "false",
		},

		// Method
		{
			Block: `
				function TestVarSelf {
					$SELF -> [Method]
				}
				TestVarSelf
			`,
			Stdout: "false",
		},
		{
			Block: `
				function TestVarSelf {
					$SELF -> [Method]
				}
				out foobar -> TestVarSelf
			`,
			Stdout: "true",
		},

		// Not
		{
			Block: `
				function TestVarSelf {
					$SELF -> [Not]
				}
				TestVarSelf
			`,
			Stdout: "false",
		},
		{
			Block: `
				function !TestVarSelf {
					$SELF -> [Not]
				}
				!TestVarSelf
			`,
			Stdout: "true",
		},

		// Background
		{
			Block: `
				function TestVarSelf {
					$SELF -> [Background]
				}
				TestVarSelf
			`,
			Stdout: "false",
		},
		{
			Block: `
				function TestVarSelf {
					$SELF -> [Background]
				}
				bg { TestVarSelf }
				sleep 1
			`,
			Stdout: "true",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestVarArgs(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				function TestVarArgs {
					out @ARGS
				}
				TestVarArgs
			`,
			Stdout: "TestVarArgs\n",
		},
		{
			Block: `
				function TestVarArgs {
					out @ARGS
				}
				TestVarArgs 1 2 3
			`,
			Stdout: "TestVarArgs 1 2 3\n",
		},
		{
			Block: `
				function TestVarArgs {
					out @ARGS
				}
				TestVarArgs 1   2   3
			`,
			Stdout: "TestVarArgs 1 2 3\n",
		},
		{
			Block: `
				function TestVarArgs {
					out $ARGS
				}
				TestVarArgs 1   2   3
			`,
			Stdout: `["TestVarArgs","1","2","3"]` + "\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestVarParams(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				function TestVarParams {
					out @PARAMS
				}
				TestVarParams
			`,
			Stderr:  "Error in `out` ( 3,6): Array '@PARAMS' is empty\n",
			ExitNum: 1,
		},
		{
			Block: `
				function TestVarParams {
					out @PARAMS
				}
				TestVarParams 1 2 3
			`,
			Stdout: "1 2 3\n",
		},
		{
			Block: `
				function TestVarParams {
					out @PARAMS
				}
				TestVarParams 1   2   3
			`,
			Stdout: "1 2 3\n",
		},
		{
			Block: `
				function TestVarParams {
					out $PARAMS
				}
				TestVarParams 1   2   3
			`,
			Stdout: `["1","2","3"]` + "\n",
		},
	}

	test.RunMurexTests(tests, t)
}
