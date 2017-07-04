package lang

import (
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"os"
	"strings"
)

func createProcess(p *proc.Process, f proc.Flow) {
	//if p.Parent.MethodRef == "" {
	//	p.Parent.MethodRef = "null"
	//}

	if p.Name[0] == '!' {
		p.IsNot = true
	}

	/*local := "[" + p.Previous.Name + "]" + p.Name
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
		// Forcing `printf` to `exec` is a bit of a kludge.
		if f.NewChain && !f.PipeOut && !f.PipeErr && p.Name != "printf"  {
			p.MethodRef = "pty"
		} else {
			p.MethodRef = "exec"
		}

	default:
		p.MethodRef = "die"
		os.Stderr.WriteString(fmt.Sprintf("Methodable function `%s` does not exist for `%s.(%s)`\n",
			p.Name, p.Previous.Name, proc.GoFunctions[p.Previous.Name].TypeOut))
	}*/

	//p.ReturnType = proc.GoFunctions[p.MethodRef].TypeOut

	if !proc.GlobalAliases.Exists(p.Name) && proc.GoFunctions[p.Name].Func == nil {
		p.Parameters.SetPrepend(p.Name)
		if f.NewChain && !f.PipeOut && !f.PipeErr && p.Name != "printf" {
			p.Name = "pty"
		} else {
			p.Name = "exec"
		}
	}

	p.IsMethod = !f.NewChain

	return
}

func executeProcess(p *proc.Process) {
	debug.Json("Executing:", p)

	parseParameters(&p.Parameters, &proc.GlobalVars)

	// A little catch for unexpected behavior.
	// This shouldn't ever happen so lets produce a stack trace for debugging.
	if proc.GoFunctions[p.Name].Func == nil {
		panic("Failed to execute GoFunc[mapRef] `" + p.Name + "`. This should never happen!!")
	}

	// Echo
	echo, err := proc.GlobalConf.Get("shell", "Echo", types.Boolean)
	if err != nil {
		panic(err.Error())
	}
	if echo.(bool) {
		params := strings.Replace(strings.Join(p.Parameters.Params, `", "`), "\n", "\n# ", -1)
		os.Stdout.WriteString("# " + p.Name + `("` + params + `");` + utils.NewLineString)
	}

	// Execute function.
	p.Stderr.SetDataType(types.String)
	err = proc.GoFunctions[p.Name].Func(p)
	p.Stdout.DefaultDataType(err != nil)
	if err != nil {
		p.Stderr.Writeln([]byte("Error in `" + p.Name + "`: " + err.Error()))
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
	p.HasTerminated = <-p.WaitForTermination
}

func destroyProcess(p *proc.Process) {
	debug.Json("Destroying:", p)
	p.Stdout.Close()
	p.Stderr.Close()
	p.WaitForTermination <- true
	debug.Log("Destroyed " + p.Name)
}
