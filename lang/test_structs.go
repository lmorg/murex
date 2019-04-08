package lang

/*
	This test library relates to the testing framework within the
	murex language itself rather than Go's test framework within
	the murex project.
*/

import (
	"fmt"
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

	// TestState is reporting the output from test state blocks
	TestState TestStatus = "STATE"

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
	mutex       sync.Mutex
	test        []*TestProperties
	Results     *TestResults
	stateBlocks map[string][]rune
}

// NewTests creates a new testing scope for Murex's test suite.NewTests.
// Please note this should NOT be confused with Go tests (go test)!
func NewTests(p *Process) (tests *Tests) {
	tests = new(Tests)
	tests.stateBlocks = make(map[string][]rune)

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
	return fmt.Errorf("Test already defined for '%s' in this scope", name)
}

// State creates a new test state
func (tests *Tests) State(name string, block []rune) error {
	tests.mutex.Lock()

	if len(tests.stateBlocks[name]) != 0 {
		tests.mutex.Unlock()
		return fmt.Errorf("Test state already defined for '%s' in this scope", name)
	}

	if len(block) == 0 {
		tests.mutex.Unlock()
		return fmt.Errorf("Test state for '%s' is an empty code block", name)
	}

	tests.stateBlocks[name] = block
	tests.mutex.Unlock()
	return nil
}

// Dump is used for `runtime --tests`
func (tests *Tests) Dump() interface{} {
	tests.mutex.Lock()

	names := make([]string, 0)
	for _, ptr := range tests.test {
		names = append(names, ptr.Name)
	}

	states := make(map[string]string)
	for name, state := range tests.stateBlocks {
		states[name] = string(state)
	}

	tests.mutex.Unlock()

	return map[string]interface{}{
		"test":  names,
		"state": states,
	}
}
