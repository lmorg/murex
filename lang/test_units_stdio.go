package lang

import (
	"fmt"

	"github.com/lmorg/murex/utils"

	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/ref"
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
			Message:    fmt.Sprintf("%s description failed on Stdout.ReadAll: %s", name, err),
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
			Message:    fmt.Sprintf("%s description failed on Stderr.ReadAll: %s", name, err),
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
