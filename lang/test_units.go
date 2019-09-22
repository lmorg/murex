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
	"errors"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/types"
)

// UnitTests is a class for all things murex unit tests
type UnitTests struct {
	units []*unitTest
	mutex sync.Mutex
}

type unitTest struct {
	Function string // if private it should contain path module path
	FileRef  *ref.File
	TestPlan *UnitTestPlan
}

// Add a new unit test
func (ut *UnitTests) Add(function string, test *UnitTestPlan, fileRef *ref.File) {
	newUnitTest := &unitTest{
		Function: function,
		TestPlan: test,
		FileRef:  fileRef,
	}

	ut.mutex.Lock()
	ut.units = append(ut.units, newUnitTest)
	ut.mutex.Unlock()
}

const testName = "unit test"

const (
	// UnitTestAutocomplete is the psuedo-module name for autocomplete blocks
	UnitTestAutocomplete = "(autocomplete)"

	// UnitTestOpen is the psuedo-module name for open handler blocks
	UnitTestOpen = "(open)"

	// UnitTestEvent is the psuedo-module name for event blocks
	UnitTestEvent = "(event)"
)

// Run all unit tests against a specific murex function
func (ut *UnitTests) Run(p *Process, function string) bool {
	ut.mutex.Lock()
	utCopy := make([]*unitTest, len(ut.units))
	copy(utCopy, ut.units)
	ut.mutex.Unlock()

	var (
		passed = true
		exists bool
	)

	for i := range utCopy {
		if function == "*" || utCopy[i].Function == function {
			passed = passed && runTest(p.Tests.Results, utCopy[i].FileRef, utCopy[i].TestPlan, utCopy[i].Function)
			exists = true
		}
	}

	if !exists {
		passed = false
		p.Tests.Results.Add(&TestResult{
			Exec:     function,
			TestName: testName,
			Status:   TestError,
			Message:  fmt.Sprintf("No unit tests exist for: `%s`", function),
		})
	}

	v, err := p.Config.Get("test", "auto-report", "bool")
	if err != nil {
		v = true
	}
	if v.(bool) {
		p.Tests.WriteResults(p.Config, p.Stdout)
	}

	return passed
}

// Dump the defined unit tests in a JSONable structure
func (ut *UnitTests) Dump() interface{} {
	ut.mutex.Lock()
	dump := ut.units
	ut.mutex.Unlock()

	return dump
}

// UnitTestPlan is defined via JSON and specifies an individual test plan
type UnitTestPlan struct {
	Parameters    []string
	Stdin         string
	StdoutMatch   string
	StderrMatch   string
	StdinType     string
	StdoutType    string
	StderrType    string
	StdoutRegex   string
	StderrRegex   string
	StdoutBlock   string
	StderrBlock   string
	StdoutIsArray bool
	StderrIsArray bool
	StdoutIsMap   bool
	StderrIsMap   bool
	ExitNum       int
	PreBlock      string
	PostBlock     string
}

func utAddReport(results *TestResults, fileRef *ref.File, plan *UnitTestPlan, function string, status TestStatus, message string) {
	results.Add(&TestResult{
		ColNumber:  fileRef.Column,
		LineNumber: fileRef.Line,
		Exec:       function,
		Params:     plan.Parameters,
		TestName:   testName,
		Status:     status,
		Message:    message,
	})
}

func runTest(results *TestResults, fileRef *ref.File, plan *UnitTestPlan, function string) bool {
	var (
		preExitNum, testExitNum, postExitNum int
		preForkErr, testForkErr, postForkErr error
		stdoutType, stderrType               string

		fStdin int
		passed = true
	)

	addReport := func(status TestStatus, message string) {
		utAddReport(results, fileRef, plan, function, status, message)
	}

	if len(plan.Stdin) == 0 {
		fStdin = F_NO_STDIN
	} else {
		fStdin = F_CREATE_STDIN
	}

	fork := ShellProcess.Fork(fStdin | F_CREATE_STDOUT | F_CREATE_STDERR | F_FUNCTION)
	fork.Parameters.Params = plan.Parameters

	if len(plan.Stdin) > 0 {
		if plan.StdinType == "" {
			plan.StdinType = types.Generic
		}
		fork.Stdin.SetDataType(plan.StdinType)
		_, err := fork.Stdin.Write([]byte(plan.Stdin))
		if err != nil {
			fmt.Println(err)
			return false
		}
	}

	// run any initializing code...if defined
	if len(plan.PreBlock) > 0 {
		preFork := ShellProcess.Fork(F_FUNCTION | F_NEW_MODULE | F_CREATE_STDOUT | F_CREATE_STDERR)
		preFork.FileRef = fileRef
		preFork.Name = "(unit test PreBlock)"
		preExitNum, preForkErr = preFork.Execute([]rune(plan.PreBlock))

		if preForkErr != nil {
			passed = false
			addReport(TestError, tMsgCompileErr("PreBlock", preForkErr))
		}

		if preExitNum != 0 {
			addReport(TestInfo, tMsgNoneZeroExit("PreBlock", preExitNum))
		}

		utReadAllOut(preFork.Stdout, results, plan, fileRef, "PreBlock", function, &passed)
		utReadAllErr(preFork.Stderr, results, plan, fileRef, "PreBlock", function, &passed)
	}

	// run function
	testExitNum, testForkErr = runFunction(function, plan.Stdin != "", fork)
	if testForkErr != nil {
		addReport(TestError, tMsgCompileErr(function, testForkErr))
		return false
	}

	// run any clear down code...if defined
	if len(plan.PostBlock) > 0 {
		postFork := ShellProcess.Fork(F_FUNCTION | F_NEW_MODULE | F_CREATE_STDOUT | F_CREATE_STDERR)
		postFork.Name = "(unit test PostBlock)"
		postFork.FileRef = fileRef
		postExitNum, postForkErr = postFork.Execute([]rune(plan.PostBlock))

		if postForkErr != nil {
			passed = false
			addReport(TestError, tMsgCompileErr("PostBlock", postForkErr))
		}

		if postExitNum != 0 {
			addReport(TestInfo, tMsgNoneZeroExit("PostBlock", preExitNum))
		}

		utReadAllOut(postFork.Stdout, results, plan, fileRef, "PostBlock", function, &passed)
		utReadAllErr(postFork.Stderr, results, plan, fileRef, "PostBlock", function, &passed)
	}

	// stdout

	stdout, err := fork.Stdout.ReadAll()
	if err != nil {
		addReport(TestFailed, tMsgReadErr("stdout", function, err))
		return false
	}
	stdoutType = fork.Stdout.GetDataType()

	// stderr

	stderr, err := fork.Stderr.ReadAll()
	if err != nil {
		addReport(TestFailed, tMsgReadErr("stderr", function, err))
		return false
	}
	stderrType = fork.Stderr.GetDataType()

	// test exit number

	if testExitNum == plan.ExitNum {
		addReport(TestInfo, tMsgExitNumMatch())
	} else {
		passed = false
		addReport(TestFailed, tMsgExitNumMismatch(plan.ExitNum, testExitNum))
	}

	// test stdout stream

	if plan.StdoutIsArray {
		status, message := testIsArray(stdout, stdoutType, "StdoutIsArray")
		addReport(status, message)
		if status != TestPassed {
			passed = false
		}
	}

	if plan.StdoutIsMap {
		status, message := testIsMap(stdout, stdoutType, "StdoutIsMap")
		addReport(status, message)
		if status != TestPassed {
			passed = false
		}
	}

	if plan.StdoutMatch != "" {
		if string(stdout) == plan.StdoutMatch {
			addReport(TestInfo, tMsgStringMatch("StdoutMatch"))
		} else {
			passed = false
			addReport(TestFailed, tMsgStringMismatch("StdoutMatch", stdout))
		}
	}

	if plan.StdoutRegex != "" {
		rx, err := regexp.Compile(plan.StdoutRegex)
		switch {
		case err != nil:
			passed = false
			addReport(TestError, tMsgRegexCompileErr("StdoutRegex", err))

		case !rx.Match(stdout):
			passed = false
			addReport(TestFailed, tMsgRegexMismatch("StdoutRegex", stdout))

		default:
			addReport(TestInfo, tMsgRegexMatch("StdoutRegex"))
		}
	}

	if plan.StdoutBlock != "" {
		utBlock(plan, fileRef, []rune(plan.StdoutBlock), stdout, stdoutType, "StdoutBlock", function, results, &passed)
	}

	if plan.StdoutType != "" {
		if stdoutType == plan.StdoutType {
			addReport(TestInfo, tMsgDataTypeMatch("stdout"))
		} else {
			passed = false
			addReport(TestFailed, tMsgDataTypeMismatch("stdout", stdoutType))
		}
	}

	// test stderr stream

	if plan.StderrIsArray {
		status, message := testIsArray(stderr, stderrType, "StderrIsArray")
		addReport(status, message)
		if status != TestPassed {
			passed = false
		}
	}

	if plan.StderrIsMap {
		status, message := testIsMap(stderr, stderrType, "StderrIsMap")
		addReport(status, message)
		if status != TestPassed {
			passed = false
		}
	}

	if plan.StderrMatch != "" {
		if string(stderr) == plan.StderrMatch {
			addReport(TestInfo, tMsgStringMatch("StderrMatch"))
		} else {
			passed = false
			addReport(TestFailed, tMsgStringMismatch("StderrMatch", stderr))
		}
	}

	if plan.StderrRegex != "" {
		rx, err := regexp.Compile(plan.StderrRegex)
		switch {
		case err != nil:
			passed = false
			addReport(TestError, tMsgRegexCompileErr("StderrRegex", err))

		case !rx.Match(stderr):
			passed = false
			addReport(TestFailed, tMsgRegexMismatch("StderrRegex", stdout))

		default:
			addReport(TestInfo, tMsgRegexMatch("StderrRegex"))
		}
	}

	if plan.StderrBlock != "" {
		utBlock(plan, fileRef, []rune(plan.StderrBlock), stderr, stderrType, "StderrBlock", function, results, &passed)
	}

	if plan.StderrType != "" {
		if stderrType == plan.StderrType {
			addReport(TestInfo, tMsgDataTypeMatch("stdout"))
		} else {
			passed = false
			addReport(TestFailed, tMsgDataTypeMismatch("stderr", stderrType))
		}
	}

	// lastly, a passed message if no errors

	if passed {
		addReport(TestPassed, tMsgPassed())
	}

	return passed
}

func runFunction(function string, isMethod bool, fork *Fork) (int, error) {
	fork.IsMethod = isMethod

	if function[0] == '/' {
		function = function[1:]
	}

	if strings.Contains(function, "/") {
		return altFunc(function, fork)
	}

	fork.Name = function

	if !MxFunctions.Exists(function) {
		return 0, errors.New("Function does not exist")
	}

	block, err := MxFunctions.Block(function)
	if err != nil {
		return 0, err
	}

	return fork.Execute(block)
}

func altFunc(path string, fork *Fork) (int, error) {
	split := strings.Split(path, "/")

	switch split[0] {
	case UnitTestAutocomplete:
		return runAutocomplete(path, split, fork)
	case UnitTestOpen:
		return runOpen(path, split, fork)
	case UnitTestEven:
		return runEvent(path, split, fork)
	default:
		return runPrivate(path, split, fork)
	}
}

func runAutocomplete(path string, split []string, fork *Fork) (int, error) {
	return 0, errors.New("Not currently supported")
	//autocomplete.MatchFlags()
}

func runOpen(path string, split []string, fork *Fork) (int, error) {
	return 0, errors.New("TODO: not currently supported")
}

func runEvent(path string, split []string, fork *Fork) (int, error) {
	return 0, errors.New("TODO: not currently supported")
}

func runPrivate(path string, split []string, fork *Fork) (int, error) {
	if len(split) < 2 {
		return 0, fmt.Errorf("Invalid module and private function path: `%s`", path)
	}

	function := split[len(split)-1]
	module := strings.Join(split[:len(split)-1], "/")

	fork.Name = function

	if !PrivateFunctions.Exists(function, module) {
		return 0, fmt.Errorf("Private (%s) does not exist or module name (%s) is wrong", function, module)
	}

	block, err := PrivateFunctions.Block(function, module)
	if err != nil {
		return 0, err
	}

	return fork.Execute(block)
}
