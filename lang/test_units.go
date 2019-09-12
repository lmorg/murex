package lang

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/types"
)

// The naming convention here is basically the inverse of Go's test naming
// convention. ie Go source files will be named "test_unit.go" (because calling
// it unit_test.go would mean it's a Go test rather than murex test) and the
// code is named UnitTestPlans (etc) rather than TestUnitPlans (etc) because
// the latter might suggest they would be used by `go test`. This naming
// convention is a little counterintuitive but it at least avoids naming
// conflicts with `go test`.

// unitTests is a class for all things murex unit tests
type unitTests struct {
	units []*unitTest
	mutex sync.Mutex
}

type unitTest struct {
	Function string // if private it should contain path module path
	FileRef  *ref.File
	TestPlan *UnitTestPlan
}

// Add a new unit test
func (ut *unitTests) Add(function string, test *UnitTestPlan, fileRef *ref.File) {
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

// Run all unit tests against a specific murex function
func (ut *unitTests) Run(tests *Tests, function string) bool {
	ut.mutex.Lock()
	utCopy := make([]*unitTest, len(ut.units))
	copy(utCopy, ut.units)
	ut.mutex.Unlock()

	var (
		passed = true
		exists bool
	)

	for i := range utCopy {
		if utCopy[i].Function == function {
			passed = passed && runTest(tests.Results, utCopy[i].FileRef, utCopy[i].TestPlan, function)
			exists = true
		}
	}

	if exists {
		return passed
	}

	tests.Results.Add(&TestResult{
		Exec:     function,
		TestName: testName,
		Status:   TestFailed,
		Message:  fmt.Sprintf("No unit tests exist for: `%s`", function),
	})
	return false
}

func (ut *unitTests) Dump() interface{} {
	ut.mutex.Lock()
	dump := ut.units
	ut.mutex.Unlock()

	return dump
}

// UnitTestPlan is defined via JSON and specifies an individual test plan
type UnitTestPlan struct {
	Parameters []string
	Stdin      string
	Stdout     string
	Stderr     string
	StdinDT    string
	StdoutDT   string
	StderrDT   string
	StdoutRx   string // check this is the same as test define
	StderrRx   string // check this is the same as test define
	ExitNumber int    // check this is the same as test define
	PreBlock   string
	PostBlock  string
}

func runTest(results *TestResults, fileRef *ref.File, plan *UnitTestPlan, function string) bool {
	var (
		//testName                             = "unit test"
		preExitNum, testExitNum, postExitNum int
		preForkErr, testForkErr, postForkErr error
		fStdin                               int
		passed                               = true
		stdout, stderr                       string
	)

	if len(plan.Stdin) == 0 {
		fStdin = F_NO_STDIN
	} else {
		fStdin = F_CREATE_STDIN
	}

	fork := ShellProcess.Fork(fStdin | F_CREATE_STDOUT | F_CREATE_STDERR | F_FUNCTION)
	fork.Parameters.Params = plan.Parameters

	if len(plan.Stdin) > 0 {
		if plan.StdinDT == "" {
			plan.StdinDT = types.Generic
		}
		fork.Stdin.SetDataType(plan.StdinDT)
		_, err := fork.Stdin.Write([]byte(plan.Stdin))
		if err != nil {
			fmt.Println(err)
			return false
		}
	}

	// Run any initializing code...if defined
	if len(plan.PreBlock) > 0 {
		fork.Name = "(unit test preblock)"
		preExitNum, preForkErr = fork.Execute([]rune(plan.PreBlock))
	}

	testExitNum, testForkErr = runFunction(function, plan.Stdin != "", fork)

	// Run any clear down code...if defined
	fork.IsMethod = false
	if len(plan.PostBlock) > 0 {
		fork.Name = "(unit test postblock)"
		postExitNum, postForkErr = fork.Execute([]rune(plan.PostBlock))
	}

	b, err := fork.Stdout.ReadAll()
	if err != nil {
		fmt.Println(err)
		return false
	}
	stdout = string(b)

	b, err = fork.Stderr.ReadAll()
	if err != nil {
		fmt.Println(err)
		return false
	}
	stderr = string(b)

	// test fork errors

	if preForkErr != nil {
		results.Add(&TestResult{
			ColNumber:  fileRef.Column,
			LineNumber: fileRef.Line,
			Exec:       function,
			Params:     plan.Parameters,
			TestName:   testName,
			Status:     TestFailed,
			Message:    fmt.Sprintf("PreBlock failed to compile: %s", preForkErr),
		})
		return false
	}

	if postForkErr != nil {
		results.Add(&TestResult{
			ColNumber:  fileRef.Column,
			LineNumber: fileRef.Line,
			Exec:       function,
			Params:     plan.Parameters,
			TestName:   testName,
			Status:     TestFailed,
			Message:    fmt.Sprintf("PostBlock failed to compile: %s", postForkErr),
		})
		return false
	}

	if testForkErr != nil {
		results.Add(&TestResult{
			ColNumber:  fileRef.Column,
			LineNumber: fileRef.Line,
			Exec:       function,
			Params:     plan.Parameters,
			TestName:   testName,
			Status:     TestFailed,
			Message:    fmt.Sprintf("Block failed to compile: %s", testForkErr),
		})
		return false
	}

	// test exit numbers

	if preExitNum != 0 {
		//passed = false
		results.Add(&TestResult{
			ColNumber:  fileRef.Column,
			LineNumber: fileRef.Line,
			Exec:       function,
			Params:     plan.Parameters,
			TestName:   testName,
			Status:     TestInfo,
			Message:    fmt.Sprintf("PreBlock exit num non-zero: `%d`", preExitNum),
		})
	}

	if postExitNum != 0 {
		//passed = false
		results.Add(&TestResult{
			ColNumber:  fileRef.Column,
			LineNumber: fileRef.Line,
			Exec:       function,
			Params:     plan.Parameters,
			TestName:   testName,
			Status:     TestInfo,
			Message:    fmt.Sprintf("PostBlock exit num non-zero: `%d`", postExitNum),
		})
	}

	if testExitNum != plan.ExitNumber {
		passed = false
		results.Add(&TestResult{
			ColNumber:  fileRef.Column,
			LineNumber: fileRef.Line,
			Exec:       function,
			Params:     plan.Parameters,
			TestName:   testName,
			Status:     TestFailed,
			Message:    fmt.Sprintf("Exit num mismatch: exp `%d` act `%d`", plan.ExitNumber, testExitNum),
		})
	}

	// test stdout stream

	if stdout != plan.Stdout && plan.StdoutRx == "" {
		passed = false
		results.Add(&TestResult{
			ColNumber:  fileRef.Column,
			LineNumber: fileRef.Line,
			Exec:       function,
			Params:     plan.Parameters,
			TestName:   testName,
			Status:     TestFailed,
			Message:    fmt.Sprintf("Unexpected stdout: `%s`", stdout),
		})
	}

	if plan.StdoutRx != "" && plan.Stdout == "" {
		rx, err := regexp.Compile(plan.StdoutRx)
		if err != nil {
			passed = false
			results.Add(&TestResult{
				ColNumber:  fileRef.Column,
				LineNumber: fileRef.Line,
				Exec:       function,
				Params:     plan.Parameters,
				TestName:   testName,
				Status:     TestFailed,
				Message:    fmt.Sprintf("StdoutRx could not compile: %s", err),
			})
		} else if !rx.MatchString(stdout) {
			passed = false
			results.Add(&TestResult{
				ColNumber:  fileRef.Column,
				LineNumber: fileRef.Line,
				Exec:       function,
				Params:     plan.Parameters,
				TestName:   testName,
				Status:     TestFailed,
				Message:    fmt.Sprintf("StdoutRx did not match stdout: `%s`", stdout),
			})
		}
	}

	if plan.StdoutDT != "" && fork.Stdout.GetDataType() != plan.StdoutDT {
		passed = false
		results.Add(&TestResult{
			ColNumber:  fileRef.Column,
			LineNumber: fileRef.Line,
			Exec:       function,
			Params:     plan.Parameters,
			TestName:   testName,
			Status:     TestFailed,
			Message:    fmt.Sprintf("Stdout data-type mismatch: exp `%s` act `%s`", plan.StdoutDT, fork.Stdout.GetDataType()),
		})
	}

	// test stderr stream

	if stderr != plan.Stderr && plan.StderrRx == "" {
		passed = false
		results.Add(&TestResult{
			ColNumber:  fileRef.Column,
			LineNumber: fileRef.Line,
			Exec:       function,
			Params:     plan.Parameters,
			TestName:   testName,
			Status:     TestFailed,
			Message:    fmt.Sprintf("Unexpected stderr: `%s`", stderr),
		})
	}

	if plan.StderrRx != "" && plan.Stderr == "" {
		rx, err := regexp.Compile(plan.StderrRx)
		if err != nil {
			passed = false
			results.Add(&TestResult{
				ColNumber:  fileRef.Column,
				LineNumber: fileRef.Line,
				Exec:       function,
				Params:     plan.Parameters,
				TestName:   testName,
				Status:     TestFailed,
				Message:    fmt.Sprintf("StderrRx could not compile: %s", err),
			})
		} else if !rx.MatchString(stderr) {
			passed = false
			results.Add(&TestResult{
				ColNumber:  fileRef.Column,
				LineNumber: fileRef.Line,
				Exec:       function,
				Params:     plan.Parameters,
				TestName:   testName,
				Status:     TestFailed,
				Message:    fmt.Sprintf("StderrRx did not match stderr: `%s`", stderr),
			})
		}
	}

	if plan.StderrDT != "" && fork.Stderr.GetDataType() != plan.StderrDT {
		passed = false
		results.Add(&TestResult{
			ColNumber:  fileRef.Column,
			LineNumber: fileRef.Line,
			Exec:       function,
			Params:     plan.Parameters,
			TestName:   testName,
			Status:     TestFailed,
			Message:    fmt.Sprintf("Stderr data-type mismatch: exp `%s` act `%s`", plan.StderrDT, fork.Stderr.GetDataType()),
		})
	}

	// lastly, a passed message if no errors

	if passed {
		results.Add(&TestResult{
			ColNumber:  fileRef.Column,
			LineNumber: fileRef.Line,
			Exec:       function,
			Params:     plan.Parameters,
			TestName:   testName,
			Status:     TestPassed,
			//Message:    "",
		})
	}

	return passed
}

func runFunction(function string, isMethod bool, fork *Fork) (int, error) {
	fork.IsMethod = isMethod

	if strings.Contains(function, "/") {
		return runPrivate(function, fork)
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

func runPrivate(path string, fork *Fork) (int, error) {
	if path[0] == '/' {
		path = path[1:]
	}

	split := strings.Split(path, "/")
	if len(path) < 2 {
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
