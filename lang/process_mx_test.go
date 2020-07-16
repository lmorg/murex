package lang_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

// TestMxProcess is a function for testing parser, builtins and other behaviors
// which are defined in process.go. This might result in duplication of tests
// where such behavior is also tested in the builtin or other package, however
// that is acceptable because it allows different packages to be altered,
// refactored and even completely rewritten while still maintaining as much
// test coverage as possible.
func TestMxProcess(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `pipe: foobar
					bg { <foobar> }
					out: "Hello, world!" -> <foobar>
					!pipe: foobar`,
			Stdout: "Hello, world!\n",
		},

		{
			Block: `alias: foobar=out Hello, world!
					foobar`,
			Stdout: "Hello, world!\n",
		},

		{
			Block:  `@g out "Hello, world!"`,
			Stdout: "Hello, world!\n",
		},

		{
			Block: `global: foobar="Hello, world!"
					$foobar`,
			Stdout: "Hello, world!",
		},

		{
			Block:   `$`,
			Stdout:  "",
			Stderr:  "Error in `$` ( 1,1): Variable token, `$`, used without specifying variable name\n",
			ExitNum: 1,
		},

		{
			Block:   `$!`,
			Stdout:  "",
			Stderr:  "Error in `$!` ( 1,1): `!` is not a valid variable name\n",
			ExitNum: 1,
		},

		{
			Block: `global: json array = ([0, 1, 2, 3])
					$array[2]`,
			Stdout: "2\n",
		},

		{
			Block: `function test-func {
						out: "Hello, world!" 
					}
					test-func`,
			Stdout: "Hello, world!\n",
		},

		{
			Block: `private test-priv {
						out: "Hello, world!"
					}
					test-priv`,
			Stdout: "Hello, world!\n",
		},

		{
			Block: `func hello-world {
						test define example {
							"StdoutRegex": (^Hello world$)
						}
						out <test_example> "Hello world"
					}
					test config enable !auto-report
					hello-world`,
			Stdout: "Enabling test mode....\nDisabling auto-report....\nHello world\n",
		},
	}

	test.RunMurexTests(tests, t)
}
