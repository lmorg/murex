//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package signalhandler

import (
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
				go Handlers.Sigint(interactive)

			case syscall.SIGTERM.String():
				go Handlers.Sigterm(interactive)

			case syscall.SIGQUIT.String():
				go Handlers.Sigquit(interactive)

			case syscall.SIGTSTP.String():
				go Handlers.Sigtstp(interactive)

			case syscall.SIGCHLD.String():
				go Handlers.Sigchld(interactive)

			default:
				panic("unhandled signal: " + sig.String()) // this shouldn't ever happen
			}
		}
	}()
}

func Register(interactive bool) {
	if interactive {
		// Interactive, so we will handle stop
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGTSTP) //, syscall.SIGCHLD) //, syscall.SIGTTIN, syscall.SIGTTOU)

	} else {
		// Non-interactive, so lets ignore the stop signal and let the OS / calling shell manage that for us
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	}
}

var Handlers *SignalFunctionsT

type SignalFunctionsT struct {
	Sigint  func(bool)
	Sigterm func(bool)
	Sigquit func(bool)
	Sigtstp func(bool)
	Sigchld func(bool)
}
