package lang

import (
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
	units []unitTests
}

// RunTests runs all unit tests against a specific murex function
func (ut UnitTests) RunTests(name string) {

}

// UnitTestPlans is defined via JSON and specifies an individual test plan
type UnitTestPlans struct {
	Parameters []string
	Stdin      string
	Stdout     string
	Stderr     string
	ExitCode   int
	PreBlock   string
	PostBlock  string
}

type unitTest struct {
	Function  string
	FileRef   *ref.File
	TestPlans []UnitTestPlans
}
