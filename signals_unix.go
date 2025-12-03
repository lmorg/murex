//go:build !js && !windows && !plan9
// +build !js,!windows,!plan9

package main

import (
	signalhandler "github.com/lmorg/murex/shell/signal_handler"
	"github.com/lmorg/murex/shell/signal_handler/sigfns"
)

func registerSignalHandlers(interactiveMode bool) {
	signalhandler.Handlers = &signalhandler.SignalFunctionsT{
		Sigint:  sigfns.Sigint,
		Sigterm: sigfns.Sigterm,
		Sigquit: sigfns.Sigquit,
		Sigtstp: sigfns.Sigtstp,
		Sigchld: sigfns.Sigchld,
		Sigcont: sigfns.Sigcont,
	}
	signalhandler.EventLoop(interactiveMode)
}
