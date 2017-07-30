package lang

import (
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	. "github.com/lmorg/murex/utils/consts"
	"os"
	"regexp"
	"strings"
)

var rxNamedPipeStdinOnly *regexp.Regexp = regexp.MustCompile(`^<[a-zA-Z0-9]+>$`)

func createProcess(p *proc.Process, isMethod bool) {
	// Create empty function so we don't have to check nil state when invoking kill, ie you try to kill a process
	// before it's fully started
	p.Kill = func() {}

	proc.GlobalFIDs.Register(p)
	parseRedirection(p)

	if rxNamedPipeStdinOnly.MatchString(p.Name) {
		p.Parameters.SetPrepend(p.Name[1 : len(p.Name)-1])
		p.Name = NamedPipeProcName
	}

	if p.Name[0] == '!' {
		p.IsNot = true
	}

	/*if (!proc.GlobalAliases.Exists(p.Name) || p.Parent.Name == "alias") &&
		p.Name[0] != '$' && proc.GoFunctions[p.Name].Func == nil {
		p.Parameters.SetPrepend(p.Name)
		// Make a special case of excluding `printf` from running inside a PTY as it hangs murex.
		// Obviously this shouldn't happen and in an ideal world I would fix murex instead of implementing this
		// horrible kludge. But I can live without `printf` being inside a PTY so I will class this bug as a low
		// priority.
		if f.NewChain && !f.PipeOut && !f.PipeErr && p.Name != "printf" {
			p.Name = CmdPty
		} else {
			p.Name = CmdExec
		}
	}*/

	//p.IsMethod = !f.NewChain
	p.IsMethod = isMethod

	p.State = state.Assigned

	return
}

func executeProcess(p *proc.Process) {
	var err error

	p.State = state.Starting

	echo, err := proc.GlobalConf.Get("shell", "echo", types.Boolean)
	if err != nil {
		echo = false
	}

	//debug.Json("Executing:", p)

	// Create a kill switch
	if p.Name != CmdExec && p.Name != CmdPty {
		p.Kill = func() { destroyProcess(p) }
	}

	if !p.IsBackground {
		proc.KillForeground = p.Kill
		proc.ForegroundProc = p
	}

	parseParameters(&p.Parameters, &proc.GlobalVars)

	switch p.NamedPipeOut {
	case "":
		//p.NamedPipeOut = "out"
	case "err":
		p.Stdout.SetDataType(types.Null)
		p.Stdout.Close()
		p.Stdout = p.Next.Stderr
	case "out":
		p.Stderr.Writeln([]byte("Invalid usage of named pipes: stdout defaults to <out>."))
	default:
		p.Stdout.SetDataType(types.Null)
		p.Stdout.Close()
		pipe, err := proc.GlobalPipes.Get(p.NamedPipeOut)
		if err == nil {
			p.Stdout = pipe
		} else {
			p.Stderr.Writeln([]byte("Invalid usage of named pipes: " + err.Error()))
		}
	}

	switch p.NamedPipeErr {
	case "":
		//p.NamedPipeErr = "err"
	case "err":
		p.Stderr.Writeln([]byte("Invalid usage of named pipes: stderr defaults to <err>."))
	case "out":
		p.Stderr.SetDataType(types.String)
		p.Stderr.Close()
		p.Stderr = p.Next.Stdout
	default:
		p.Stderr.SetDataType(types.String)
		p.Stderr.Close()
		pipe, err := proc.GlobalPipes.Get(p.NamedPipeErr)
		if err == nil {
			p.Stderr = pipe
		} else {
			p.Stderr.Writeln([]byte("Invalid usage of named pipes: " + err.Error()))
		}
	}

	p.Stderr.SetDataType(types.String)

	// Execute function.
	var parsedAlias bool
	p.State = state.Executing

executeProcess:
	if echo.(bool) {
		params := strings.Replace(strings.Join(p.Parameters.Params, `", "`), "\n", "\n# ", -1)
		os.Stdout.WriteString("# " + p.Name + `("` + params + `");` + utils.NewLineString)
	}

	switch {
	case proc.GlobalAliases.Exists(p.Name) && p.Parent.Name != "alias" && parsedAlias == false:
		//r := append(proc.GlobalAliases.Get(p.Name), []rune(" "+p.Parameters.StringAll())...)
		//p.Name = "alias"
		//p.ExitNum, err = ProcessNewBlock(r, p.Stdin, p.Stdout, p.Stderr, p)
		alias := proc.GlobalAliases.Get(p.Name)
		p.Name = alias[0]
		p.Parameters.Params = append(alias[1:], p.Parameters.Params...)
		parsedAlias = true
		goto executeProcess

	case proc.MxFunctions.Exists(p.Name):
		p.Scope = p
		r, err := proc.MxFunctions.Block(p.Name)
		if err == nil {
			p.ExitNum, err = ProcessNewBlock(r, p.Stdin, p.Stdout, p.Stderr, p)
		}

	case p.Name[0] == '$' && len(p.Name) > 1:
		s := proc.GlobalVars.GetString(p.Name[1:])
		p.Stdout.SetDataType(proc.GlobalVars.GetType(p.Name[1:]))
		_, err = p.Stdout.Write([]byte(s))

	case proc.GoFunctions[p.Name].Func != nil:
		err = proc.GoFunctions[p.Name].Func(p)

	default:
		//err = errors.New("Function not found (" + p.Name + ")! This is likely due to a bad alias.")
		p.Parameters.Params = append([]string{p.Name}, p.Parameters.Params...)

		// Make a special case of excluding `printf` from running inside a PTY as it hangs murex.
		// Obviously this shouldn't happen and in an ideal world I would fix murex instead of implementing this
		// horrible kludge. But I can live without `printf` being inside a PTY so I will class this bug as a low
		// priority.
		if !p.IsMethod && p.Stdout.IsTTY() && p.Name != "printf" {
			p.Name = CmdPty
		} else {
			p.Name = CmdExec
		}

		err = proc.GoFunctions[p.Name].Func(p)
	}
	p.State = state.Executed

	p.Stdout.DefaultDataType(err != nil)

	if err != nil {
		p.Stderr.Writeln([]byte("Error in `" + p.Name + "`: " + err.Error()))
		if p.ExitNum == 0 {
			p.ExitNum = 1
		}
	}

	for !p.Previous.HasTerminated() {
		// Code shouldn't really get stuck here.
		// This would only happen if someone abuses pipes on a function that has no stdin.
	}

	destroyProcess(p)
}

func waitProcess(p *proc.Process) {
	debug.Log("Waiting for", p.Name)
	<-p.WaitForTermination
}

func destroyProcess(p *proc.Process) {
	p.State = state.Terminating
	//debug.Json("Destroying:", p)
	p.Stdout.Close()
	p.Stderr.Close()

	p.SetTerminatedState(true)
	if p.Name != "fork" { // make special case for `fork` because that doesn't wait
		p.WaitForTermination <- false
	}
	debug.Log("Destroyed " + p.Name)

	proc.GlobalFIDs.Deregister(p.Id)
	p.State = state.AwaitingGC
}
