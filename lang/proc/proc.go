package proc

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/pipes"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"sync"
)

// Process - Each process running inside the murex shell will be one of these objects.
// It is equivalent to the /proc directory on Linux, albeit queried through murex as JSON.
// External processes will also appear in the host OS's process list.
type Process struct {
	Stdin              streams.Io
	Stdout             streams.Io
	Stderr             streams.Io
	Parameters         parameters.Parameters
	ExitNum            int
	Name               string
	Id                 int
	Path               string
	IsMethod           bool
	Scope              *Process  `json:"-"`
	Parent             *Process  `json:"-"`
	Previous           *Process  `json:"-"`
	Next               *Process  `json:"-"`
	WaitForTermination chan bool `json:"-"`
	Kill               func()    `json:"-"`
	IsNot              bool
	NamedPipeOut       string
	NamedPipeErr       string
	hasTerminatedM     sync.Mutex
	hasTerminatedV     bool
	State              state.FunctionStates
	IsBackground       bool
}

// HasTerminated checks if process has terminated.
// This is a function because terminated state can be subject to race conditions so we need a mutex to make the state
// thread safe.
func (p *Process) HasTerminated() (state bool) {
	p.hasTerminatedM.Lock()
	state = p.hasTerminatedV
	p.hasTerminatedM.Unlock()
	return
}

// SetTerminatedState sets the process terminated state.
// This is a function because terminated state can be subject to race conditions so we need a mutex to make the state
// thread safe.
func (p *Process) SetTerminatedState(state bool) {
	p.hasTerminatedM.Lock()
	p.hasTerminatedV = state
	p.hasTerminatedM.Unlock()
	return
}

var (
	ShellProcess   *Process                        = &Process{}
	MxFunctions    MurexFuncs                      = NewMurexFuncs()
	GoFunctions    map[string]func(*Process) error = make(map[string]func(*Process) error)
	GlobalVars     types.Vars                      = types.NewVariableGroup()
	GlobalConf     config.Config                   = config.NewConfiguration()
	GlobalAliases  Aliases                         = NewAliases()
	GlobalPipes    pipes.Named                     = pipes.NewNamed()
	GlobalFIDs     funcID                          = *newFuncID()
	KillForeground func()                          = func() {}
	ForegroundProc *Process                        = ShellProcess
)

// ExportRuntime exports a JSONable structure of the shell running state minus the FIDs
func ExportRuntime() map[string]interface{} {
	/*ListMap := func(m map[string]interface{}) (s []string) {
		for name := range m {
			s = append(s, name)
		}
		sort.Strings(s)
		return
	}*/

	m := make(map[string]interface{})
	m["Vars"] = GlobalVars.Dump()
	m["Aliases"] = GlobalAliases.Dump()
	m["Config"] = GlobalConf.Dump()
	m["Pipes"] = GlobalPipes.Dump()
	m["Funcs"] = MxFunctions.Dump()
	//m["Fids"] = GlobalFIDs.Dump()

	/*ctypes := make(map[string]interface{})
	ctypes["foreach"] = ListMap(streams.ReadArray)
	ctypes["formap"] = ListMap(streams.ReadMap)
	ctypes["index"] = ListMap(data.ReadIndexes)
	ctypes["format-in"] = ListMap(data.Marshal)
	ctypes["format-out"] = ListMap(data.Unmarshal)*/

	return m
}
