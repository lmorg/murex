package proc

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/streams"
	"github.com/lmorg/murex/lang/types"
)

type Process struct {
	Stdin      streams.Io
	Stdout     streams.Io
	Stderr     streams.Io
	Parameters parameters.Parameters
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
	MethodRef  string
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
)
