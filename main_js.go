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
	startMurex()

	js.Global().Set("wasmShellExec", wasmShellExec())

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

// wasmShellExec returns a Promise
func wasmShellExec() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		block := args[0].String()

		// Handler for the Promise: this is a JS function
		// It receives two arguments, which are JS functions themselves: resolve and reject
		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			reject := args[1]

			// Now that we have a way to return the response to JS, spawn a goroutine
			// This way, we don't block the event loop and avoid a deadlock
			go func() {
				fork := lang.ShellProcess.Fork(lang.F_PARENT_VARTABLE | lang.F_NEW_MODULE | lang.F_NO_STDIN)
				fork.FileRef.Source.Module = app.Name
				fork.Stderr = term.NewErr(ansi.IsAllowed())
				var err error
				lang.ShellExitNum, err = fork.Execute([]rune(block))
				if err != nil {
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
				}
				resolve.Invoke("wasmShellExec(): " + block)
			}()

			// The handler of a Promise doesn't return any value
			return nil
		})

		// Create and return the Promise object
		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}
