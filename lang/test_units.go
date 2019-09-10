package lang

import (
	"strings"

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
	units []unitTest
}

// RunTests runs all unit tests against a specific murex function
func (ut UnitTests) RunTests(name string) {
	for i := range ut.units {
		if ut.units[i].Function == name {
			runTest(&ut.units[i].TestPlan, name)
		}
	}
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

type unitTest struct {
	Function string // if private it should contain path module path
	FileRef  *ref.File
	TestPlan UnitTestPlan
}

func runTest(test *UnitTestPlan, name string) {
	var (
		preExitNum, testExitNum, postExitNum int
		preForkErr, testForkErr, postForkErr error
		F_NO_STDIN                           int
	)

	if len(test.Stdin) == 0 {
		F_STDIN = F_NO_STDIN
	} else {
		F_STDIN = F_CREATE_STDIN
	}

	fork := ShellProcess.Fork(F_STDIN | F_CREATE_STDOUT | F_CREATE_STDERR | F_FUNCTION)
	if len(test.Stdin) > 0 {
		fork.Stdin.Write([]byte(test.Stdin))
	}

	// Run any initializing code...if defined
	if len(test.PreBlock) > 0 {
		preExitNum, preForkErr = fork.Execute([]rune(test.PreBlock))
	}

	testExitNum, testForkErr := runFunction(name, fork)

	// Run any clear down code...if defined
	if len(test.PostBlock) > 0 {
		postExitNum, postForkErr = fork.Execute([]rune(test.PostBlock))
	}
}

func runFunction(name string, fork *Fork) (int, error) {
	if strings.Contains(name, "/") {
		// run private
	} else {
		block, err := MxFunctions.Block(name)
		if err != nil {
			return 0, err
		}

		return fork.Execute(block)
	}
}
