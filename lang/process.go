package lang

import (
	"fmt"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"os"
)

var ShellEnabled bool

func createProcess(p *proc.Process, f proc.Flow) {
	proc.ProcIDs.Add(p)

	if p.Parent.MethodRef == "" {
		p.Parent.MethodRef = "null"
	}

	if p.Name[0] == '!' {
		p.IsNot = true
	}

	local := "[" + p.Previous.Name + "]" + p.Name
	switch {
	case proc.GoFunctions[local].Func != nil && p.IsMethod &&
		(proc.GoFunctions[local].TypeIn == proc.GoFunctions[p.Previous.MethodRef].TypeOut ||
			proc.GoFunctions[local].TypeIn == types.Generic ||
			proc.GoFunctions[p.Previous.MethodRef].TypeOut == types.Generic):
		p.MethodRef = local

	case proc.GoFunctions[p.Name].Func != nil && p.IsMethod &&
		(proc.GoFunctions[p.Name].TypeIn == proc.GoFunctions[p.Previous.MethodRef].TypeOut ||
			proc.GoFunctions[p.Name].TypeIn == types.Generic ||
			proc.GoFunctions[p.Previous.MethodRef].TypeOut == types.Generic):
		p.MethodRef = p.Name

	case proc.GoFunctions[p.Name].Func != nil && !f.NewChain && !p.IsMethod &&
		(proc.GoFunctions[p.Name].TypeIn == types.Null ||
			proc.GoFunctions[p.Name].TypeIn == types.Generic):
		p.MethodRef = p.Name

	case proc.GoFunctions[p.Name].Func != nil && f.NewChain &&
		(proc.GoFunctions[p.Name].TypeIn == types.Null ||
			proc.GoFunctions[p.Name].TypeIn == types.Generic):
		p.MethodRef = p.Name

	case !p.IsMethod:
		p.Parameters.SetPrepend(p.Name)
		if f.NewChain && !f.PipeOut && !f.PipeErr && ShellEnabled {
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

func executeProcess(p *proc.Process) {
	debug.Json("Executing:", p)
	proc.GlobalVars.Dump()

	// Expand variables if parameter isn't a code block.
	/*for i := range p.Parameters {
		if len(p.Parameters[i]) > 1 && (p.Parameters[i][0] != '{' || p.Parameters[i][len(p.Parameters[i])-1] != '}') {
			GlobalVars.KeyValueReplace(&p.Parameters[i])
		}
	}*/
	parseParameters(&p.Parameters, &proc.GlobalVars)

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

	for !p.Previous.HasTerminated {
		// Code shouldn't really get stuck here.
		// This would only happen if someone abuses pipes on a function that has no stdin.
	}

	destroyProcess(p)
}

func waitProcess(p *proc.Process) {
	debug.Log("Waiting for", p.Name)
	for !p.HasTerminated {
		// Wait for process to terminate
	}
}

func destroyProcess(p *proc.Process) {
	debug.Json("Destroying:", p)
	p.Stdout.Close()
	p.Stderr.Close()
	p.HasTerminated = true
	debug.Log("Destroyed")
}
