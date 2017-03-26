package proc

import (
	"fmt"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"os"
)

type Flow struct {
	PipeOut  bool
	PipeErr  bool
	NewChain bool
	Last     bool
}

type Process struct {
	Stdin      streams.Io
	Stdout     streams.Io
	Stderr     streams.Io
	Parameters Parameters
	ExitNum    int
	Name       string
	Id         int
	Path       string
	Method     bool
	Parent     *Process `json:"-"`
	Previous   *Process `json:"-"`
	Next       *Process `json:"-"`
	Terminated bool
	Not        bool
	mapRef     string
	ReturnType string
}

type GoFunction struct {
	Func    func(*Process) error
	TypeIn  string
	TypeOut string
}

var (
	GlobalVars  types.Vars            = types.NewVariableGroup()
	GoFunctions map[string]GoFunction = make(map[string]GoFunction)
	GlobalConf  config.Config         = config.NewConfiguration()
	ProcIDs     Pid
)

func (p *Process) OverrideProcName(name string) {
	p.Name = name
	p.mapRef = name
}

func CreateProcess(p *Process, f Flow) {
	ProcIDs.Add(p)

	if p.Parent.mapRef == "" {
		p.Parent.mapRef = "null"
	}

	if p.Name[0] == '!' {
		p.Not = true
	}

	local := "[" + p.Previous.Name + "]" + p.Name
	switch {
	case GoFunctions[local].Func != nil && p.Method &&
		(GoFunctions[local].TypeIn == GoFunctions[p.Previous.mapRef].TypeOut ||
			GoFunctions[local].TypeIn == types.Generic ||
			GoFunctions[p.Previous.mapRef].TypeOut == types.Generic):
		p.mapRef = local

	case GoFunctions[p.Name].Func != nil && p.Method &&
		(GoFunctions[p.Name].TypeIn == GoFunctions[p.Previous.mapRef].TypeOut ||
			GoFunctions[p.Name].TypeIn == types.Generic ||
			GoFunctions[p.Previous.mapRef].TypeOut == types.Generic):
		p.mapRef = p.Name

	case GoFunctions[p.Name].Func != nil && !f.NewChain && !p.Method &&
		(GoFunctions[p.Name].TypeIn == types.Null ||
			GoFunctions[p.Name].TypeIn == types.Generic):
		p.mapRef = p.Name

	case GoFunctions[p.Name].Func != nil && f.NewChain &&
		(GoFunctions[p.Name].TypeIn == types.Null ||
			GoFunctions[p.Name].TypeIn == types.Generic):
		p.mapRef = p.Name

	case !p.Method:
		p.Parameters = append(Parameters{p.Name}, p.Parameters...)
		if f.NewChain && !f.PipeOut && !f.PipeErr {
			p.mapRef = "pty"
		} else {
			p.mapRef = "exec"
		}

	default:
		p.mapRef = "die"
		os.Stderr.WriteString(fmt.Sprintf("Methodable function `%s` does not exist for `%s.(%s)`\n",
			p.Name, p.Previous.Name, GoFunctions[p.Previous.Name].TypeOut))
	}

	p.ReturnType = GoFunctions[p.mapRef].TypeOut
	return
}

func DestroyProcess(p *Process) {
	debug.Json("Destroying:", p)
	p.Stdout.Close()
	p.Stderr.Close()
	p.Terminated = true
	debug.Log("Destroyed")
}

func (p *Process) Execute() {
	debug.Json("Executing:", p)
	GlobalVars.Dump()

	// Expand variables if parameter isn't a code block.
	for i := range p.Parameters {
		if len(p.Parameters[i]) > 1 && (p.Parameters[i][0] != '{' || p.Parameters[i][len(p.Parameters[i])-1] != '}') {
			GlobalVars.KeyValueReplace(&p.Parameters[i])
		}
	}

	// A little catch for unexpected behavior.
	// This shouldn't ever happen so lets produce a stack trace for debugging.
	if GoFunctions[p.mapRef].Func == nil {
		panic("Failed to execute GoFunc[mapRef] `" + p.mapRef + "`. This should never happen!!")
	}

	// Execute function.
	err := GoFunctions[p.mapRef].Func(p)
	if err != nil {
		p.Stderr.Writeln([]byte(err.Error()))
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
	}
}
