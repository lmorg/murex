package lang

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/lmorg/murex/utils"

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
			Status:   TestFailed,
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
	Parameters  []string
	Stdin       string
	StdoutMatch string
	StderrMatch string
	StdinType   string
	StdoutType  string
	StderrType  string
	StdoutRegex string // check this is the same as test define
	StderrRegex string // check this is the same as test define
	StdoutBlock string // check this is the same as test define
	StderrBlock string // check this is the same as test define
	ExitNum     int    // check this is the same as test define
	PreBlock    string
	PostBlock   string
}

func runTest(results *TestResults, fileRef *ref.File, plan *UnitTestPlan, function string) bool {
	var (
		preExitNum, testExitNum, postExitNum int
		preForkErr, testForkErr, postForkErr error
		stdoutType, stderrType               string
		bOut, bErr                           []byte

		fStdin         int
		passed         = true
		stdout, stderr string

		oblkStdout, eblkStdout   []byte // Std(out|err)Block; Stdout.ReadAll() []byte
		oblkStderr, eblkStderr   []byte // Std(out|err)Block; Stderr.ReadAll() []byte
		oblkExitNum, eblkExitNum int    // Std(out|err)Block; fork.Execute() int
		oblkErr, eblkErr         error  // Std(out|err)Block; fork.Execute() error
	)

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
			results.Add(&TestResult{
				ColNumber:  fileRef.Column,
				LineNumber: fileRef.Line,
				Exec:       function,
				Params:     plan.Parameters,
				TestName:   testName,
				Status:     TestFailed,
				Message:    fmt.Sprintf("PreBlock failed to compile: %s", preForkErr),
			})
		}

		if preExitNum != 0 {
			//passed = false
			results.Add(&TestResult{
				ColNumber:  fileRef.Column,
				LineNumber: fileRef.Line,
				Exec:       function,
				Params:     plan.Parameters,
				TestName:   testName,
				Status:     TestInfo,
				Message:    fmt.Sprintf("PreBlock exit num non-zero: %d", preExitNum),
			})
		}

		utReadAllOut(preFork.Stdout, results, plan, fileRef, "PreBlock", function, &passed)
		utReadAllErr(preFork.Stderr, results, plan, fileRef, "PreBlock", function, &passed)
	}

	// run function
	testExitNum, testForkErr = runFunction(function, plan.Stdin != "", fork)
	if testForkErr != nil {
		results.Add(&TestResult{
			ColNumber:  fileRef.Column,
			LineNumber: fileRef.Line,
			Exec:       function,
			Params:     plan.Parameters,
			TestName:   testName,
			Status:     TestFailed,
			Message:    fmt.Sprintf("testBlock failed to compile: %s", testForkErr),
		})
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
			results.Add(&TestResult{
				ColNumber:  fileRef.Column,
				LineNumber: fileRef.Line,
				Exec:       function,
				Params:     plan.Parameters,
				TestName:   testName,
				Status:     TestFailed,
				Message:    fmt.Sprintf("PostBlock failed to compile: %s", postForkErr),
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
				Message:    fmt.Sprintf("PostBlock exit num non-zero: %d", postExitNum),
			})
		}

		utReadAllOut(postFork.Stdout, results, plan, fileRef, "PostBlock", function, &passed)
		utReadAllErr(postFork.Stderr, results, plan, fileRef, "PostBlock", function, &passed)
	}

	// stdout

	bOut, err := fork.Stdout.ReadAll()
	if err != nil {
		fmt.Println(err)
		return false
	}
	stdout = string(bOut)
	stdoutType = fork.Stdout.GetDataType()

	// stderr

	bErr, err = fork.Stderr.ReadAll()
	if err != nil {
		fmt.Println(err)
		return false
	}
	stderr = string(bErr)
	stderrType = fork.Stderr.GetDataType()

	// stdout block

	if plan.StdoutBlock != "" {
		oFork := ShellProcess.Fork(fStdin | F_FUNCTION | F_CREATE_STDIN | F_CREATE_STDOUT | F_CREATE_STDERR)
		oFork.IsMethod = true
		oFork.Name = "(unit test StdoutBlock)"
		oFork.Stdin.SetDataType(stdoutType)
		_, err = oFork.Stdin.Write(bOut)
		if err != nil {
			fmt.Println(err)
			return false
		}
		oblkExitNum, oblkErr = oFork.Execute([]rune(plan.StdoutBlock))
		utReadAllOut(oFork.Stdout, results, plan, fileRef, "StdoutBlock", function, &passed)
		utReadAllErr(oFork.Stderr, results, plan, fileRef, "StdoutBlock", function, &passed)
	}

	// stderr block

	if plan.StderrBlock != "" {
		eFork := ShellProcess.Fork(F_FUNCTION | F_CREATE_STDIN | F_CREATE_STDOUT | F_CREATE_STDERR)
		eFork.IsMethod = true
		eFork.Name = "(unit test StderrBlock)"
		eFork.Stderr.SetDataType(stderrType)
		_, err = eFork.Stdin.Write(bErr)
		if err != nil {
			fmt.Println(err)
			return false
		}
		eblkExitNum, eblkErr = eFork.Execute([]rune(plan.StderrBlock))
		utReadAllOut(eFork.Stdout, results, plan, fileRef, "StderrBlock", function, &passed)
		utReadAllErr(eFork.Stderr, results, plan, fileRef, "StderrBlock", function, &passed)
	}

	// test exit number

	if testExitNum != plan.ExitNum {
		passed = false
		results.Add(&TestResult{
			ColNumber:  fileRef.Column,
			LineNumber: fileRef.Line,
			Exec:       function,
			Params:     plan.Parameters,
			TestName:   testName,
			Status:     TestFailed,
			Message:    fmt.Sprintf("ExitNum mismatch: exp %d act %d", plan.ExitNum, testExitNum),
		})
	}

	// test stdout stream

	if stdout != plan.StdoutMatch && plan.StdoutRegex == "" && plan.StdoutBlock == "" {
		passed = false
		results.Add(&TestResult{
			ColNumber:  fileRef.Column,
			LineNumber: fileRef.Line,
			Exec:       function,
			Params:     plan.Parameters,
			TestName:   testName,
			Status:     TestFailed,
			Message:    fmt.Sprintf("Unexpected stdout: '%s'", stdout),
		})
	}

	if plan.StdoutRegex != "" {
		rx, err := regexp.Compile(plan.StdoutRegex)
		if err != nil {
			passed = false
			results.Add(&TestResult{
				ColNumber:  fileRef.Column,
				LineNumber: fileRef.Line,
				Exec:       function,
				Params:     plan.Parameters,
				TestName:   testName,
				Status:     TestFailed,
				Message:    fmt.Sprintf("StdoutRegex could not compile: %s", err),
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
				Message:    fmt.Sprintf("StdoutRegex did not match stdout: '%s'", stdout),
			})
		}
	}

	if plan.StdoutBlock != "" {
		if oblkExitNum != 0 {
			passed = false
			results.Add(&TestResult{
				ColNumber:  fileRef.Column,
				LineNumber: fileRef.Line,
				Exec:       function,
				Params:     plan.Parameters,
				TestName:   testName,
				Status:     TestFailed,
				Message:    fmt.Sprintf("StdoutBlock exit num non-zero: %d", oblkExitNum),
			})
		}
		if oblkErr != nil {
			passed = false
			results.Add(&TestResult{
				ColNumber:  fileRef.Column,
				LineNumber: fileRef.Line,
				Exec:       function,
				Params:     plan.Parameters,
				TestName:   testName,
				Status:     TestFailed,
				Message:    fmt.Sprintf("StdoutBlock failed to compile: %s", oblkErr),
			})
		}
		if len(oblkStderr) != 0 {
			passed = false
			results.Add(&TestResult{
				ColNumber:  fileRef.Column,
				LineNumber: fileRef.Line,
				Exec:       function,
				Params:     plan.Parameters,
				TestName:   testName,
				Status:     TestFailed,
				Message:    fmt.Sprintf("StdoutBlock returned an error: %s", oblkStderr),
			})
		}
		if len(oblkStdout) != 0 {
			results.Add(&TestResult{
				ColNumber:  fileRef.Column,
				LineNumber: fileRef.Line,
				Exec:       function,
				Params:     plan.Parameters,
				TestName:   testName,
				Status:     TestInfo,
				Message:    fmt.Sprintf("StdoutBlock comment: %s", utils.CrLfTrim(oblkStdout)),
			})
		}
	}

	if plan.StdoutType != "" && stdoutType != plan.StdoutType {
		passed = false
		results.Add(&TestResult{
			ColNumber:  fileRef.Column,
			LineNumber: fileRef.Line,
			Exec:       function,
			Params:     plan.Parameters,
			TestName:   testName,
			Status:     TestFailed,
			Message:    fmt.Sprintf("Stdout data-type mismatch: exp '%s' act '%s'", plan.StdoutType, fork.Stdout.GetDataType()),
		})
	}

	// test stderr stream

	if stderr != plan.StderrMatch && plan.StderrRegex == "" && plan.StderrBlock == "" {
		passed = false
		results.Add(&TestResult{
			ColNumber:  fileRef.Column,
			LineNumber: fileRef.Line,
			Exec:       function,
			Params:     plan.Parameters,
			TestName:   testName,
			Status:     TestFailed,
			Message:    fmt.Sprintf("Unexpected stderr: '%s'", stderr),
		})
	}

	if plan.StderrRegex != "" {
		rx, err := regexp.Compile(plan.StderrRegex)
		if err != nil {
			passed = false
			results.Add(&TestResult{
				ColNumber:  fileRef.Column,
				LineNumber: fileRef.Line,
				Exec:       function,
				Params:     plan.Parameters,
				TestName:   testName,
				Status:     TestFailed,
				Message:    fmt.Sprintf("StderrRegex could not compile: %s", err),
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
				Message:    fmt.Sprintf("StderrRegex did not match stderr: '%s'", stderr),
			})
		}
	}

	if plan.StderrBlock != "" {
		if eblkExitNum != 0 {
			passed = false
			results.Add(&TestResult{
				ColNumber:  fileRef.Column,
				LineNumber: fileRef.Line,
				Exec:       function,
				Params:     plan.Parameters,
				TestName:   testName,
				Status:     TestFailed,
				Message:    fmt.Sprintf("StderrBlock exit num non-zero: %d", eblkExitNum),
			})
		}
		if eblkErr != nil {
			passed = false
			results.Add(&TestResult{
				ColNumber:  fileRef.Column,
				LineNumber: fileRef.Line,
				Exec:       function,
				Params:     plan.Parameters,
				TestName:   testName,
				Status:     TestFailed,
				Message:    fmt.Sprintf("StderrBlock failed to compile: %s", eblkErr),
			})
		}
		if len(eblkStderr) != 0 {
			passed = false
			results.Add(&TestResult{
				ColNumber:  fileRef.Column,
				LineNumber: fileRef.Line,
				Exec:       function,
				Params:     plan.Parameters,
				TestName:   testName,
				Status:     TestFailed,
				Message:    fmt.Sprintf("StderrBlock returned an error: %s", eblkStderr),
			})
		}
		if len(eblkStdout) != 0 {
			results.Add(&TestResult{
				ColNumber:  fileRef.Column,
				LineNumber: fileRef.Line,
				Exec:       function,
				Params:     plan.Parameters,
				TestName:   testName,
				Status:     TestInfo,
				Message:    fmt.Sprintf("StderrBlock comment: %s", utils.CrLfTrim(eblkStdout)),
			})
		}
	}

	if plan.StderrType != "" && stderrType != plan.StderrType {
		passed = false
		results.Add(&TestResult{
			ColNumber:  fileRef.Column,
			LineNumber: fileRef.Line,
			Exec:       function,
			Params:     plan.Parameters,
			TestName:   testName,
			Status:     TestFailed,
			Message:    fmt.Sprintf("Stderr data-type mismatch: exp '%s' act '%s'", plan.StderrType, fork.Stderr.GetDataType()),
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
			Message:    testPassedMessage,
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
