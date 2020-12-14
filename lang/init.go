package lang

import (
	"context"
	"os"
	"time"

	"github.com/lmorg/murex/builtins/pipes/null"
	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/runmode"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

// InitEnv initialises murex. Exported function to enable unit tests.
func InitEnv() {
	ShellProcess.State.Set(state.Executing)
	ShellProcess.Name = os.Args[0]
	ShellProcess.Parameters.Params = os.Args[1:]
	ShellProcess.Scope = ShellProcess
	ShellProcess.Parent = ShellProcess
	ShellProcess.Previous = ShellProcess
	ShellProcess.Next = ShellProcess
	ShellProcess.Config = config.InitConf //InitConf.Copy(ShellProcess.Id)
	ShellProcess.Tests = NewTests(ShellProcess)
	ShellProcess.Variables = NewVariables(ShellProcess)
	ShellProcess.RunMode = runmode.Shell
	ShellProcess.Stdout = new(term.Out)
	ShellProcess.Stderr = term.NewErr(true) // TODO: check this is overridden by `config set ...`
	ShellProcess.Kill = func() {}
	ShellProcess.FileRef = &ref.File{Source: &ref.Source{Module: config.AppName}}

	// Sets $SHELL to be murex
	shellEnv, err := os.Executable()
	if err != nil {
		shellEnv = ShellProcess.Name
	}
	os.Setenv("SHELL", shellEnv)

	// Pre-populate $PWDHIST with current working directory
	s, _ := os.Getwd()
	pwd := []string{s}
	if b, err := json.Marshal(&pwd, false); err == nil {
		//ShellProcess.Variables.Set("PWDHIST", string(b), types.Json)
		GlobalVariables.Set(ShellProcess, "PWDHIST", string(b), types.Json)
	}
}

// NewTestProcess creates a dummy process for testing in Go (ie `go test`)
func NewTestProcess() (p *Process) {
	p = new(Process)
	p.Stdin = new(null.Null)
	p.Stdout = new(null.Null)
	p.Stderr = new(null.Null)
	p.Config = config.InitConf.Copy()
	p.Variables = NewVariables(p)
	p.FileRef = &ref.File{Source: &ref.Source{Module: "builtin/testing"}}
	p.Context, p.Done = context.WithTimeout(context.Background(), 60*time.Second)
	p.Parent = ShellProcess
	p.Scope = ShellProcess
	p.Next = ShellProcess
	p.Previous = ShellProcess

	GlobalFIDs.Register(p)

	return
}
