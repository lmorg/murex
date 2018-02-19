package proc

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/pipes"
	"github.com/lmorg/murex/lang/proc/runmode"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
	"github.com/lmorg/murex/lang/types"

	"errors"
	"os"
	"strings"
	"sync"
)

// Process - Each process running inside the murex shell will be one of these objects.
// It is equivalent to the /proc directory on Linux, albeit queried through murex as JSON.
// External processes will also appear in the host OS's process list.
type Process struct {
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
	RunMode            runmode.RunMode
	Config             *config.Config
	ScopedVars         *types.Vars
}

var (
	// ShellProcess is the root murex process
	ShellProcess *Process = &Process{}

	// MxFunctions is a table of global murex functions
	MxFunctions MurexFuncs = NewMurexFuncs()

	// GoFunctions is a table of available builtin functions
	GoFunctions map[string]func(*Process) error = make(map[string]func(*Process) error)

	// InitConf is a table of global config options
	InitConf *config.Config = config.NewConfiguration()

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

// ErrIfNotAMethod returns a standard error message for builtins not run as methods
func (p *Process) ErrIfNotAMethod() (err error) {
	if !p.IsMethod {
		err = errors.New("`" + p.Name + "` must be run as a method.")
	}
	return
}

// VarGetValue returns a Go-typed value for the variable with the requested name.
// If no local scoped variable exists, then value is taken from globals
// (ShellProcess) and if no variables exist from there then value is taken from
// the OS environmental variables
func (p *Process) VarGetValue(name string) interface{} {
	v := p.ScopedVars.GetValue(name)
	if v != nil {
		return v
	}

	if p.Id != ShellProcess.Id {
		return ShellProcess.VarGetValue(name)

	} else {
		s, b := os.LookupEnv(name)
		if !b {
			return nil
		}
		return s
	}
}

// VarGetString returns a string value for the variable with the requested name.
// If no local scoped variable exists, then value is taken from globals
// (ShellProcess) and if no variables exist from there then value is taken from
// the OS environmental variables
func (p *Process) VarGetString(name string) string {
	//defer func() {
	//r := recover()
	//if r != nil {
	//	p.Stderr.Writeln([]byte("Error converting `" + name + "` to string: " + err.Error()))
	//}
	//v.mutex.Unlock()
	//}()
	v := p.VarGetValue(name)
	str, err := types.ConvertGoType(v, types.String)
	if err != nil {
		//ansi.Stderrln(ansi.FgRed, "Error converting `"+name+"` to string: "+err.Error())
		p.Stderr.Writeln([]byte("Error converting `" + name + "` to string: " + err.Error()))
	}
	return str.(string)
}

// VarGetType returns the murex type for the variable with the requested name.
// If no local scoped variable exists, then value is taken from globals
// (ShellProcess) and if no variables exist from there then value is assumed a
// string (if an OS environmental variable exists) or null (if no env var exists)
func (p *Process) VarGetType(name string) string {
	v := p.ScopedVars.GetValue(name)
	if v != nil {
		return p.ScopedVars.GetType(name)
	}

	if p.Id != ShellProcess.Id {
		return ShellProcess.VarGetType(name)

	} else {
		_, b := os.LookupEnv(name)
		if !b {
			return types.Null
		}
	}
	return types.String
}

// VarDumpMap returns a map of all the variables (murex and environmental)
// visible in scope
func (p *Process) VarDumpMap() map[string]interface{} {
	m := make(map[string]interface{})

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		m[pair[0]] = pair[1]
	}

	if p.Id != ShellProcess.Id {
		ShellProcess.ScopedVars.DumpMap(m)
	}

	p.ScopedVars.DumpMap(m)

	return m
}
