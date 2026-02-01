//go:build windows
// +build windows

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
	}
	signalhandler.EventLoop(interactiveMode)
}
