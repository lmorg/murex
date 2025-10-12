package lang

/*
	This test library relates to the testing framework within the murex
	language itself rather than Go's test framework within the murex project.

	The naming convention here is basically the inverse of Go's test naming
	convention. ie Go source files will be named "test_unit.go" (because
	calling it unit_test.go would mean it's a Go test rather than murex test)
	and the code is named UnitTestPlans (etc) rather than TestUnitPlans (etc)
	because the latter might suggest they would be used by `go test`. This
	naming convention is a little counterintuitive but it at least avoids
	naming conflicts with `go test`.
*/

import (
	"github.com/lmorg/murex/utils"
)

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

	tests.AddResult(&TestProperties{Name: name}, p, TestError, "Test named but there is no test defined")
	return

compare:

	var failed bool
	test := tests.test[i]
	test.HasRan = true
	tests.mutex.Unlock()

	// read stdout
	stdout, err := test.out.stdio.ReadAll()
	if err != nil {
		failed = true
		tests.AddResult(test, p, TestError, tMsgReadErr("stdout", name, err))
	}
	stdout = utils.CrLfTrim(stdout)

	// read stderr
	stderr, err := test.err.stdio.ReadAll()
	if err != nil {
		failed = true
		tests.AddResult(test, p, TestError, tMsgReadErr("stderr", name, err))
	}
	stderr = utils.CrLfTrim(stderr)

	// compare stdout
	if len(test.out.Block) != 0 {
		testBlock(test, p, test.out.Block, stdout, test.out.stdio.GetDataType(), "StdoutBlock", &failed)
	}

	if test.out.Regexp != nil {
		if test.out.Regexp.Match(stdout) {
			tests.AddResult(test, p, TestInfo, tMsgRegexMatch("StdoutRegex"))
		} else {
			failed = true
			tests.AddResult(test, p, TestFailed, tMsgRegexMismatch("StdoutRegex", stdout))
		}
	}

	// compare stderr
	if len(test.err.Block) != 0 {
		testBlock(test, p, test.err.Block, stderr, test.err.stdio.GetDataType(), "StderrBlock", &failed)
	}

	if test.err.Regexp != nil {
		if test.err.Regexp.Match(stderr) {
			tests.AddResult(test, p, TestInfo, tMsgRegexMatch("StderrRegex"))
		} else {
			failed = true
			tests.AddResult(test, p, TestFailed, tMsgRegexMismatch("StderrRegex", stderr))
		}
	}

	// test exit number
	if test.exitNum != *test.exitNumPtr {
		failed = true
		tests.AddResult(test, p, TestFailed, tMsgExitNumMismatch(test.exitNum, *test.exitNumPtr))

	} else {
		tests.AddResult(test, p, TestInfo, tMsgExitNumMatch())
	}

	// if not failed, log a success result
	if !failed {
		tests.AddResult(test, p, TestPassed, tMsgPassed())
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

func testIsMap(b []byte, dt string, property string) (TestStatus, string) {
	fork := ShellProcess.Fork(F_CREATE_STDIN)
	fork.Stdin.SetDataType(dt)
	_, err := fork.Stdin.Write(b)
	if err != nil {
		return TestError, tMsgWriteErr(property, err)
	}

	v, err := UnmarshalData(fork.Process, dt)
	if err != nil {
		return TestFailed, tMsgUnmarshalErr(property, dt, err)
	}

	switch v.(type) {
	case map[string]string, map[string]any, map[any]string, map[any]any:
		return TestPassed, tMsgDataFormatValid(property, dt, v)

	default:
		return TestFailed, tMsgDataFormatInvalid(property, dt, v)
	}
}

func testIsArray(b []byte, dt string, property string) (TestStatus, string) {
	fork := ShellProcess.Fork(F_CREATE_STDIN)
	fork.Stdin.SetDataType(dt)
	_, err := fork.Stdin.Write(b)
	if err != nil {
		return TestError, tMsgWriteErr(property, err)
	}

	v, err := UnmarshalData(fork.Process, dt)
	if err != nil {
		return TestFailed, tMsgUnmarshalErr(property, dt, err)
	}

	switch v.(type) {
	case []string, []any:
		return TestPassed, tMsgDataFormatValid(property, dt, v)

	default:
		return TestFailed, tMsgDataFormatInvalid(property, dt, v)
	}
}

func testIsGreaterThanOrEqualTo(b []byte, dt string, property string, comparison int) (TestStatus, string) {
	fork := ShellProcess.Fork(F_CREATE_STDIN)
	fork.Stdin.SetDataType(dt)
	_, err := fork.Stdin.Write(b)
	if err != nil {
		return TestError, tMsgWriteErr(property, err)
	}

	v, err := UnmarshalData(fork.Process, dt)
	if err != nil {
		return TestFailed, tMsgUnmarshalErr(property, dt, err)
	}

	var l int
	switch t := v.(type) {
	case []string:
		l = len(t)
	case []float64:
		l = len(t)
	case []int:
		l = len(t)
	case []bool:
		l = len(t)
	case []any:
		l = len(t)

	case [][]string:
		l = len(t)
	case [][]any:
		l = len(t)

	case map[string]string:
		l = len(t)
	case map[string]any:
		l = len(t)
	case map[any]any:
		l = len(t)

	default:
		return TestFailed, tMsgDataFormatInvalid(property, dt, v)
	}

	if l >= comparison {
		return TestPassed, tMsgGtEqMatch(property, l, comparison)
	}
	return TestFailed, tMsgGtEqFail(property, l, comparison)
}
