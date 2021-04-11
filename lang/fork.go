package lang

import (
	"errors"
	"fmt"

	"github.com/lmorg/murex/builtins/pipes/null"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc/runmode"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/types"
)

const (
	// F_DEFAULTS is forking with within the existing function
	F_DEFAULTS = 0

	// F_NEW_MODULE will skip the stage of inheriting the module name from the
	// calling function. You will still then need to specify that module name
	// yourself. eg
	//
	//     fork := p.Fork(F_SHELL|F_NEW_MODULE)
	//     fork.Module = "package/module"
	//     exitNum, err := fork.Execute([]rune{})
	F_NEW_MODULE = 1 << iota

	// F_FUNCTION will assign a bunch of sane default properties for a function
	// call
	F_FUNCTION

	// F_PARENT_VARTABLE will bypass the automatic forking of the var table.
	// The plan is to make this the default because it's what you'd expect to
	// use inside builtins
	F_PARENT_VARTABLE

	// F_NEW_VARTABLE will fork the variable table (not needed when using
	// F_FUNCTION)
	// For reasons I haven't got to the bottom of yet, this is rather glitchy
	// inside builtins.
	F_NEW_VARTABLE

	// F_NEW_CONFIG will fork the config table - eg when calling a new function
	// (not needed when calling F_FUNCTION)
	F_NEW_CONFIG

	// F_NEW_TESTS will start a new scope for the testing framework (not needed
	// when calling F_FUNCTION)
	F_NEW_TESTS

	// F_BACKGROUND this process will run in the background
	F_BACKGROUND

	// F_CREATE_STDIN will create a new stdin stdio.Io interface
	F_CREATE_STDIN

	// F_CREATE_STDOUT will create a new stdout stdio.Io interface
	F_CREATE_STDOUT

	// F_CREATE_STDERR will create a new stderr stdio.Io interface
	F_CREATE_STDERR

	// F_NO_STDIN will ensure stdin will be a nil interface
	F_NO_STDIN

	// F_NO_STDOUT will ensure stdout will be a nil interface
	F_NO_STDOUT

	// F_NO_STDERR will ensure stderr will be a nil interface
	F_NO_STDERR
)

// Fork is a forked process
type Fork struct {
	*Process
	fidRegistered bool
	newTestScope  bool
}

// Fork will create a new handle for executing a code block
func (p *Process) Fork(flags int) *Fork {
	fork := new(Fork)
	fork.Process = new(Process)
	fork.Kill = func() {
		ShellProcess.Stderr.Writeln([]byte("!!! Murex currently doesn't support killing `(fork)` functions !!!"))
	}

	fork.State.Set(state.MemAllocated)
	fork.PromptId = p.PromptId
	fork.IsBackground = flags&F_BACKGROUND != 0 || p.IsBackground
	fork.PromptId = p.PromptId

	fork.Previous = p
	fork.Next = p.Next

	if p.Id == ShellProcess.Id {
		fork.ExitNum = ShellExitNum
	} else {
		fork.RunMode = p.RunMode
	}

	if flags&F_NEW_MODULE == 0 {
		fork.FileRef = p.FileRef
	} else {
		fork.FileRef = &ref.File{Source: new(ref.Source)}
	}

	if flags&F_FUNCTION != 0 {
		fork.Scope = fork.Process
		fork.Parent = fork.Process

		fork.Variables = NewVariables(fork.Process)
		GlobalFIDs.Register(fork.Process)
		fork.fidRegistered = true

		fork.Config = p.Config.Copy()

		fork.newTestScope = true
		fork.Tests = NewTests(fork.Process)

	} else {
		fork.Scope = p.Scope
		fork.Name = p.Name
		fork.Parameters = p.Parameters

		switch {
		case flags&F_PARENT_VARTABLE != 0:
			fork.Parent = p
			fork.Variables = p.Variables
			fork.Id = p.Id

		case flags&F_NEW_VARTABLE != 0:
			fork.Parent = p
			fork.Variables = p.Variables
			fork.Name += " (fork)"
			GlobalFIDs.Register(fork.Process)
			fork.fidRegistered = true

		default:
			//panic("must include either F_PARENT_VARTABLE or F_NEW_VARTABLE")
			fork.Parent = p
			fork.Variables = NewVariables(fork.Process)
			fork.Variables = p.Variables
			fork.Name += " (fork)"
			GlobalFIDs.Register(fork.Process)
			fork.fidRegistered = true
		}

		if flags&F_NEW_CONFIG != 0 {
			fork.Config = p.Config.Copy()
		} else {
			fork.Config = p.Config
		}

		if flags&F_NEW_TESTS != 0 {
			fork.newTestScope = true
			fork.Tests = NewTests(fork.Process)
		} else {
			fork.Tests = p.Tests
		}
	}

	switch {
	case flags&F_CREATE_STDIN != 0:
		fork.Stdin = streams.NewStdin()
	case flags&F_NO_STDIN != 0:
		fork.Stdin = streams.NewStdin()
		fork.Stdin.SetDataType(types.Null)
	default:
		fork.Stdin = p.Stdin
	}

	switch {
	case flags&F_CREATE_STDOUT != 0:
		fork.Stdout = streams.NewStdin()
	case flags&F_NO_STDOUT != 0:
		if debug.Enabled {
			// This is TermErr despite being a Stdout stream because it is a debug
			// stream so we don't want to taint stdout with unexpected output.
			fork.Stdout = term.NewErr(true)
		} else {
			fork.Stdout = new(null.Null)
		}
	default:
		fork.Stdout = p.Stdout
	}

	switch {
	case flags&F_CREATE_STDERR != 0:
		fork.Stderr = streams.NewStdin()
	case flags&F_NO_STDERR != 0:
		if debug.Enabled {
			// This is TermErr despite being a Stdout stream because it is a debug
			// stream so we don't want to taint stdout with unexpected output.
			fork.Stderr = term.NewErr(true)
		} else {
			fork.Stderr = new(null.Null)
		}
	default:
		fork.Stderr = p.Stderr
	}

	return fork
}

// ExecuteAsRunMode is a wrapper function for handling forks that need to
// comply with runmode changes (eg `try` and `trypipe` blocks). It returns err
// if the child process raises a runmode error and that should be returned in
// the calling builtin. Functions that shouldn't make use of this is processes
// that are spawned by the shell (eg dynamic autocomplete blocks or events).
func (fork *Fork) ExecuteAsRunMode(block []rune) error {
	fork.RunMode = fork.Parent.RunMode
	i, err := fork.Execute(block)
	if fork.RunMode != runmode.Try && fork.RunMode != runmode.TryPipe {
		return nil
	}

	if err != nil {
		return err
	}
	if i != 0 {
		return fmt.Errorf("Non-zero exit code: %d", i)
	}

	return nil
}

// Execute will run a murex code block
func (fork *Fork) Execute(block []rune) (exitNum int, err error) {
	switch {
	case fork.FileRef == nil:
		panic("fork.FileRef == nil in (fork *Fork).Execute()")
	case fork.FileRef.Source == nil:
		panic("fork.FileRef.Source == nil in (fork *Fork).Execute()")
	case fork.FileRef.Source.Module == "":
		panic("missing module name in (fork *Fork).Execute()")
	case fork.Name == "":
		panic("missing function name in (fork *Fork).Execute()")
	}

	fork.Stdout.Open()
	fork.Stderr.Open()

	if len(block) > 2 && block[0] == '{' && block[len(block)-1] == '}' {
		block = block[1 : len(block)-1]
	}

	if fork.fidRegistered {
		defer deregisterProcess(fork.Process)
	} else {
		defer fork.Stdout.Close()
		defer fork.Stderr.Close()
	}

	tree, pErr := ParseBlock(block)
	if pErr.Code != 0 {
		errMsg := fmt.Sprintf("Syntax error at %d,%d+%d: %s", fork.FileRef.Line, fork.FileRef.Column, pErr.EndByte, pErr.Message)
		fork.Stderr.Writeln([]byte(errMsg))
		err = errors.New(errMsg)
		return 1, err
	}

	procs, errNo := compile(&tree, fork.Process)
	if errNo != 0 {
		errMsg := fmt.Sprintf("Compilation Error at %d,%d+0: %s", fork.FileRef.Line, fork.FileRef.Column, errMessages[errNo])
		fork.Stderr.Writeln([]byte(errMsg))
		return errNo, errors.New(errMsg)
	}
	if len(procs) == 0 {
		if debug.Enabled {
			err = fmt.Errorf("Compilation Error at %d,%d+0: Empty code block", fork.FileRef.Line, fork.FileRef.Column)
		}
		return 0, err
	}

	if !fork.IsBackground {
		ForegroundProc.Set(&procs[0])
	}

	// Support for different run modes:
	switch fork.RunMode {
	case runmode.Normal:
		exitNum = runModeNormal(procs)

	case runmode.Try:
		exitNum = runModeTry(procs)

	case runmode.TryPipe:
		exitNum = runModeTryPipe(procs)

	//case runmode.Evil:
	//	panic("Not yet implemented")

	default:
		panic("Unknown run mode")
	}

	if fork.newTestScope {
		fork.Tests.ReportMissedTests(fork.Process)

		testAutoReport, configErr := fork.Config.Get("test", "auto-report", types.Boolean)
		if configErr == nil && testAutoReport.(bool) {
			err = fork.Tests.WriteResults(fork.Config, ShellProcess.Stderr)
			if err != nil {
				message := fmt.Sprintf("Error generating test results: %s.", err.Error())
				ShellProcess.Stderr.Writeln([]byte(message))
			}
		}
	}

	return
}
