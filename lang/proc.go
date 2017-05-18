package lang

import (
	"fmt"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/proc"
	"os"
)

type Flow struct {
	PipeOut  bool
	PipeErr  bool
	NewChain bool
	Last     bool
}

var (
	ProcIDs Pid
)

/*
func (p *Process) OverrideProcName(name string) {
	p.Name = name
	p.MethodRef = name
}
*/
func CreateProcess(p *proc.Process, f Flow) {
	ProcIDs.Add(p)

	if p.Parent.MethodRef == "" {
		p.Parent.MethodRef = "null"
	}

	if p.Name[0] == '!' {
		p.Not = true
	}

	local := "[" + p.Previous.Name + "]" + p.Name
	switch {
	case proc.GoFunctions[local].Func != nil && p.Method &&
		(proc.GoFunctions[local].TypeIn == proc.GoFunctions[p.Previous.MethodRef].TypeOut ||
			proc.GoFunctions[local].TypeIn == types.Generic ||
			proc.GoFunctions[p.Previous.MethodRef].TypeOut == types.Generic):
		p.MethodRef = local

	case proc.GoFunctions[p.Name].Func != nil && p.Method &&
		(proc.GoFunctions[p.Name].TypeIn == proc.GoFunctions[p.Previous.MethodRef].TypeOut ||
			proc.GoFunctions[p.Name].TypeIn == types.Generic ||
			proc.GoFunctions[p.Previous.MethodRef].TypeOut == types.Generic):
		p.MethodRef = p.Name

	case proc.GoFunctions[p.Name].Func != nil && !f.NewChain && !p.Method &&
		(proc.GoFunctions[p.Name].TypeIn == types.Null ||
			proc.GoFunctions[p.Name].TypeIn == types.Generic):
		p.MethodRef = p.Name

	case proc.GoFunctions[p.Name].Func != nil && f.NewChain &&
		(proc.GoFunctions[p.Name].TypeIn == types.Null ||
			proc.GoFunctions[p.Name].TypeIn == types.Generic):
		p.MethodRef = p.Name

	case !p.Method:
		p.Parameters.SetPrepend(p.Name)
		if f.NewChain && !f.PipeOut && !f.PipeErr {
			p.MethodRef = "pty"
		} else {
			p.MethodRef = "exec"
		}

	default:
		p.MethodRef = "die"
		os.Stderr.WriteString(fmt.Sprintf("Methodable function `%s` does not exist for `%s.(%s)`\n",
			p.Name, p.Previous.Name, proc.GoFunctions[p.Previous.Name].TypeOut))
	}

	p.ReturnType = proc.GoFunctions[p.MethodRef].TypeOut
	return
}

func DestroyProcess(p *proc.Process) {
	debug.Json("Destroying:", p)
	p.Stdout.Close()
	p.Stderr.Close()
	p.Terminated = true
	debug.Log("Destroyed")
}

func ExecuteProcess(p *proc.Process) {
	debug.Json("Executing:", p)
	//proc.GlobalVars.Dump()

	// Expand variables if parameter isn't a code block.
	/*for i := range p.Parameters {
		if len(p.Parameters[i]) > 1 && (p.Parameters[i][0] != '{' || p.Parameters[i][len(p.Parameters[i])-1] != '}') {
			GlobalVars.KeyValueReplace(&p.Parameters[i])
		}
	}*/
	ParseParameters(&p.Parameters, &proc.GlobalVars)

	// A little catch for unexpected behavior.
	// This shouldn't ever happen so lets produce a stack trace for debugging.
	if proc.GoFunctions[p.MethodRef].Func == nil {
		panic("Failed to execute GoFunc[mapRef] `" + p.MethodRef + "`. This should never happen!!")
	}

	// Execute function.
	err := proc.GoFunctions[p.MethodRef].Func(p)
	if err != nil {
		p.Stderr.Writeln([]byte("Error in `" + p.MethodRef + "`: " + err.Error()))
		if p.ExitNum == 0 {
			p.ExitNum = 1
		}
	}

	for !p.Previous.Terminated {
		// Code shouldn't really get stuck here.
		// This would only happen if someone abuses pipes on a function that has no stdin.
	}

	DestroyProcess(p)
}

func (p *Process) Wait() {
	debug.Log("Waiting for", p.Name)
	for !p.Terminated {
		// Wait for process to terminate
	}
}
