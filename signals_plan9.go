//go:build plan9
// +build plan9

package main

import (
	signalhandler "github.com/lmorg/murex/shell/signal_handler"
	"github.com/lmorg/murex/shell/signal_handler/sigfns"
)

func registerSignalHandlers(interactiveMode bool) {
	signalhandler.Handlers = &signalhandler.SignalFunctionsT{
		Sigint:  sigfns.Sigint,
		Sigterm: sigfns.Sigterm,
	}
	signalhandler.EventLoop(interactiveMode)
}
