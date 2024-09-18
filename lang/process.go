package lang

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/pipes"
	"github.com/lmorg/murex/lang/state"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansititle"
	"github.com/lmorg/murex/utils/crash"
)

var (
	// Interactive describes whether murex is running as an interactive shell or not
	Interactive bool

	// ShellProcess is the root murex process
	ShellProcess = &Process{}

	// MxFunctions is a table of global murex functions
	MxFunctions = NewMurexFuncs()

	// PrivateFunctions is a table of private murex functions
	PrivateFunctions = NewMurexPrivs()

	// GoFunctions is a table of available builtin functions
	GoFunctions = make(map[string]func(*Process) error)

	// MethodStdin is a table of all the different commands that can be used as methods
	MethodStdin = newMethods()

	// MethodStdout is a table of all the different output formats supported by a given command (by default)
	MethodStdout = newMethods()

	// GlobalVariables is a table of global variables
	GlobalVariables = NewGlobals()

	// ModuleVariables is a table of module specific variables
	ModuleVariables = NewModuleVars()

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

func DefineFunction(name string, fn func(*Process) error, StdoutDataType string) {
	GoFunctions[name] = fn
	MethodStdout.Define(name, StdoutDataType)
}

func DefineMethod(name string, fn func(*Process) error, StdinDataType, StdoutDataType string) {
	GoFunctions[name] = fn
	MethodStdin.Define(name, StdinDataType)
	MethodStdout.Define(name, StdoutDataType)
}

func indentError(err error) error {
	if err == nil {
		return err
	}

	s := strings.ReplaceAll(err.Error(), "\n", "\n    ")
	return errors.New(s)
}

func writeError(p *Process, err error) []byte {
	var msg string

	name := p.Name.String()
	if name == "exec" {
		exec, pErr := p.Parameters.String(0)
		if pErr == nil {
			name = exec
		}
	}

	var indentLen int
	if p.FileRef.Source.Module == app.ShellModule {
		msg = fmt.Sprintf("Error in `%s` (%d,%d): ", name, p.FileRef.Line, p.FileRef.Column)
		indentLen = len(msg) - 2
	} else {
		msg = fmt.Sprintf("Error in `%s` (%s %d,%d):\n      Command: %s\n      Error: ", name, p.FileRef.Source.Filename, p.FileRef.Line+1, p.FileRef.Column, string(p.raw))
		indentLen = 6 + 5
	}

	sErr := strings.ReplaceAll(err.Error(), utils.NewLineString, utils.NewLineString+strings.Repeat(" ", indentLen)+"> ")
	return []byte(msg + sErr)
}

func createProcess(p *Process, isMethod bool) {
	GlobalFIDs.Register(p) // This also registers the variables process
	p.CreationTime = time.Now()

	parseRedirection(p)

	name := p.Name.String()

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
		p.Stderr = p.Next.Stdin
	default:
		pipe, err := GlobalPipes.Get(p.NamedPipeErr)
		if err == nil {
			p.Stderr = pipe
		} else {
			p.Stderr.Writeln([]byte("invalid usage of named pipes: " + err.Error()))
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
		pipe, err := GlobalPipes.Get(p.NamedPipeOut)
		if err == nil {
			p.stdoutOldPtr = p.Stdout
			p.Stdout = pipe
		} else {
			p.Stderr.Writeln([]byte("invalid usage of named pipes: " + err.Error()))
		}
	}

	// Test cases
	if p.NamedPipeTest != "" {
		var stdout2, stderr2 *streams.Stdin
		p.Stdout, stdout2 = streams.NewTee(p.Stdout)
		p.Stderr, stderr2 = streams.NewTee(p.Stderr)
		err := p.Tests.SetStreams(p.NamedPipeTest, stdout2, stderr2, &p.ExitNum)
		if err != nil {
			p.Stderr.Writeln([]byte("invalid usage of named pipes: " + err.Error()))
		}
	}

	ttys(p)

	p.Stdout.Open()
	p.Stderr.Open()

	p.State.Set(state.Assigned)

	// Lets run `pipe` and `test` ahead of time to fudge the use of named pipes
	if name == "pipe" || (name == "test" && len(p.Parameters.PreParsed) > 0 &&
		(string(p.Parameters.PreParsed[0]) == "config" || // test config
			string(p.Parameters.PreParsed[0]) == "define" || // test define
			string(p.Parameters.PreParsed[0]) == "state")) { // test state

		_, params, err := ParseStatementParameters(p.raw, p)
		if err != nil {
			ShellProcess.Stderr.Writeln(writeError(p, err))
			if p.ExitNum == 0 {
				p.ExitNum = 1
			}

		} else {
			p.Parameters.DefineParsed(params)
			err = GoFunctions[name[:4]](p)
			if err != nil {
				ShellProcess.Stderr.Writeln(writeError(p, err))
				if p.ExitNum == 0 {
					p.ExitNum = 1
				}
			}
		}

		p.SetTerminatedState(true)
		p.State.Set(state.Executed)
	}
}

func executeProcess(p *Process) {
	defer crash.Handler()

	testStates(p)

	if p.HasTerminated() || p.HasCancelled() ||
		/*p.Parent.HasTerminated() ||*/ p.Parent.HasCancelled() {
		destroyProcess(p)
		return
	}

	p.State.Set(state.Starting)

	var err error
	name := p.Name.String()
	echo, err := p.Config.Get("proc", "echo", types.Boolean)
	if err != nil {
		echo = false
	}

	tmux, err := p.Config.Get("proc", "echo-tmux", types.Boolean)
	if err != nil {
		tmux = false
	}

	var parsedAlias bool

	n, params, err := ParseStatementParameters(p.raw, p)
	if err != nil {
		goto cleanUpProcess
	}
	if n != name {
		p.Name.Set(n)
		name = n
	}
	p.Parameters.DefineParsed(params)

	// Execute function
	p.StartTime = time.Now()

	if p.cache != nil && p.cache.use {
		// we have a preview cache, lets just write that and skip execution
		_, err = p.Stdout.Write(p.cache.b.stdout)
		if err != nil {
			panic(err)
		}
		p.Stdout.SetDataType(p.cache.dt.stdout)

		_, err = p.Stderr.Write(p.cache.b.stderr)
		if err != nil {
			panic(err)
		}
		p.Stderr.SetDataType(p.cache.dt.stderr)

		goto cleanUpProcess
	}

executeProcess:
	if !p.Background.Get() || debug.Enabled {
		if echo.(bool) {
			params := strings.Replace(strings.Join(p.Parameters.StringArray(), `", "`), "\n", "\n# ", -1)
			os.Stdout.WriteString("# " + name + `("` + params + `");` + utils.NewLineString)
		}

		if tmux.(bool) {
			ansititle.Tmux([]byte(name))
		}

		//ansititle.Write([]byte(name))
	}

	// execution mode:
	switch {
	case p.Scope.Id != ShellProcess.Id && PrivateFunctions.Exists(name, p.FileRef):
		// murex privates
		fn := PrivateFunctions.get(name, p.FileRef)
		if fn != nil {
			fork := p.Fork(F_FUNCTION)
			fork.Name.Set(name)
			fork.Parameters.CopyFrom(&p.Parameters)
			fork.FileRef = fn.FileRef
			p.State.Set(state.Executing)
			p.ExitNum, err = fork.Execute(fn.Block)
		}

	case GlobalAliases.Exists(name) && p.Parent.Name.String() != "alias" && !parsedAlias:
		// murex aliases
		alias := GlobalAliases.Get(name)
		p.Name.Set(alias[0])
		name = alias[0]
		p.Parameters.Prepend(alias[1:])
		parsedAlias = true
		goto executeProcess

	case MxFunctions.Exists(name):
		// murex functions
		fn := MxFunctions.get(name)
		if fn != nil {
			fork := p.Fork(F_FUNCTION)
			fork.Name.Set(name)
			fork.Parameters.CopyFrom(&p.Parameters)
			fork.FileRef = fn.FileRef
			err = fn.castParameters(fork.Process)
			if err == nil {
				p.State.Set(state.Executing)
				p.ExitNum, err = fork.Execute(fn.Block)
			}
		}

	case GoFunctions[name] != nil:
		// murex builtins
		p.State.Set(state.Executing)
		err = GoFunctions[name](p)

	default:
		if p.Parameters.Len() == 0 {
			v, _ := ShellProcess.Config.Get("shell", "auto-cd", types.Boolean)
			autoCd, _ := v.(bool)
			if autoCd {
				fileInfo, _ := os.Stat(name)
				if fileInfo != nil && fileInfo.IsDir() {
					p.Parameters.Prepend([]string{name})
					name = "cd"
					goto executeProcess
				}
			}
		}

		// shell execute
		p.Parameters.Prepend([]string{name})
		p.Name.Set("exec")
		// Don't put the following here: p.State.Set(state.Executing)
		// That should be done in the `exec` builtin instead
		err = GoFunctions["exec"](p)
		if err != nil && strings.Contains(err.Error(), "executable file not found") {
			_, cpErr := ParseExpression(p.raw, 0, false)
			err = fmt.Errorf("Not a valid expression:\n    %v\nNor a valid statement:\n    %v",
				indentError(cpErr), indentError(err))
		}
	}

	if p.cache != nil {
		p.cache.b.stdout, _ = p.cache.tee.stdout.ReadAll()
		p.cache.dt.stdout = p.cache.tee.stdout.GetDataType()
		p.cache.b.stderr, _ = p.cache.tee.stderr.ReadAll()
		p.cache.dt.stderr = p.cache.tee.stderr.GetDataType()
	}

cleanUpProcess:

	if err != nil {
		p.Stderr.Writeln(writeError(p, err))
		if p.ExitNum == 0 {
			p.ExitNum = 1
		}
	}

	if p.CCEvent != nil {
		p.CCEvent(name, p)
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
			p.Stderr.Writeln([]byte(fmt.Sprintf("!!! Error connecting to temporary named pipe '%s': %s", p.NamedPipeOut, err.Error())))
		} else {
			p.stdoutOldPtr.Open()
			_, err = out.WriteTo(p.stdoutOldPtr)
			p.stdoutOldPtr.Close()
			if err != nil && err != io.EOF {
				p.Stderr.Writeln([]byte(fmt.Sprintf("!!! Error piping from temporary named pipe '%s': %s", p.NamedPipeOut, err.Error())))
			}
		}

		err = GlobalPipes.Close(p.NamedPipeOut)
		if err != nil {
			p.Stderr.Writeln([]byte(fmt.Sprintf("!!! Error closing temporary named pipe '%s': %s", p.NamedPipeOut, err.Error())))
		}
	}

	for !p.Previous.HasTerminated() {
		// Code shouldn't really get stuck here.
		// This would only happen if someone abuses pipes on a function that has no stdin.
		time.Sleep(10 * time.Microsecond)
	}

	//debug.Json("Execute process (destroyProcess)", p)
	destroyProcess(p)
}

func waitProcess(p *Process) {
	//debug.Log("Waiting for", p.Name.String())
	<-p.WaitForTermination
	//debug.Log("Finished waiting for", p.Name.String())
}

func destroyProcess(p *Process) {
	//debug.Json("destroyProcess ()", p.Dump())
	// Clean up any context goroutines
	go p.Done()

	// Make special case for `bg` because that doesn't wait.
	if p.Name.String() != "bg" && !p.IsFork {
		//debug.Json("destroyProcess (p.WaitForTermination <- false)", p.Dump())
		p.WaitForTermination <- false
	}

	//debug.Json("destroyProcess (deregisterProcess)", p.Dump())
	deregisterProcess(p)
	//debug.Json("destroyProcess (end)", p.Dump())
}

// deregisterProcess deregisters a murex process, FID and mark variables for
// garbage collection.
func deregisterProcess(p *Process) {
	//debug.Json("deregisterProcess ()", p)

	p.State.Set(state.Terminating)

	p.Stdout.Close()
	p.Stderr.Close()

	p.SetTerminatedState(true)
	if !p.Background.Get() {
		ForegroundProc.Set(p.Next)
	}

	go func() {
		p.State.Set(state.AwaitingGC)
		GlobalFIDs.Deregister(p.Id)
	}()
}
