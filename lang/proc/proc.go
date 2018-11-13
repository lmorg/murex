package proc

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/pipes"
	"github.com/lmorg/murex/lang/proc/runmode"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
)

// Process - Each process running inside the murex shell will be one of these objects.
// It is equivalent to the /proc directory on Linux, albeit queried through murex as JSON.
// External processes will also appear in the host OS's process list.
type Process struct {
	Context            context.Context
	Stdin              stdio.Io
	Stdout             stdio.Io
	Stderr             stdio.Io
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
	Done               func()    `json:"-"`
	Kill               func()    `json:"-"`
	IsNot              bool
	NamedPipeOut       string
	NamedPipeErr       string
	NamedPipeTest      string
	hasTerminatedM     sync.Mutex
	hasTerminatedV     bool
	State              state.FunctionState
	IsBackground       bool
	LineNumber         int
	ColNumber          int
	RunMode            runmode.RunMode
	Config             *config.Config
	Tests              *Tests
	Variables          *Variables
	FidTree            []int
	CreationTime       time.Time
	StartTime          time.Time
}

var (
	// ShellProcess is the root murex process
	ShellProcess = &Process{}

	// MxFunctions is a table of global murex functions
	MxFunctions = NewMurexFuncs()

	// GoFunctions is a table of available builtin functions
	GoFunctions = make(map[string]func(*Process) error)

	// This will hold all variables
	masterVarTable = newVarTable()

	// InitConf is a table of global config options
	InitConf = config.NewConfiguration()

	// GlobalAliases is a table of global aliases
	GlobalAliases = NewAliases()

	// GlobalPipes is a table of  named pipes
	GlobalPipes = pipes.NewNamed()

	// GlobalFIDs is a table of running murex processes
	GlobalFIDs = *newFuncID()

	// ForegroundProc is the murex FID which currently has "focus"
	ForegroundProc = ShellProcess
)

// HasTerminated checks if process has terminated.
// This is a function because terminated state can be subject to race conditions
// so we need a mutex to make the state thread safe.
func (p *Process) HasTerminated() (state bool) {
	p.hasTerminatedM.Lock()
	state = p.hasTerminatedV
	p.hasTerminatedM.Unlock()
	return
}

// HasCancelled is a wrapper function around context because it's a pretty ugly API
func (p *Process) HasCancelled() (state bool) {
	if p.Context == nil {
		return false
	}

	select {
	case <-p.Context.Done():
		return true
	default:
		return false
	}
}

// SetTerminatedState sets the process terminated state.
// This is a function because terminated state can be subject to race conditions
// so we need a mutex to make the state thread safe.
func (p *Process) SetTerminatedState(state bool) {
	p.hasTerminatedM.Lock()
	p.hasTerminatedV = state
	p.hasTerminatedM.Unlock()
	return
}

// ErrIfNotAMethod returns a standard error message for builtins not run as methods
func (p *Process) ErrIfNotAMethod() (err error) {
	if !p.IsMethod {
		err = errors.New("`" + p.Name + "` expects to be pipelined.")
	}
	return
}

// DeregisterProcess deregisters a murex process
func DeregisterProcess(p *Process) {
	p.State = state.Terminating

	p.Stdout.Close()
	p.Stderr.Close()

	p.SetTerminatedState(true)
	if !p.IsBackground {
		ForegroundProc = p.Next
	}

	go deregister(p)
}

// deregister FID and mark variables for garbage collection.
func deregister(p *Process) {
	p.State = state.AwaitingGC
	CloseScopedVariables(p)
	GlobalFIDs.Deregister(p.Id)
}
