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
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/utils"
)

func utBlock(plan *UnitTestPlan, fileRef *ref.File, block []rune, stdin []byte, dt string, property string, function string, results *TestResults, passed *bool) {
	fork := ShellProcess.Fork(F_FUNCTION | F_CREATE_STDIN | F_CREATE_STDERR | F_CREATE_STDOUT)
	fork.IsMethod = true
	fork.Name.Set("(unit test " + property + ")")
	fork.Stdin.SetDataType(dt)
	_, err := fork.Stdin.Write(stdin)
	if err != nil {
		utAddReport(results, fileRef, plan, function, TestError, tMsgWriteErr(property, err))
		*passed = false
		return
	}

	exitNum, err := fork.Execute(block)
	if err != nil {
		utAddReport(results, fileRef, plan, function, TestError, tMsgCompileErr(property, err))
		*passed = false
		return
	}

	if exitNum == 0 {
		utAddReport(results, fileRef, plan, function, TestInfo, tMsgExitNumZero(property))
	} else {
		utAddReport(results, fileRef, plan, function, TestFailed, tMsgExitNumNotZero(property, exitNum))
		*passed = false
	}

	utReadAllOut(fork.Stdout, results, plan, fileRef, property, function, passed)
	utReadAllErr(fork.Stderr, results, plan, fileRef, property, function, passed)
}

func utReadAllOut(std stdio.Io, results *TestResults, plan *UnitTestPlan, fileRef *ref.File, property string, function string, passed *bool) {
	b, err := std.ReadAll()
	if err != nil {
		utAddReport(results, fileRef, plan, function, TestError,
			tMsgReadErr("stdout", property, err))
		*passed = false
		return
	}

	if len(b) != 0 {
		utAddReport(results, fileRef, plan, function, TestInfo,
			tMsgStdout(property, utils.CrLfTrim(b)))
	}
}

func utReadAllErr(std stdio.Io, results *TestResults, plan *UnitTestPlan, fileRef *ref.File, property string, function string, passed *bool) {
	b, err := std.ReadAll()
	if err != nil {
		utAddReport(results, fileRef, plan, function, TestError,
			tMsgReadErr("stderr", property, err))
		*passed = false
		return
	}

	if len(b) != 0 {
		utAddReport(results, fileRef, plan, function, TestFailed,
			tMsgStdout(property, utils.CrLfTrim(b)))
		*passed = false
	}
}
