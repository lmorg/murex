//go:build js
// +build js

package main

import (
	"syscall/js"

	"github.com/lmorg/murex/app"
	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/readline/v4"
)

const interactive = true

func main() {
	startMurex()

	js.Global().Set("wasmShellExec", wasmShellExec())
	js.Global().Set("wasmShellStart", wasmShellStart())
	js.Global().Set("wasmKeyPress", wasmKeyPress())

	wait := make(chan bool)
	<-wait
}

func startMurex() {
	lang.InitEnv()

	// default config
	defaults.Config(lang.ShellProcess.Config, interactive)

	// compiled profile
	profile.Execute(profile.F_BUILTIN)
}

// wasmShellExec returns a Promise
func wasmShellExec() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		block := args[0].String()

		// Handler for the Promise: this is a JS function
		// It receives two arguments, which are JS functions themselves: resolve and reject
		handler := js.FuncOf(func(this js.Value, args []js.Value) any {
			resolve := args[0]
			reject := args[1]

			// Now that we have a way to return the response to JS, spawn a goroutine
			// This way, we don't block the event loop and avoid a deadlock
			go func() {
				fork := lang.ShellProcess.Fork(lang.F_PARENT_VARTABLE | lang.F_NEW_MODULE | lang.F_NO_STDIN)
				fork.FileRef.Source.Module = app.ShellModule
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

// wasmShellStart starts the interactive shell as a Promise
func wasmShellStart() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {

		handler := js.FuncOf(func(this js.Value, args []js.Value) any {
			resolve := args[0]
			//reject := args[1]

			go func() {
				resolve.Invoke("Starting interactive shell....")
				shell.Start()
			}()

			// The handler of a Promise doesn't return any value
			return nil
		})

		// Create and return the Promise object
		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}

// wasmKeyPress starts the interactive shell as a Promise
func wasmKeyPress() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		stdin := args[0].String()

		handler := js.FuncOf(func(this js.Value, args []js.Value) any {
			//resolve := args[0]
			//reject := args[1]

			go func() {
				readline.Stdin <- stdin
			}()

			// The handler of a Promise doesn't return any value
			return nil
		})

		// Create and return the Promise object
		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}

/*
func registerSignalHandlers(interactiveMode bool) {
	signalhandler.Handlers = &signalhandler.SignalFunctionsT{}
	signalhandler.EventLoop(interactiveMode)
}*/
