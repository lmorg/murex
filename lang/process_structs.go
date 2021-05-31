package lang

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/runmode"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/process"
	"github.com/lmorg/murex/lang/ref"
)

// Process - Each process running inside the murex shell will be one of these objects.
// It is equivalent to the /proc directory on Linux, albeit queried through murex as JSON.
// External processes will also appear in the host OS's process list.
type Process struct {
	Id                 uint32
	Name               process.Name
	Parameters         parameters.Parameters
	Context            context.Context
	Stdin              stdio.Io
	Stdout             stdio.Io
	stdoutOldPtr       stdio.Io // only used when stdout is a tmp named pipe
	Stderr             stdio.Io
	ExitNum            int
	WaitForTermination chan bool `json:"-"`
	Done               func()    `json:"-"`
	Kill               func()    `json:"-"`
	Exec               process.Exec
	PromptId           int
	Scope              *Process `json:"-"`
	Parent             *Process `json:"-"`
	Previous           *Process `json:"-"`
	Next               *Process `json:"-"`
	IsNot              bool
	IsMethod           bool
	NamedPipeOut       string
	NamedPipeErr       string
	NamedPipeTest      string
	hasTerminatedM     sync.Mutex
	hasTerminatedV     bool
	State              state.State
	Background         process.Background
	RunMode            runmode.RunMode
	Config             *config.Config
	Tests              *Tests
	testState          []string
	Variables          *Variables
	CreationTime       time.Time
	StartTime          time.Time
	FileRef            *ref.File
}

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
}

// ErrIfNotAMethod returns a standard error message for builtins not run as methods
func (p *Process) ErrIfNotAMethod() error {
	if !p.IsMethod {
		return fmt.Errorf("`%s` expects to be pipelined", p.Name.String())
	}
	return nil
}

type foregroundProc struct {
	mutex sync.Mutex
	p     *Process
}

func newForegroundProc() *foregroundProc {
	return &foregroundProc{p: ShellProcess}
}

func (fp *foregroundProc) Get() *Process {
	fp.mutex.Lock()
	p := fp.p
	//if p == nil {
	//	panic("Get() retrieved p")
	//}
	fp.mutex.Unlock()

	return p
}

func (fp *foregroundProc) Set(p *Process) {
	fp.mutex.Lock()
	if p == nil {
		panic("nil p in (fp *foregroundProc) Set(p *Process)")
	}
	fp.p = p
	//debug.Json("fp.Set", p)
	//debug.Json("fp.p", fp.p)
	fp.mutex.Unlock()
}
