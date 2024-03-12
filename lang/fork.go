package lang

import (
	"context"
	"errors"
	"fmt"

	"github.com/lmorg/murex/builtins/pipes/null"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/runmode"
	"github.com/lmorg/murex/lang/state"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/crash"
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

	// F_PREVIEW
	F_PREVIEW
)

var (
	ShowPrompt = make(chan bool, 1)
	HidePrompt = make(chan bool, 1)

	ModuleRunModes map[string]runmode.RunMode = make(map[string]runmode.RunMode)
)

// Fork is a forked process
type Fork struct {
	*Process
	fidRegistered bool
	newTestScope  bool
	preview       bool
}

const ForkSuffix = " (fork)"

// Fork will create a new handle for executing a code block
func (p *Process) Fork(flags int) *Fork {
	fork := new(Fork)
	fork.Process = new(Process)
	fork.SetTerminatedState(true)
	fork.Forks = p.Forks
	trace(fork.Process)
	fork.raw = p.raw

	fork.State.Set(state.MemAllocated)
	fork.Background.Set(flags&F_BACKGROUND != 0 || p.Background.Get())

	fork.IsMethod = p.IsMethod
	fork.OperatorLogicAnd = p.OperatorLogicAnd
	fork.OperatorLogicOr = p.OperatorLogicOr
	fork.IsNot = p.IsNot

	fork.Previous = p.Previous
	fork.Next = p.Next

	fork.preview = flags&F_PREVIEW != 0

	if p.Id == ShellProcess.Id {
		fork.ExitNum = ShellExitNum
	}

	if flags&F_NEW_MODULE == 0 {
		fork.FileRef = p.FileRef
	}

	if flags&F_FUNCTION != 0 {
		fork.Scope = fork.Process
		fork.Parent = fork.Process
		fork.Context, fork.Done = context.WithCancel(context.Background())
		fork.Kill = fork.Done

		fork.Variables = NewVariables(fork.Process)
		GlobalFIDs.Register(fork.Process)
		fork.fidRegistered = true

		fork.Config = p.Config.Copy()

		fork.newTestScope = true
		fork.Tests = NewTests(fork.Process)

	} else {
		fork.Scope = p.Scope
		fork.Name.Set(p.Name.String())
		//fork.Context, fork.Done = p.Context, p.Done
		fork.Context = p.Context
		fork.Done = func() { deregisterProcess(fork.Process) }

		if p.Scope.RunMode > runmode.Default {
			fork.RunMode = p.Scope.RunMode
		}
		if p.RunMode > runmode.Default {
			fork.RunMode = p.RunMode
		}

		switch {
		case flags&F_PARENT_VARTABLE != 0:
			fork.Parent = p.Parent
			fork.Variables = p.Variables
			fork.Id = p.Id

		case flags&F_NEW_VARTABLE != 0:
			fork.Parent = p.Parent
			fork.Variables = p.Variables
			fork.Name.Append(ForkSuffix)
			GlobalFIDs.Register(fork.Process)
			fork.fidRegistered = true
			fork.IsFork = true

		default:
			//panic("must include either F_PARENT_VARTABLE or F_NEW_VARTABLE")
			fork.Parent = p.Parent
			fork.Variables = NewVariables(fork.Process)
			fork.Variables = p.Variables
			fork.Name.Append(ForkSuffix)
			GlobalFIDs.Register(fork.Process)
			fork.fidRegistered = true
			fork.IsFork = true
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

// Execute will run a murex code block
func (fork *Fork) Execute(block []rune) (exitNum int, err error) {
	defer crash.Handler()

	switch {
	case fork.FileRef == nil:
		panic("fork.FileRef == nil in (fork *Fork).Execute()")
	case fork.FileRef.Source == nil:
		panic("fork.FileRef.Source == nil in (fork *Fork).Execute()")
	case fork.FileRef.Source.Module == "":
		panic("missing module name in (fork *Fork).Execute()")
	case fork.Name.String() == "":
		panic("missing function name in (fork *Fork).Execute()")
	}

	moduleRunMode := ModuleRunModes[fork.FileRef.Source.Module]
	if moduleRunMode > 0 && fork.RunMode == 0 {
		fork.RunMode = moduleRunMode
	}

	fork.Stdout.Open()
	fork.Stderr.Open()

	if len(block) > 2 && block[0] == '{' && block[len(block)-1] == '}' {
		block = block[1 : len(block)-1]
	}

	if fork.fidRegistered {
		defer deregisterProcess(fork.Process)
	} else {
		defer fork.SetTerminatedState(true)
		defer fork.Stdout.Close()
		defer fork.Stderr.Close()
	}

	tree, err := ParseBlock(block)
	if err != nil {
		return 1, err
	}

	procs, errNo := compile(tree, fork.Process)
	if errNo != 0 {
		errMsg := fmt.Sprintf("compilation Error at %d,%d+0 (%s): %s",
			fork.FileRef.Line, fork.FileRef.Column, fork.FileRef.Source.Module, errMessages[errNo])
		fork.Stderr.Writeln([]byte(errMsg))
		return errNo, errors.New(errMsg)
	}
	if len(*procs) == 0 {
		return 0, nil
	}

	id := fork.Process.Forks.add(procs)
	defer fork.Process.Forks.delete(id)

	if fork.preview {
		err := previewCache.compile(tree, procs)
		if err != nil {
			return 0, err
		}
	}

	if !fork.Background.Get() {
		ForegroundProc.Set(&(*procs)[0])
	}

	// Support for different run modes:
	switch fork.RunMode {
	case runmode.Default, runmode.Normal:
		exitNum = runModeNormal(procs)

	case runmode.BlockUnsafe, runmode.FunctionUnsafe, runmode.ModuleUnsafe:
		_ = runModeNormal(procs)
		exitNum = 0

	case runmode.BlockTry, runmode.FunctionTry, runmode.ModuleTry:
		exitNum = runModeTry(procs, false)

	case runmode.BlockTryPipe, runmode.FunctionTryPipe, runmode.ModuleTryPipe:
		exitNum = runModeTryPipe(procs, false)

	case runmode.BlockTryErr, runmode.FunctionTryErr, runmode.ModuleTryErr:
		exitNum = runModeTry(procs, true)

	case runmode.BlockTryPipeErr, runmode.FunctionTryPipeErr, runmode.ModuleTryPipeErr:
		exitNum = runModeTryPipe(procs, true)

	default:
		panic("unknown run mode")
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
