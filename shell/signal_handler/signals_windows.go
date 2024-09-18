//go:build windows
// +build windows

package signalhandler

import (
	"os"
	"os/signal"
	"syscall"
)

// EventLoop is an internal function to capture and handle OS signals (eg SIGTERM).
func EventLoop(interactive bool) {
	Register(interactive)

	go func() {
		for {
			sig := <-signalChan
			switch sig.String() {

			case syscall.SIGINT.String():
				Handlers.Sigint(interactive)

			case syscall.SIGTERM.String():
				Handlers.Sigterm(interactive)

			case syscall.SIGQUIT.String():
				Handlers.Sigquit(interactive)

			default:
				os.Stderr.WriteString("Unhandled signal: " + sig.String())
			}
		}
	}()
}

func Register(_ bool) {
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
}

var Handlers *SignalFunctionsT

type SignalFunctionsT struct {
	Sigint  func(bool)
	Sigterm func(bool)
	Sigquit func(bool)
}
