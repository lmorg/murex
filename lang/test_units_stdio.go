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
	"fmt"

	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/utils"
)

func utReadAllOut(std stdio.Io, results *TestResults, plan *UnitTestPlan, fileRef *ref.File, name string, function string, passed *bool) {
	b, err := std.ReadAll()
	if err != nil {
		results.Add(&TestResult{
			ColNumber:  fileRef.Column,
			LineNumber: fileRef.Line,
			Exec:       function,
			Params:     plan.Parameters,
			TestName:   testName,
			Status:     TestFailed,
			Message:    fmt.Sprintf("%s failed on Stdout.ReadAll: %s", name, err),
		})
		*passed = false
	}

	if len(b) != 0 {
		results.Add(&TestResult{
			ColNumber:  fileRef.Column,
			LineNumber: fileRef.Line,
			Exec:       function,
			Params:     plan.Parameters,
			TestName:   testName,
			Status:     TestInfo,
			Message:    fmt.Sprintf("%s output: %s", name, utils.CrLfTrim(b)),
		})
	}
}

func utReadAllErr(std stdio.Io, results *TestResults, plan *UnitTestPlan, fileRef *ref.File, name string, function string, passed *bool) {
	b, err := std.ReadAll()
	if err != nil {
		results.Add(&TestResult{
			ColNumber:  fileRef.Column,
			LineNumber: fileRef.Line,
			Exec:       function,
			Params:     plan.Parameters,
			TestName:   testName,
			Status:     TestFailed,
			Message:    fmt.Sprintf("%s failed on Stderr.ReadAll: %s", name, err),
		})
		*passed = false
	}

	if len(b) != 0 {
		results.Add(&TestResult{
			ColNumber:  fileRef.Column,
			LineNumber: fileRef.Line,
			Exec:       function,
			Params:     plan.Parameters,
			TestName:   testName,
			Status:     TestFailed,
			Message:    fmt.Sprintf("%s returned an err: %s", name, b),
		})
		*passed = false
	}
}
