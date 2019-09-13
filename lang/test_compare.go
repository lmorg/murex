package lang

/*
	This test library relates to the testing framework within the
	murex language itself rather than Go's test framework within
	the murex project.
*/

import (
	"fmt"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

const testPassedMessage = "All test conditions were met"

// Compare is the method which actually runs the individual test cases
// to see if they pass or fail.
func (tests *Tests) Compare(name string, p *Process) {
	tests.mutex.Lock()

	var i int
	for ; i < len(tests.test); i++ {
		if tests.test[i].Name == name {
			goto compare
		}
	}

	tests.mutex.Unlock()

	tests.AddResult(&TestProperties{Name: name}, p, TestError, "Test named but there is no test defined.")
	return

compare:

	var failed, verbose bool
	test := tests.test[i]
	test.HasRan = true
	tests.mutex.Unlock()

	v, err := p.Config.Get("test", "verbose", types.Boolean)
	if err == nil && v.(bool) {
		verbose = true
	}

	// read stdout
	stdout, err := test.out.stdio.ReadAll()
	if err != nil {
		failed = true
		tests.AddResult(test, p, TestError, "Cannot read from stdout")
	}
	stdout = utils.CrLfTrim(stdout)

	// read stderr
	stderr, err := test.err.stdio.ReadAll()
	if err != nil {
		failed = true
		tests.AddResult(test, p, TestError, "Cannot read from stderr")
	}
	stderr = utils.CrLfTrim(stderr)

	// compare stdout
	if len(test.out.Block) > 0 {
		b, bErr, err := test.out.RunBlock(p, test.out.Block, stdout)
		if err != nil {
			failed = true
			tests.AddResult(test, p, TestError, err.Error())

		} else if string(b) != string(stdout) {
			failed = true
			tests.AddResult(test, p, TestFailed,
				fmt.Sprintf("stdout: got '%s' returned '%s'",
					stdout, b))

		} else if verbose {
			tests.AddResult(test, p, TestPassed, fmt.Sprintf("stdout: block passed '%s'", stdout))
		}

		if verbose {
			tests.AddResult(test, p, TestInfo, fmt.Sprintf("stdout: stderr returned from block '%s'", bErr))
		}

	} else if test.out.Regexp != nil {
		if !test.out.Regexp.Match(stdout) {
			failed = true
			tests.AddResult(test, p, TestFailed,
				fmt.Sprintf("stdout: regexp did not match '%s'",
					stdout))

		} else if verbose {
			tests.AddResult(test, p, TestPassed, fmt.Sprintf("stdout: regexp matched '%s'", stdout))
		}
	}

	// compare stderr
	if len(test.err.Block) > 0 {
		b, bErr, err := test.err.RunBlock(p, test.err.Block, stderr)
		if err != nil {
			failed = true
			tests.AddResult(test, p, TestError, err.Error())

		} else if string(b) != string(stderr) {
			failed = true
			tests.AddResult(test, p, TestFailed,
				fmt.Sprintf("stderr: got '%s' returned '%s'",
					stderr, b))

		} else if verbose {
			tests.AddResult(test, p, TestPassed, fmt.Sprintf("stderr: block passed '%s'", stderr))
		}

		if verbose {
			tests.AddResult(test, p, TestInfo, fmt.Sprintf("stderr: stderr returned from block '%s'", bErr))
		}

	} else if test.err.Regexp != nil {
		if !test.err.Regexp.Match(stderr) {
			failed = true
			tests.AddResult(test, p, TestFailed,
				fmt.Sprintf("stderr: regexp did not match '%s'.",
					stderr))

		} else if verbose {
			tests.AddResult(test, p, TestPassed, fmt.Sprintf("stderr: regexp matched '%s'", stderr))
		}
	}

	// test exit number
	if test.exitNum != *test.exitNumPtr {
		failed = true
		tests.AddResult(test, p, TestFailed,
			fmt.Sprintf("exit number: wanted %d got %d",
				test.exitNum, *test.exitNumPtr))

	} else if verbose {
		tests.AddResult(test, p, TestPassed, fmt.Sprintf("exit number: returned '%d'", test.exitNum))
	}

	// if not failed, log a success result
	if !failed {
		tests.AddResult(test, p, TestPassed, testPassedMessage)
	}
}

// ReportMissedTests is used so we have a result of tests that didn't run
func (tests *Tests) ReportMissedTests(p *Process) {
	tests.mutex.Lock()

	for _, test := range tests.test {
		if test.HasRan {
			continue
		}

		tests.AddResult(test, &Process{Config: p.Config}, TestMissed, "Test was defined but no function ran against that test pipe.")
	}

	tests.mutex.Unlock()
}
