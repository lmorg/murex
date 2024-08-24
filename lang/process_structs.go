package lang

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/process"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/runmode"
	"github.com/lmorg/murex/lang/state"
	"github.com/lmorg/murex/lang/stdio"
)

// Process - Each process running inside the murex shell will be one of these objects.
// It is equivalent to the /proc directory on Linux, albeit queried through murex as JSON.
// External processes will also appear in the host OS's process list.
type Process struct {
	Id                 uint32
	cache              *cacheT
	raw                []rune
	Name               process.Name
	Parameters         parameters.Parameters
	namedPipes         []string
	Context            context.Context
	Stdin              stdio.Io
	Stdout             stdio.Io
	stdoutOldPtr       stdio.Io // only used when stdout is a tmp named pipe
	Stderr             stdio.Io
	ExitNum            int
	Forks              *ForkManagement
	IsFork             bool
	WaitForTermination chan bool `json:"-"`
	WaitForStopped     chan bool `json:"-"`
	HasStopped         chan bool `json:"-"`
	Done               func()    `json:"-"`
	Kill               func()    `json:"-"`
	SystemProcess      process.SystemProcess
	Envs               []string
	Scope              *Process `json:"-"`
	Parent             *Process `json:"-"`
	Previous           *Process `json:"-"`
	Next               *Process `json:"-"`
	IsNot              bool
	IsMethod           bool
	OperatorLogicAnd   bool
	OperatorLogicOr    bool
	NamedPipeOut       string
	NamedPipeErr       string
	NamedPipeTest      string
	hasTerminatedM     sync.Mutex `json:"-"`
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
	CCEvent            func(string, *Process) `json:"-"`
	CCExists           func(string) bool      `json:"-"`
	CCOut              *streams.Stdin         `json:"-"`
	CCErr              *streams.Stdin         `json:"-"`
	Trace              []any
}

func (p *Process) Dump() map[string]any {
	return map[string]any{
		"Id":           p.Id,
		"Cache":        p.cache != nil && p.cache.use,
		"Raw":          string(p.raw),
		"Name":         p.Name.String(),
		"Parameters":   p.Parameters.Dump(),
		"NamedPipes":   p.namedPipes,
		"Stdin":        statsToStruct(p.Stdin),
		"Stdout":       statsToStruct(p.Stdout),
		"Stderr":       statsToStruct(p.Stderr),
		"stdoutOldPtr": statsToStruct(p.stdoutOldPtr),
		"ExitNum":      p.ExitNum,
		"Forks":        p.Forks.GetForks(),
		"IsFork":       p.IsFork,
		//"Done":             p.Done,
		//"Kill":             p.Kill,
		"SysProcId":        _jsonfySysProc(p),
		"ScopeFid":         p.Scope.Id,
		"ParentFid":        p.Parent.Id,
		"PreviousFid":      p.Previous.Id,
		"NextFid":          p.Next.Id,
		"IsNot":            p.IsNot,
		"IsMethod":         p.IsMethod,
		"OperatorLogicAnd": p.OperatorLogicAnd,
		"OperatorLogicOr":  p.OperatorLogicOr,
		"NamedPipeOut":     p.NamedPipeOut,
		"NamedPipeErr":     p.NamedPipeErr,
		"NamedPipeTest":    p.NamedPipeTest,
		"HasTerminated":    p.HasTerminated(),
		"HasCancelled":     p.HasCancelled(),
		//"HasStopped":       p.hasCancelledStopped(),
		"State":         p.State.String(),
		"Background":    p.Background.Get(),
		"RunMode":       p.RunMode.String(),
		"RunModeStrict": p.RunMode.IsStrict(),
		//Config             *config.Config
		//Tests              *Tests
		//testState          []string
		//Variables          *Variables
		"CreationTime": p.CreationTime.String(),
		"StartTime":    p.StartTime.String(),
		"FileRef":      p.FileRef,
		//CCEvent            func(string, *Process) `json:"-"`
		//CCExists           func(string) bool      `json:"-"`
		//CCOut              *streams.Stdin         `json:"-"`
		//CCErr              *streams.Stdin         `json:"-"`
		"Trace": p.Trace,
	}
}

func _jsonfySysProc(p *Process) any {
	if p.SystemProcess == nil {
		return nil
	}

	return p.SystemProcess.Pid()
}

type statsT struct {
	Read    uint64
	Written uint64
}

func statsToStruct(stdio stdio.Io) any {
	if stdio == nil {
		return nil
	}

	r, w := stdio.Stats()
	return statsT{
		Read:    r,
		Written: w,
	}
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
func (p *Process) HasCancelled() bool {
	select {
	case <-p.Context.Done():
		return true

	default:
		if p.State.Get() == state.Stopped {
			return p.hasCancelledStopped()
		}
		return false
	}
}

func (p *Process) hasCancelledStopped() bool {
	select {
	case <-p.Context.Done():
		return true
	case <-p.WaitForStopped:
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

func (p *Process) KillForks(exitNum int) {
	forks := p.Forks.GetForks()
	for _, procs := range forks {
		for i := range *procs {
			(*procs)[i].ExitNum = exitNum
			(*procs)[i].Done()
		}
	}
}

// Args returns a normalised function name and parameters
func (p *Process) Args() (string, []string) {
	return args(p.Name.String(), p.Parameters.StringArray())
}

func args(name string, params []string) (string, []string) {
	if len(params) == 0 {
		return name, []string{}
	}

	switch name {
	case "exec":
		return params[0], params[1:]

	default:
		return name, params
	}
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
