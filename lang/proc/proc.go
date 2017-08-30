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
	LineNumber         int
	ColNumber          int
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
	// ShellProcess is the root murex process
	ShellProcess *Process = &Process{}

	// MxFunctions is a table of global murex functions
	MxFunctions MurexFuncs = NewMurexFuncs()

	// GoFunctions is a table of available builtin functions
	GoFunctions map[string]func(*Process) error = make(map[string]func(*Process) error)

	// GlobalVars is a table of global variables
	GlobalVars types.Vars = types.NewVariableGroup()

	// GlobalConf is a table of global config options
	GlobalConf config.Config = config.NewConfiguration()

	// GlobalAliases is a table of global aliases
	GlobalAliases Aliases = NewAliases()

	// GlobalPipes is a table of  named pipes
	GlobalPipes pipes.Named = pipes.NewNamed()

	// GlobalFIDs is a table of running murex processes
	GlobalFIDs funcID = *newFuncID()

	// KillForeground is the `kill` function for whichever FID currently has "focus"
	KillForeground func() = func() {}

	// ForegroundProc is the murex FID which currently has "focus"  Em3w
	ForegroundProc *Process = ShellProcess
)

// ExportRuntime exports a JSONable structure of the shell running state
func ExportRuntime() map[string]interface{} {
	m := make(map[string]interface{})
	m["Vars"] = GlobalVars.Dump()
	m["Aliases"] = GlobalAliases.Dump()
	m["Config"] = GlobalConf.Dump()
	m["Pipes"] = GlobalPipes.Dump()
	m["Funcs"] = MxFunctions.Dump()
	m["Fids"] = GlobalFIDs.Dump()

	return m
}
