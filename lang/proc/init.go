package proc

import (
	"os"

	"github.com/lmorg/murex/lang/proc/runmode"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

// InitEnv initialises murex. Exported function to enable unit tests.
func InitEnv() {
	ShellProcess.State = state.Executing
	ShellProcess.Name = os.Args[0]
	ShellProcess.Parameters.Params = os.Args[1:]
	ShellProcess.Scope = ShellProcess
	ShellProcess.Parent = ShellProcess
	ShellProcess.Config = InitConf.Copy()
	ShellProcess.Variables = &Variables{varTable: masterVarTable, process: ShellProcess}
	ShellProcess.RunMode = runmode.Shell
	ShellProcess.FidTree = []int{0}
	ShellProcess.Stdout = new(streams.TermOut)
	ShellProcess.Stderr = new(streams.TermErr)

	// Sets $SHELL to be murex
	shellEnv, err := utils.Executable()
	if err != nil {
		shellEnv = ShellProcess.Name
	}
	os.Setenv("SHELL", shellEnv)

	// Pre-populate $PWDHIST with current working directory
	s, _ := os.Getwd()
	pwd := []string{s}
	//if b, err := json.MarshalIndent(&pwd, "", "    "); err == nil {
	if b, err := utils.JsonMarshal(&pwd, false); err == nil {
		ShellProcess.Variables.Set("PWDHIST", string(b), types.Json)
	}
}
