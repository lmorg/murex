package lang

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc/pipes"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/consts"
)

var (
	// ShellProcess is the root murex process
	ShellProcess = &Process{}

	// MxFunctions is a table of global murex functions
	MxFunctions = NewMurexFuncs()

	// PrivateFunctions is a table of private murex functions
	PrivateFunctions = NewMurexPrivs()

	// GoFunctions is a table of available builtin functions
	GoFunctions = make(map[string]func(*Process) error)

	// GlobalVariables is a table of global variables
	GlobalVariables = NewGlobals()

	// GlobalAliases is a table of global aliases
	GlobalAliases = NewAliases()

	// GlobalPipes is a table of  named pipes
	GlobalPipes = pipes.NewNamed()

	// GlobalFIDs is a table of running murex processes
	GlobalFIDs = *newFuncID()

	// GlobalUnitTests is a class for all things murex unit tests
	GlobalUnitTests = new(UnitTests)

	// ForegroundProc is the murex FID which currently has "focus"
	ForegroundProc = newForegroundProc()

	// ShellExitNum is for when running murex in interactive shell mode
	ShellExitNum int
)

var (
	rxNamedPipeStdinOnly = regexp.MustCompile(`^<[a-zA-Z0-9]+>$`)
	rxVariables          = regexp.MustCompile(`^\$([_a-zA-Z0-9]+)(\[(.*?)\]|)$`)
)

func writeError(p *Process, err error) []byte {
	if p.FileRef.Source.Module == app.Name {
		return []byte(fmt.Sprintf("Error in `%s` (%d,%d): %s", p.Name.String(), p.FileRef.Line, p.FileRef.Column, err.Error()))
	}
	return []byte(fmt.Sprintf("Error in `%s` (%s %d,%d): %s", p.Name.String(), p.FileRef.Source.Filename, p.FileRef.Line+1, p.FileRef.Column, err.Error()))
}

func createProcess(p *Process, isMethod bool) {
	GlobalFIDs.Register(p) // This also registers the variables process
	p.CreationTime = time.Now()

	parseRedirection(p)

	name := p.Name.String()

	if rxNamedPipeStdinOnly.MatchString(name) {
		p.Parameters.SetPrepend(name[1 : len(name)-1])
		p.Name.Set(consts.NamedPipeProcName)
		name = consts.NamedPipeProcName
	}

	if name[0] == '!' {
		p.IsNot = true
	}

	p.IsMethod = isMethod

	// We do stderr first so we can log errors in the stdout pipe to stderr
	switch p.NamedPipeErr {
	case "":
		p.NamedPipeErr = "err"
	case "err":
		//p.Stderr.Writeln([]byte("Invalid usage of named pipes: stderr defaults to <err>."))
	case "out":
		//p.Stderr.SetDataType(types.Generic)
		p.Stderr = p.Next.Stdin //p.Next.Stdout
	default:
		//p.Stderr.SetDataType(types.Generic)
		pipe, err := GlobalPipes.Get(p.NamedPipeErr)
		if err == nil {
			p.Stderr = pipe
		} else {
			p.Stderr.Writeln([]byte("Invalid usage of named pipes: " + err.Error()))
		}
	}

	// We do stdout last so we can log errors in the stdout pipe to stderr
	switch p.NamedPipeOut {
	case "":
		p.NamedPipeOut = "out"
	case "err":
		p.Stdout.SetDataType(types.Generic)
		p.Stdout = p.Next.Stderr
	case "out":
		//p.Stderr.Writeln([]byte("Invalid usage of named pipes: stdout defaults to <out>."))
	default:
		//p.Stdout.SetDataType(types.Null)
		pipe, err := GlobalPipes.Get(p.NamedPipeOut)
		if err == nil {
			p.stdoutOldPtr = p.Stdout
			p.Stdout = pipe
		} else {
			p.Stderr.Writeln([]byte("Invalid usage of named pipes: " + err.Error()))
		}
	}

	// Test cases
	if p.NamedPipeTest != "" {
		var stdout2, stderr2 *streams.Stdin
		p.Stdout, stdout2 = streams.NewTee(p.Stdout)
		p.Stderr, stderr2 = streams.NewTee(p.Stderr)
		err := p.Tests.SetStreams(p.NamedPipeTest, stdout2, stderr2, &p.ExitNum)
		if err != nil {
			p.Stderr.Writeln([]byte("Invalid usage of named pipes: " + err.Error()))
		}
	}

	p.Stdout.Open()
	p.Stderr.Open()

	//p.Stderr.SetDataType(types.Generic)

	p.State.Set(state.Assigned)

	// Lets run `pipe` and `test` ahead of time to fudge the use of named pipes
	if name == "pipe" || name == "test" {
		ParseParameters(p, &p.Parameters)
		err := GoFunctions[name](p)
		if err != nil {
			ShellProcess.Stderr.Writeln(writeError(p, err))
			if p.ExitNum == 0 {
				p.ExitNum = 1
			}
		}
		p.SetTerminatedState(true)
		p.State.Set(state.Executed)
	}

	return
}

func executeProcess(p *Process) {
	testStates(p)

	name := p.Name.String()

	if p.HasTerminated() {
		destroyProcess(p)
		return
	}

	p.State.Set(state.Starting)

	echo, err := p.Config.Get("proc", "echo", types.Boolean)
	if err != nil {
		echo = false
		err = nil
	}

	p.Context, p.Done = context.WithCancel(context.Background())

	p.Kill = func() {
		p.Stdin.ForceClose()
		p.Stdout.ForceClose()
		p.Stderr.ForceClose()
		p.Done()
	}

	ParseParameters(p, &p.Parameters)

	// Execute function.
	var parsedAlias bool
	p.State.Set(state.Executing)
	p.StartTime = time.Now()

	if err := GlobalFIDs.Executing(p.Id); err != nil {
		panic(err)
	}
executeProcess:

	if echo.(bool) {
		params := strings.Replace(strings.Join(p.Parameters.Params, `", "`), "\n", "\n# ", -1)
		os.Stdout.WriteString("# " + name + `("` + params + `");` + utils.NewLineString)
	}

	// execution mode:
	switch {
	case GlobalAliases.Exists(name) && p.Parent.Name.String() != "alias" && !parsedAlias:
		// murex aliases
		alias := GlobalAliases.Get(name)
		p.Name.Set(alias[0])
		name = alias[0]
		p.Parameters.Params = append(alias[1:], p.Parameters.Params...)
		parsedAlias = true
		goto executeProcess

	case MxFunctions.Exists(name):
		// murex functions
		fn := MxFunctions.get(name)
		if fn != nil {
			fork := p.Fork(F_FUNCTION)
			fork.Name.Set(name)
			fork.Parameters = p.Parameters
			fork.FileRef = fn.FileRef
			p.ExitNum, err = fork.Execute(fn.Block)
		}

	case p.Scope.Id != ShellProcess.Id && PrivateFunctions.Exists(name, p.FileRef.Source.Module):
		// murex privates
		fn := PrivateFunctions.get(name, p.FileRef.Source.Module)
		if fn != nil {
			fork := p.Fork(F_FUNCTION)
			fork.Name.Set(name)
			fork.Parameters = p.Parameters
			fork.FileRef = fn.FileRef
			p.ExitNum, err = fork.Execute(fn.Block)
		}

	case name[0] == '$':
		// variables as functions
		match := rxVariables.FindAllStringSubmatch(name+p.Parameters.StringAll(), -1)
		switch {
		case len(name) == 1:
			err = errors.New("Variable token, `$`, used without specifying variable name")
		case len(match) == 0 || len(match[0]) == 0:
			err = errors.New("`" + name[1:] + "` is not a valid variable name")
		case match[0][2] == "":
			s := p.Variables.GetString(match[0][1])
			p.Stdout.SetDataType(p.Variables.GetDataType(match[0][1]))
			_, err = p.Stdout.Write([]byte(s))
		default:
			block := []rune("$" + match[0][1] + "->[" + match[0][3] + "]")
			p.Fork(F_PARENT_VARTABLE).Execute(block)
		}

	case name == "@g":
		// auto globbing
		err = autoGlob(p)
		if err == nil {
			goto executeProcess
		}

	case GoFunctions[name] != nil:
		// murex builtins
		err = GoFunctions[name](p)

	default:
		// shell execute
		p.Parameters.Params = append([]string{name}, p.Parameters.Params...)
		p.Name.Set("exec")
		name = "exec"
		err = GoFunctions["exec"](p)
	}

	//p.Stdout.DefaultDataType(err != nil)

	if err != nil {
		p.Stderr.Writeln(writeError(p, err))
		if p.ExitNum == 0 {
			p.ExitNum = 1
		}
	}

	p.State.Set(state.Executed)

	if p.NamedPipeTest != "" {
		testEnabled, err := p.Config.Get("test", "enabled", types.Boolean)
		if err == nil && testEnabled.(bool) {
			p.Tests.Compare(p.NamedPipeTest, p)
		}
	}

	if len(p.NamedPipeOut) > 7 /* tmp:$FID/$MD5 */ && p.NamedPipeOut[:4] == "tmp:" {
		out, err := GlobalPipes.Get(p.NamedPipeOut)
		if err != nil {
			p.Stderr.Writeln([]byte(fmt.Sprintf("Error connecting to temporary named pipe '%s': %s", p.NamedPipeOut, err.Error())))
		} else {
			p.stdoutOldPtr.Open()
			_, err = out.WriteTo(p.stdoutOldPtr)
			p.stdoutOldPtr.Close()
			if err != nil && err != io.EOF {
				p.Stderr.Writeln([]byte(fmt.Sprintf("Error piping from temporary named pipe '%s': %s", p.NamedPipeOut, err.Error())))
			}
		}

		err = GlobalPipes.Close(p.NamedPipeOut)
		if err != nil {
			p.Stderr.Writeln([]byte(fmt.Sprintf("Error closing temporary named pipe '%s': %s", p.NamedPipeOut, err.Error())))
		}
	}

	for !p.Previous.HasTerminated() {
		// Code shouldn't really get stuck here.
		// This would only happen if someone abuses pipes on a function that has no stdin.
	}

	destroyProcess(p)
}

func waitProcess(p *Process) {
	//debug.Log("Waiting for", p.Name)
	<-p.WaitForTermination
}

func destroyProcess(p *Process) {
	// Clean up any context goroutines
	go p.Done()

	// Make special case for `bg` because that doesn't wait.
	if p.Name.String() != "bg" {
		p.WaitForTermination <- false
	}

	deregisterProcess(p)
}

// deregisterProcess deregisters a murex process, FID and mark variables for
// garbage collection.
func deregisterProcess(p *Process) {
	p.State.Set(state.Terminating)

	p.Stdout.Close()
	p.Stderr.Close()

	p.SetTerminatedState(true)
	if !p.Background.Get() {
		if p.Next == nil {
			debug.Json("p", p)
		}
		ForegroundProc.Set(p.Next)
	}

	go func() {
		p.State.Set(state.AwaitingGC)
		//CloseScopedVariables(p)
		GlobalFIDs.Deregister(p.Id)
	}()
}
