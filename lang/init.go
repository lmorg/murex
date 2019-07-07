package lang

import (
	"context"
	"os"
	"time"

	"github.com/lmorg/murex/lang/ref"

	"github.com/lmorg/murex/builtins/pipes/null"
	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/runmode"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/json"
)

// InitEnv initialises murex. Exported function to enable unit tests.
func InitEnv() {
	ShellProcess.State = state.Executing
	ShellProcess.Name = os.Args[0]
	ShellProcess.Parameters.Params = os.Args[1:]
	ShellProcess.Scope = ShellProcess
	ShellProcess.Parent = ShellProcess
	ShellProcess.Config = InitConf.Copy()
	ShellProcess.Tests = NewTests(ShellProcess)
	ShellProcess.Variables = newVariables(ShellProcess, masterVarTable)
	ShellProcess.RunMode = runmode.Shell
	ShellProcess.FidTree = []uint32{0}
	ShellProcess.Stdout = new(term.Out)
	ShellProcess.Stderr = term.NewErr(true) // TODO: check this is overridden by `config set ...`

	ShellProcess.FileRef = &ref.File{Source: &ref.Source{Module: config.AppName}}

	// Sets $SHELL to be murex
	shellEnv, err := utils.Executable()
	if err != nil {
		shellEnv = ShellProcess.Name
	}
	os.Setenv("SHELL", shellEnv)

	// Pre-populate $PWDHIST with current working directory
	s, _ := os.Getwd()
	pwd := []string{s}
	if b, err := json.Marshal(&pwd, false); err == nil {
		ShellProcess.Variables.Set("PWDHIST", string(b), types.Json)
	}
}

// NewTestProcess creates a dummy process for testing in Go (ie `go test`)
func NewTestProcess() (p *Process) {
	p = new(Process)
	p.Stdin = new(null.Null)
	p.Stdout = new(null.Null)
	p.Stderr = new(null.Null)
	p.Config = config.NewConfiguration()
	p.Variables = newVariables(p, masterVarTable)
	p.Context, p.Done = context.WithTimeout(context.Background(), 60*time.Second)

	GlobalFIDs.Register(p)

	return
}
