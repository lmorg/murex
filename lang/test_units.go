package lang

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/lmorg/murex/lang/ref"
)

// The naming convention here is basically the inverse of Go's test naming
// convention. ie Go source files will be named "test_unit.go" (because calling
// it unit_test.go would mean it's a Go test rather than murex test) and the
// code is named UnitTestPlans (etc) rather than TestUnitPlans (etc) because
// the latter might suggest they would be used by `go test`. This naming
// convention is a little counterintuitive but it at least avoids naming
// conflicts with `go test`.

// UnitTests is an exportable class for all things murex unit tests
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

// Run all unit tests against a specific murex function
func (ut *UnitTests) Run(function string) bool {
	ut.mutex.Lock()
	utCopy := make([]*unitTest, len(ut.units))
	copy(utCopy, ut.units)
	ut.mutex.Unlock()

	passed := true

	for i := range utCopy {
		if utCopy[i].Function == function {
			passed = passed && runTest(utCopy[i].TestPlan, function)
		}
	}

	return passed
}

// UnitTestPlan is defined via JSON and specifies an individual test plan
type UnitTestPlan struct {
	Parameters []string
	Stdin      string
	Stdout     string
	Stderr     string
	ExitNumber int // check this is the same as test define
	PreBlock   string
	PostBlock  string
}

func runTest(test *UnitTestPlan, function string) bool {
	var (
		preExitNum, testExitNum, postExitNum int
		preForkErr, testForkErr, postForkErr error
		F_STDIN                              int
		passed                               bool
		stdout, stderr                       string
	)

	if len(test.Stdin) == 0 {
		F_STDIN = F_NO_STDIN
	} else {
		F_STDIN = F_CREATE_STDIN
	}

	fork := ShellProcess.Fork(F_STDIN | F_CREATE_STDOUT | F_CREATE_STDERR | F_FUNCTION)
	fork.Name = function
	if len(test.Stdin) > 0 {
		_, err := fork.Stdin.Write([]byte(test.Stdin))
		//fork.Stdin.Close()
		if err != nil {
			fmt.Println(err)
			return false
		}
	}

	// Run any initializing code...if defined
	if len(test.PreBlock) > 0 {
		preExitNum, preForkErr = fork.Execute([]rune(test.PreBlock))
	}

	testExitNum, testForkErr = runFunction(function, fork)

	// Run any clear down code...if defined
	if len(test.PostBlock) > 0 {
		postExitNum, postForkErr = fork.Execute([]rune(test.PostBlock))
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

	fmt.Println("unit test:",
		preExitNum, testExitNum, postExitNum,
		preForkErr, testForkErr, postForkErr,
		F_STDIN)

	passed = testExitNum != test.ExitNumber || preExitNum != 0 || postExitNum != 0 ||
		testForkErr != nil || preForkErr != nil || postForkErr != nil ||
		stdout != test.Stdout || stderr != test.Stderr

	return passed
}

func runFunction(function string, fork *Fork) (int, error) {
	if strings.Contains(function, "/") {
		return 0, errors.New("TODO: support me!")
	} else {
		block, err := MxFunctions.Block(function)
		if err != nil {
			return 0, err
		}

		return fork.Execute(block)
	}
}
