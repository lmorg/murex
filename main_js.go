// +build js

package main

import (
	"syscall/js"

	"github.com/lmorg/murex/app"
	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/utils/ansi"
)

const interactive = true

func main() {
	js.Global().Set("shellExec", js.FuncOf(jsShellExec))

	startMurex()

	wait := make(chan bool)
	<-wait
}

func startMurex() {
	lang.InitEnv()

	// default config
	defaults.Defaults(lang.ShellProcess.Config, interactive)

	// compiled profile
	source := defaults.DefaultMurexProfile()
	ref := ref.History.AddSource("(builtin)", "source/builtin", []byte(string(source)))
	execSource(defaults.DefaultMurexProfile(), ref)

	// load modules and profile
	//profile.Execute()

	// start interactive shell
	//shell.Start()
}

func jsShellExec(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return "Invalid number of args passed"
	}

	block := args[0].String()

	fork := lang.ShellProcess.Fork(lang.F_PARENT_VARTABLE | lang.F_NEW_MODULE | lang.F_NO_STDIN)
	fork.FileRef.Source.Module = app.Name
	fork.Stderr = term.NewErr(ansi.IsAllowed())
	lang.ShellExitNum, _ = fork.Execute([]rune(block))

	return nil
}
