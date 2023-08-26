package lang

import (
	"context"
	"os"
	"sync/atomic"
	"time"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/builtins/pipes/null"
	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/runmode"
	"github.com/lmorg/murex/lang/state"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/json"
)

var (
	// FlagTry is true if murex was started with `--try`
	FlagTry bool

	// FlagTryPipe is true if murex was started with `--trypipe`
	FlagTryPipe bool

	hasMurexAlreadyBeenInitialised int32 = -1
)

// InitEnv initialises murex. Exported function to enable unit tests.
func InitEnv() {
	i := atomic.AddInt32(&hasMurexAlreadyBeenInitialised, 1)
	if i > 0 {
		return
	}

	ShellProcess.CreationTime = time.Now()
	ShellProcess.State.Set(state.Executing)
	ShellProcess.Name.Set(os.Args[0])
	ShellProcess.Parameters.DefineParsed(os.Args[1:])
	ShellProcess.Scope = ShellProcess
	ShellProcess.Parent = ShellProcess
	ShellProcess.Previous = ShellProcess
	ShellProcess.Next = ShellProcess
	ShellProcess.Config = config.InitConf
	ShellProcess.Tests = NewTests(ShellProcess)
	ShellProcess.Variables = NewVariables(ShellProcess)
	ShellProcess.Stdout = new(term.Out)
	ShellProcess.Stderr = term.NewErr(true) // TODO: check this is overridden by `config set ...`
	ShellProcess.FileRef = ref.NewModule(app.ShellModule)
	ShellProcess.Context = context.Background()
	ShellProcess.Done = func() { /* we don't want to accidentally terminate the shell process */ }
	ShellProcess.Kill = func() { /* we don't want to accidentally terminate the shell process */ }
	ShellProcess.Forks = NewForkManagement()

	if FlagTry {
		ShellProcess.RunMode = runmode.ModuleTry
	}

	if FlagTryPipe {
		ShellProcess.RunMode = runmode.ModuleTryPipe
	}

	// Sets $SHELL to be murex
	shellEnv, err := os.Executable()
	if err != nil {
		shellEnv = ShellProcess.Name.String()
	}
	os.Setenv("SHELL", shellEnv)

	if os.Getenv(consts.EnvMethod) == consts.EnvTrue {
		ShellProcess.Stdin = term.NewIn(os.Getenv(consts.EnvDataType))
		ShellProcess.IsMethod = true
	}

	if os.Getenv(consts.EnvBackground) == consts.EnvTrue {
		ShellProcess.Background.Set(true)
	}

	// Pre-populate $PWDHIST with current working directory
	s, _ := os.Getwd()
	pwd := []string{s}
	if b, err := json.Marshal(&pwd, false); err == nil {
		GlobalVariables.Set(ShellProcess, "PWDHIST", string(b), types.Json)
	}

	if err = MethodStdin.Degroup(); err != nil {
		panic(err.Error())
	}
	if err = MethodStdout.Degroup(); err != nil {
		panic(err.Error())
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
	p.Forks = NewForkManagement()

	GlobalFIDs.Register(p)

	return
}
