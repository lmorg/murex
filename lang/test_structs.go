package lang

/*
	This test library relates to the testing framework within the
	murex language itself rather than Go's test framework within
	the murex project.
*/

import (
	"errors"
	"regexp"
	"sync"

	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
)

// TestProperties are the values prescribed to an individual test case
type TestProperties struct {
	Name       string
	out        *TestChecks
	err        *TestChecks
	exitNumPtr *int
	exitNum    int
	HasRan     bool
}

// TestChecks are the pipe streams and what test case to check against
type TestChecks struct {
	stdio    stdio.Io
	Regexp   *regexp.Regexp
	Block    []rune
	RunBlock func(*Process, []rune, []byte) ([]byte, []byte, error)
}

// TestResult is a record for each test result
type TestResult struct {
	Status     TestStatus
	TestName   string
	Message    string
	Exec       string
	Params     []string
	LineNumber int
	ColNumber  int
}

// TestResults is a class for the entire result set
type TestResults struct {
	mutex   sync.Mutex
	results []*TestResult
}

// Add appends a result to TestResults
func (tr *TestResults) Add(result *TestResult) {
	tr.mutex.Lock()
	tr.results = append(tr.results, result)
	tr.mutex.Unlock()
}

// Len returns the length of the results slice
func (tr *TestResults) Len() int {
	tr.mutex.Lock()
	i := len(tr.results)
	tr.mutex.Unlock()
	return i
}

// Dump returns the slice for runtime diagnositics
func (tr *TestResults) Dump() interface{} {
	return tr.results
}

// TestStatus is a summarised stamp for a particular result
type TestStatus string

const (
	// TestPassed means the test has passed
	TestPassed TestStatus = "PASSED"

	// TestFailed means the test has failed
	TestFailed TestStatus = "FAILED"

	// TestError means there was an error running that test case
	TestError TestStatus = "ERROR"

	// TestInfo is for any additional information on a test that might help
	// debug. This is only provided when `verbose` is enabled: `test verbose`
	TestInfo TestStatus = "INFORM"

	// TestMissed means that test was not run (this is usually because
	// it was inside a parent control block - eg if / switch / etc -
	// which flowed down a different pathway. eg:
	//
	//     if { true } else { out <test_example> "example" }
	//
	// `test_example` would not run because `if` would not run the
	// `else` block.
	TestMissed TestStatus = "MISSED"
)

// Tests is a class of all the tests that needs to run inside a
// particular scope, plus all of it's results.
type Tests struct {
	mutex   sync.Mutex
	test    []*TestProperties
	Results *TestResults
}

// NewTests creates a new testing scope for Murex's test suite.NewTests.
// Please note this should NOT be confused with Go tests (go test)!
func NewTests(p *Process) (tests *Tests) {
	tests = new(Tests)

	if p.Id == ShellProcess.Id {
		tests.Results = new(TestResults)
		return
	}

	autoReport, err := p.Parent.Config.Get("test", "auto-report", types.Boolean)
	if err != nil {
		autoReport = true
	}

	if autoReport.(bool) {
		tests.Results = new(TestResults)
	} else {
		tests.Results = ShellProcess.Tests.Results
	}

	return
}

// Define is the method used to define a new test case
func (tests *Tests) Define(name string, out *TestChecks, err *TestChecks, exitNum int) error {
	tests.mutex.Lock()

	var i int
	for ; i < len(tests.test); i++ {
		if tests.test[i].Name == name {
			goto define
		}
	}

	tests.test = append(tests.test, &TestProperties{
		Name:    name,
		out:     out,
		err:     err,
		exitNum: exitNum,
	})

	tests.mutex.Unlock()
	return nil

define:
	tests.mutex.Unlock()
	return errors.New("Test already defined for '" + name + "' in this scope.")
}

// Dump is used for `runtime --tests`
func (tests *Tests) Dump() []string {
	tests.mutex.Lock()

	names := make([]string, 0)

	for _, ptr := range tests.test {
		names = append(names, ptr.Name)
	}

	tests.mutex.Unlock()
	return names
}
