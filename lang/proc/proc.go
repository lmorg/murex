package proc

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
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
	Parent             *Process `json:"-"`
	Previous           *Process `json:"-"`
	Next               *Process `json:"-"`
	HasTerminated      bool
	WaitForTermination chan bool `json:"-"`
	IsNot              bool
	MethodRef          string
	ReturnType         string
}

type GoFunction struct {
	Func    func(*Process) error
	TypeIn  string
	TypeOut string
}

var (
	GlobalVars   types.Vars            = types.NewVariableGroup()
	GoFunctions  map[string]GoFunction = make(map[string]GoFunction)
	GlobalConf   config.Config         = config.NewConfiguration()
	ShellEnabled bool
)
