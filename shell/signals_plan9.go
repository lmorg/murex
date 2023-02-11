//go:build plan9
// +build plan9

package shell

import (
	"os/signal"
	"syscall"

	"github.com/lmorg/murex/lang/tty"
)

// SignalHandler is an internal function to capture and handle OS signals (eg SIGTERM).
func SignalHandler(interactive bool) {
	signalRegister(interactive)

	go func() {
		for {
			sig := <-signalChan
			switch sig.String() {

			case syscall.SIGINT.String():
				sigint(interactive)

			case syscall.SIGTERM.String():
				sigterm(interactive)

			default:
				tty.Stderr.WriteString("Unhandled signal: " + sig.String())
			}
		}
	}()
}

func signalRegister(_ bool) {
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
}
