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
			Stdout: "\n",
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
