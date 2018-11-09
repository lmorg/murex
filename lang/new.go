package lang

import (
	"errors"
	"fmt"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/runmode"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
	"github.com/lmorg/murex/lang/types"
)

// ShellExitNum is for when running murex in interactive shell mode
var ShellExitNum int

// RunBlockShellConfigSpace is used for calling code blocks using the shell's
// namespace. eg commands initiated inside the interactive shell.
// This shouldn't be used under normal conditions; however if you do need to
// spawn a process alongside the shells config space then please first branch it:
//
//     branch := proc.ShellProcess.BranchFID()
//     defer branch.Close()
//
// Then call `RunBlockExistingConfigSpace` with `branch.Process`.
func RunBlockShellConfigSpace(block []rune, stdin, stdout, stderr stdio.Io) (exitNum int, err error) {
	return processNewBlock(
		block, stdin, stdout, stderr,
		proc.ShellProcess, proc.ShellProcess.Config, nil,
	)
}

// RunBlockNewConfigSpace is for spawning new murex functions. eg `func {}`
// This shouldn't be used under normal conditions.
func RunBlockNewConfigSpace(block []rune, stdin, stdout, stderr stdio.Io, caller *proc.Process) (exitNum int, err error) {
	return processNewBlock(
		block, stdin, stdout, stderr,
		caller, caller.Config.Copy(), nil,
	)
}

// RunBlockExistingConfigSpace is for code blocks as parameters (eg `if {}`,
// `try {}` etc) or inlining code blocks (eg `out @{g *}`)
// This should be the default way to call code blocks.
func RunBlockExistingConfigSpace(block []rune, stdin, stdout, stderr stdio.Io, caller *proc.Process) (exitNum int, err error) {
	return processNewBlock(
		block, stdin, stdout, stderr,
		caller, caller.Config, caller.Tests,
	)
}

// processNewBlock parses new block and execute the code.
// Inputs are:
//     * the code block ([]rune),
//     * Stdin, stdout and stderr streams; or nil to black hole those data streams,
//     * caller proc.Process to determine scope and any inherited properties,
//     * config namespace.
// Outputs are:
//     * exit number of the last process in the block,
//     * any errors raised during the parse.
func processNewBlock(block []rune, stdin, stdout, stderr stdio.Io, caller *proc.Process, conf *config.Config, tests *proc.Tests) (exitNum int, err error) {
	//debug.Log(string(block))

	if len(block) > 2 && block[0] == '{' && block[len(block)-1] == '}' {
		block = block[1 : len(block)-1]
	}

	container := new(proc.Process)
	container.State = state.MemAllocated
	container.IsBackground = caller.IsBackground
	container.Name = caller.Name
	container.Scope = caller.Scope
	container.Parent = caller
	container.Id = caller.Id
	container.LineNumber = caller.LineNumber
	container.ColNumber = caller.ColNumber
	container.Config = conf
	container.Variables = caller.Variables

	if tests != nil {
		container.Tests = tests
	} else {
		container.Tests = proc.NewTests()
	}

	if caller.Id == proc.ShellProcess.Id {
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
		if debug.Enable {
			// This is TermErr despite being a Stdout stream because it is a debug
			// stream so we don't want to taint stdout with unexpected output.
			container.Stdout = streams.NewTermErr(true)
		} else {
			container.Stdout = new(streams.Null)
		}
	}
	container.Stdout.Open()
	defer container.Stdout.Close()

	if stderr != nil {
		container.Stderr = stderr
	} else {
		if debug.Enable {
			container.Stderr = streams.NewTermErr(true)
		} else {
			container.Stderr = new(streams.Null)
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
			err = container.Tests.WriteResults(container.Config, proc.ShellProcess.Stderr)
			if err != nil {
				message := fmt.Sprintf("Error generating test results: %s.", err.Error())
				proc.ShellProcess.Stderr.Writeln([]byte(message))
			}
		}
	}

	//debug.Json("Finished running &tree", tree)
	return
}
