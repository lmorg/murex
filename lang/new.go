package lang

import (
	"errors"
	"fmt"

	"github.com/lmorg/murex/builtins/pipes/null"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc/runmode"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
)

// ShellExitNum is for when running murex in interactive shell mode
var ShellExitNum int

// RunBlockShellConfigSpace is used for calling code blocks using the shell's
// namespace. eg commands initiated inside the interactive shell.
// This shouldn't be used under normal conditions; however if you do need to
// spawn a process alongside the shells config space then please first branch it:
//
//     branch := ShellProcess.BranchFID()
//     defer branch.Close()
//
// Then call `RunBlockExistingConfigSpace` with `branch.Process`.
func RunBlockShellConfigSpace(block []rune, stdin, stdout, stderr stdio.Io) (exitNum int, err error) {
	return processNewBlock(
		block, stdin, stdout, stderr,
		ShellProcess, ShellProcess.Config, nil,
		0,
	)
}

// RunBlockNewConfigSpace is for spawning new murex functions. eg `func {}`
// This shouldn't be used under normal conditions.
func RunBlockNewConfigSpace(block []rune, stdin, stdout, stderr stdio.Io, caller *Process) (exitNum int, err error) {
	return processNewBlock(
		block, stdin, stdout, stderr,
		caller, caller.Config.Copy(), nil,
		caller.PromptId,
	)
}

// RunBlockExistingConfigSpace is for code blocks as parameters (eg `if {}`,
// `try {}` etc) or inlining code blocks (eg `out @{g *}`)
// This should be the default way to call code blocks.
func RunBlockExistingConfigSpace(block []rune, stdin, stdout, stderr stdio.Io, caller *Process) (exitNum int, err error) {
	return processNewBlock(
		block, stdin, stdout, stderr,
		caller, caller.Config, caller.Tests,
		caller.PromptId,
	)
}

// RunBlockShellConfigSpaceWithPrompt is specifically for code being executed directly from the readline.
// This function should not be called by any other process aside shell.prompt()
func RunBlockShellConfigSpaceWithPrompt(block []rune, stdin, stdout, stderr stdio.Io, promptGoProc int) (exitNum int, err error) {
	return processNewBlock(
		block, stdin, stdout, stderr,
		ShellProcess, ShellProcess.Config, nil,
		promptGoProc,
	)
}

// processNewBlock parses new block and execute the code.
// Inputs are:
//     * the code block ([]rune),
//     * Stdin, stdout and stderr streams; or nil to black hole those data streams,
//     * caller Process to determine scope and any inherited properties,
//     * config namespace.
// Outputs are:
//     * exit number of the last process in the block,
//     * any errors raised during the parse.
func processNewBlock(block []rune, stdin, stdout, stderr stdio.Io, caller *Process, conf *config.Config, tests *Tests, promptGoProc int) (exitNum int, err error) {
	//debug.Log(string(block))

	if len(block) > 2 && block[0] == '{' && block[len(block)-1] == '}' {
		block = block[1 : len(block)-1]
	}

	container := new(Process)
	container.State = state.MemAllocated
	container.IsBackground = caller.IsBackground
	container.Name = caller.Name
	container.Scope = caller.Scope
	container.Parent = caller
	container.Id = caller.Id
	container.PromptId = promptGoProc
	container.LineNumber = caller.LineNumber
	container.ColNumber = caller.ColNumber
	container.Config = conf
	container.Variables = caller.Variables
	container.Module = caller.Module

	if tests != nil {
		container.Tests = tests
	} else {
		container.Tests = NewTests()
	}

	if caller.Id == ShellProcess.Id {
		container.ExitNum = ShellExitNum
	} else {
		container.RunMode = caller.RunMode
	}

	if stdin != nil {
		container.Stdin = stdin
	} else {
		container.Stdin = streams.NewStdin()
		container.Stdin.SetDataType(types.Null)
	}

	if stdout != nil {
		container.Stdout = stdout
	} else {
		if debug.Enabled {
			// This is TermErr despite being a Stdout stream because it is a debug
			// stream so we don't want to taint stdout with unexpected output.
			container.Stdout = term.NewErr(true)
		} else {
			container.Stdout = new(null.Null)
		}
	}
	container.Stdout.Open()
	defer container.Stdout.Close()

	if stderr != nil {
		container.Stderr = stderr
	} else {
		if debug.Enabled {
			container.Stderr = term.NewErr(true)
		} else {
			container.Stderr = new(null.Null)
		}
	}
	container.Stderr.Open()
	defer container.Stderr.Close()

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

	if tests == nil {
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
