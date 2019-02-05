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
	"github.com/lmorg/murex/lang/types"
)

const (
	// F_SHELL will fork from the shell process rather than the calling process
	//F_SHELL = 1 << iota

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
	// You will almost never want to enable this as it breaks the expected
	// pattern of scoped variables
	F_PARENT_VARTABLE

	// F_CONFIG will fork the config table - eg when calling a new function
	F_CONFIG

	// F_NEW_TESTS will start a new scope for the testing framework
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

	/*// F_NEW_PROMPT_ID is used exclusively by the interactive shell to ensure
	// we have readline support even when proceses are shifted from the
	// background and forground (POSIX jobs)
	F_NEW_PROMPT_ID*/
)

type Fork struct {
	*Process
	fidRegistered bool
	newTestScope  bool
	/*Module      string
	Stdin       stdio.Io
	Stdout      stdio.Io
	Stderr      stdio.Io
	Variables   *Variables
	Config      *config.Config
	tests       *Tests
	parent      *Process
	Scope       *Process
	background  bool
	PromptId    int
	registerFid bool*/
}

// Fork will create a new handle for executing a code block
func (p *Process) Fork(flags int) *Fork {
	fork := new(Fork)
	fork.Process = new(Process)
	//p.hasTerminatedM.Lock()

	// This copies a sync.Mutex value, but on this occasion it is perfectly safe.
	//process := *p

	//p.hasTerminatedM.Unlock()
	//process.hasTerminatedM.Unlock()
	//fork.Process = &process

	fork.State = state.MemAllocated
	fork.PromptId = p.PromptId
	fork.LineNumber = p.LineNumber
	fork.ColNumber = p.ColNumber
	fork.IsBackground = flags&F_BACKGROUND != 0
	fork.PromptId = p.PromptId
	//fork.Id = p.Id

	fork.FidTree = make([]int, len(p.FidTree))
	copy(fork.FidTree, p.FidTree)

	if p.Id == ShellProcess.Id {
		fork.ExitNum = ShellExitNum
	} else {
		fork.RunMode = p.RunMode
	}

	if flags&F_NEW_MODULE == 0 {
		fork.Module = p.Module
	}

	if flags&F_FUNCTION != 0 {
		fork.Scope = p
		fork.Parent = p
		fork.newTestScope = true
		fork.Tests = NewTests()
		fork.Config = p.Config.Copy()
		fork.Variables = p.Variables
		fork.Id = p.Id
		//fork.Variables = ReferenceVariables(p.Variables)
		//fork.Name += " (fork)"
		//GlobalFIDs.Register(fork.Process)
		//fork.fidRegistered = true

	} else {
		fork.Scope = p.Scope

		if flags&F_PARENT_VARTABLE != 0 {
			fork.Variables = p.Variables
			fork.Id = p.Id
			fork.Parent = p

		} else {
			fork.Variables = ReferenceVariables(p.Variables)
			fork.Name += " (fork)"
			GlobalFIDs.Register(fork.Process)
			fork.fidRegistered = true
			fork.Parent = p
		}

		if flags&F_CONFIG != 0 {
			fork.Config = p.Config.Copy()
		} else {
			fork.Config = p.Config
		}

		if flags&F_NEW_TESTS != 0 {
			fork.newTestScope = true
			fork.Tests = NewTests()
		} else {
			fork.Tests = p.Tests
		}
	}

	//fork.Name += " (fork)"
	//GlobalFIDs.Register(fork.Process)
	//fork.fidRegistered = true

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

	fork.Stdout.Open()
	fork.Stderr.Open()

	return fork
}

/*func (p *Process) _Fork(flags int) *Fork {
	fork := new(Fork)

	if flags&F_SHELL != 0 {
		fork.parent = ShellProcess
	} else {
		fork.parent = p
	}

	if flags&F_NEW_MODULE == 0 {
		fork.Module = p.Module
	}

	if flags&F_FUNCTION != 0 {
		fork.Scope = p
		fork.Variables = ReferenceVariables(p.Variables)
		fork.Config = p.Config.Copy()
		fork.registerFid = true

	} else {
		fork.Scope = p.Scope

		if flags&F_PARENT_VARTABLE != 0 {
			fork.Variables = p.Variables
			//fork.Variables = ReferenceVariables(p.Variables)
			//fork.registerFid = true
		} else {
			fork.Variables = ReferenceVariables(p.Variables)
			fork.registerFid = true
			//fork.Variables = p.Variables

		}

		if flags&F_CONFIG != 0 {
			fork.Config = p.Config.Copy()
		} else {
			fork.Config = p.Config
		}

		if flags&F_NEW_TESTS == 0 {
			fork.tests = p.Tests
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

	fork.background = flags&F_BACKGROUND != 0
	fork.PromptId = p.PromptId

	return fork
}*/

func (fork *Fork) Execute(block []rune) (exitNum int, err error) {
	if fork.Module == "" {
		panic("missing module name")
	}

	if len(block) > 2 && block[0] == '{' && block[len(block)-1] == '}' {
		block = block[1 : len(block)-1]
	}

	if fork.fidRegistered {
		defer DeregisterProcess(fork.Process)
	} else {
		defer fork.Stdout.Close()
		defer fork.Stderr.Close()
	}

	tree, pErr := ParseBlock(block)
	if pErr.Code != 0 {
		fork.Stderr.Writeln([]byte(pErr.Message))
		//debug.Json("ParseBlock returned:", pErr)
		err = errors.New(pErr.Message)
		return 1, err
	}

	procs := compile(&tree, fork.Process)
	if len(procs) == 0 {
		return
	}
	ForegroundProc = &procs[0]

	// Support for different run modes:
	switch fork.RunMode {
	case runmode.Normal, runmode.Shell:
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
		testAutoReport, configErr := fork.Config.Get("test", "auto-report", types.Boolean)
		if configErr == nil && testAutoReport.(bool) {
			fork.Tests.ReportMissedTests(fork.Process)
			err = fork.Tests.WriteResults(fork.Config, ShellProcess.Stderr)
			if err != nil {
				message := fmt.Sprintf("Error generating test results: %s.", err.Error())
				ShellProcess.Stderr.Writeln([]byte(message))
			}
		}
	}

	return
}

/*// Execute will run a murex code block
func (fork *Fork) _Execute(block []rune) (exitNum int, err error) {
	if fork.Module == "" {
		panic("missing module name")
	}

	if len(block) > 2 && block[0] == '{' && block[len(block)-1] == '}' {
		block = block[1 : len(block)-1]
	}

	container := new(Process)
	container.State = state.MemAllocated
	container.IsBackground = fork.background
	container.Name = fork.parent.Name
	container.Scope = fork.Scope
	container.Module = fork.Module
	container.Parent = fork.parent.Parent
	container.Id = fork.parent.Id
	container.PromptId = fork.PromptId
	container.LineNumber = fork.parent.LineNumber
	container.ColNumber = fork.parent.ColNumber
	container.Config = fork.Config
	container.Variables = fork.Variables

	container.FidTree = make([]int, len(fork.parent.FidTree))
	copy(container.FidTree, fork.parent.FidTree)

	if container.Id == ShellProcess.Id {
		container.ExitNum = ShellExitNum
	} else {
		container.RunMode = fork.parent.RunMode
	}

	container.Stdin = fork.Stdin
	container.Stdout = fork.Stdout
	container.Stderr = fork.Stderr

	container.Stdout.Open()
	container.Stderr.Open()

	if fork.registerFid {
		container.Name += " (fork)"
		GlobalFIDs.Register(container)
		defer DeregisterProcess(container)
	} else {
		defer container.Stdout.Close()
		defer container.Stderr.Close()
	}

	if fork.tests != nil {
		container.Tests = fork.tests
	} else {
		container.Tests = NewTests()
	}

	tree, pErr := ParseBlock(block)
	if pErr.Code != 0 {
		container.Stderr.Writeln([]byte(pErr.Message))
		//debug.Json("ParseBlock returned:", pErr)
		err = errors.New(pErr.Message)
		return 1, err
	}

	procs := compile(&tree, container)
	if len(procs) == 0 {
		return
	}
	ForegroundProc = &procs[0]

	// Support for different run modes:
	switch container.RunMode {
	case runmode.Normal, runmode.Shell:
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

	if fork.tests == nil {
		testAutoReport, configErr := container.Config.Get("test", "auto-report", types.Boolean)
		if configErr == nil && testAutoReport.(bool) {
			container.Tests.ReportMissedTests(container)
			err = container.Tests.WriteResults(container.Config, ShellProcess.Stderr)
			if err != nil {
				message := fmt.Sprintf("Error generating test results: %s.", err.Error())
				ShellProcess.Stderr.Writeln([]byte(message))
			}
		}
	}

	return
}
*/
