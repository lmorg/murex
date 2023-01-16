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
			Block: `pipe: TestMxProcess
					bg { <TestMxProcess> }
					out: "Hello, world!" -> <TestMxProcess>
					!pipe: TestMxProcess`,
			Stdout: "^Hello, world!\n$",
		},

		{
			Block: `alias: TestMxProcess=out Hello, world!
					TestMxProcess`,
			Stdout: "^Hello, world!\n$",
		},

		{
			Block: `global: TestMxProcess="Hello, world!"
					$TestMxProcess`,
			Stdout: "^Hello, world!$",
		},

		{
			Block: `global: json array = ([0, 1, 2, 3])
					$array[2]`,
			Stdout: "2",
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
			Block: `function hello-world {
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

	test.RunMurexTestsRx(tests, t)
}
