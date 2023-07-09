package lang_test

import (
	"fmt"
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

/*
	TestMurexTestingFramework tests murex's testing framework using Go's testing
	framework (confused? Essentially murex shell scripts have a testing framework
	leveredged via the `test` builtin. This can be used for testing and debugging
	murex shell script. The concept behind them is that you place all of the test
	code within the normal shell script and they sit there idle while murex is
	running. However the moment you enable the testing flag (via `config`) those
	test builtins start writing their results to the test report for your review)

	This Go source file tests that murex's test builtins and report functions
	work by testing the Go code that resides behind them.
*/

func TestMurexTestDefine(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: fmt.Sprintf(`
					pipe: %s%d

					function test.%s.%d {
						test: <null> config enable auto-report !verbose
						config: set test report-format json
						config: set test report-pipe %s%d

						test: define bob {
							"StdoutMatch": "%s.%d\n"
						}

						out: <test_bob> "%s.%d"
					}

					bg <err> {
						<%s%d> 
					}

					test.%s.%d
					sleep: 2
					!pipe: %s%d
				`,
				t.Name(), 0,
				t.Name(), 0,
				t.Name(), 0,
				t.Name(), 0,
				t.Name(), 0,
				t.Name(), 0,
				t.Name(), 0,
				t.Name(), 0,
			),
			Stdout: fmt.Sprintf("%s.[0-9]+", t.Name()),
			Stderr: "PASSED",
		},
		{
			Block: fmt.Sprintf(`
					pipe: %s%d

					function test.%s.%d {
						test: <null> config enable auto-report !verbose
						config: set test report-format json
						config: set test report-pipe %s%d

						test: define bob {
							"StderrMatch": "%s.%d\n",
							"ExitNum": 1
						}

						err: <test_bob> "%s.%d"
					}

					bg {
						<%s%d> 
					}

					test.%s.%d
					sleep: 2
					!pipe: %s%d
				`,
				t.Name(), 1,
				t.Name(), 1,
				t.Name(), 1,
				t.Name(), 1,
				t.Name(), 1,
				t.Name(), 1,
				t.Name(), 1,
				t.Name(), 1,
			),
			Stderr: fmt.Sprintf("%s.[0-9]+", t.Name()),
			Stdout: "PASSED",
		},
	}

	test.RunMurexTestsRx(tests, t)
}

func TestMurexTestUnit(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: fmt.Sprintf(`
					test: unit function test.%s.%d {
						"StdoutMatch": "%s.%d\n"
					}

					function test.%s.%d {
						out: "%s.%d"
					}

					source { test: run test.%s.%d }
				`,
				t.Name(), 0,
				t.Name(), 0,
				t.Name(), 0,
				t.Name(), 0,
				t.Name(), 0,
			),
			Stdout: "PASSED",
		},
		{
			Block: fmt.Sprintf(`
					test: unit function test.%s.%d {
						"StdoutRegex": "%s.%d"
					}

					function test.%s.%d {
						out: "%s.%d"
					}

					source { test: run test.%s.%d }
				`,
				t.Name(), 1,
				t.Name(), 1,
				t.Name(), 1,
				t.Name(), 1,
				t.Name(), 1,
			),
			Stdout: "PASSED",
		},
		{
			Block: fmt.Sprintf(`
					test: unit function test.%s.%d {
						"StderrMatch": "%s.%d\n",
						"ExitNum": 1
					}

					function test.%s.%d {
						err: "%s.%d"
					}

					source { test: run test.%s.%d }
				`,
				t.Name(), 2,
				t.Name(), 2,
				t.Name(), 2,
				t.Name(), 2,
				t.Name(), 2,
			),
			Stdout: "PASSED",
		},
		{
			Block: fmt.Sprintf(`
					test: unit function test.%s.%d {
						"StderrRegex": "%s.%d",
						"ExitNum": 1
					}

					function test.%s.%d {
						err: "%s.%d"
					}

					source { test: run test.%s.%d }
				`,
				t.Name(), 3,
				t.Name(), 3,
				t.Name(), 3,
				t.Name(), 3,
				t.Name(), 3,
			),
			Stdout: "PASSED",
		},
	}

	test.RunMurexTestsRx(tests, t)
}
