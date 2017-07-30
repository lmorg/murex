package proc

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"sync"
)

type Flow struct {
	PipeOut  bool
	PipeErr  bool
	NewChain bool
	Last     bool
}

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

func (p *Process) HasTerminated() (state bool) {
	p.hasTerminatedM.Lock()
	state = p.hasTerminatedV
	p.hasTerminatedM.Unlock()
	return
}

func (p *Process) SetTerminatedState(state bool) {
	p.hasTerminatedM.Lock()
	p.hasTerminatedV = state
	p.hasTerminatedM.Unlock()
	return
}

type GoFunction struct {
	Func    func(*Process) error `json:"-"`
	TypeIn  string
	TypeOut string
}

var (
	ShellProcess   *Process              = &Process{Name: "$SHELL"}
	MxFunctions    MurexFuncs            = NewMurexFuncs()
	GoFunctions    map[string]GoFunction = make(map[string]GoFunction)
	GlobalVars     types.Vars            = types.NewVariableGroup()
	GlobalConf     config.Config         = config.NewConfiguration()
	GlobalAliases  Aliases               = NewAliases()
	GlobalPipes    Named                 = NewNamed()
	GlobalFIDs     funcID                = newFuncID()
	KillForeground func()                = func() {}
	ForegroundProc *Process              = ShellProcess
)
