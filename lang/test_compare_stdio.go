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
)

func testBlock(test *TestProperties, p *Process, block []rune, stdin []byte, dt string, property string, failed *bool) {
	fork := p.Fork(F_FUNCTION | F_CREATE_STDIN | F_CREATE_STDERR | F_CREATE_STDOUT)
	fork.IsMethod = true
	fork.Name.Set("(pipe test " + property + ")")
	fork.Stdin.SetDataType(dt)
	_, err := fork.Stdin.Write(stdin)
	if err != nil {
		p.Tests.AddResult(test, p, TestError, tMsgWriteErr(property, err))
		*failed = true
		return
	}

	exitNum, err := fork.Execute(block)
	if err != nil {
		p.Tests.AddResult(test, p, TestError, tMsgCompileErr(property, err))
		return
	}

	if exitNum == 0 {
		p.Tests.AddResult(test, p, TestInfo, tMsgExitNumZero(property))
	} else {
		p.Tests.AddResult(test, p, TestFailed, tMsgExitNumNotZero(property, exitNum))
		*failed = true
	}

	testReadAllOut(test, p, fork.Stdout, property, failed)
	testReadAllErr(test, p, fork.Stderr, property, failed)
}

func testReadAllOut(test *TestProperties, p *Process, std stdio.Io, property string, failed *bool) {
	b, err := std.ReadAll()
	if err != nil {
		p.Tests.AddResult(test, p, TestError, tMsgReadErr("stdout", property, err))
		*failed = true
	}

	if len(b) != 0 {
		p.Tests.AddResult(test, p, TestInfo, tMsgStdout(property, b))
		*failed = true
	}
}

func testReadAllErr(test *TestProperties, p *Process, std stdio.Io, property string, failed *bool) {
	b, err := std.ReadAll()
	if err != nil {
		p.Tests.AddResult(test, p, TestError, tMsgReadErr("stderr", property, err))
		*failed = true
	}

	if len(b) != 0 {
		p.Tests.AddResult(test, p, TestFailed, tMsgStderr(property, b))
		*failed = true
	}
}
