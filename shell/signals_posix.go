// +build !windows

package shell

import (
	"os"
	"os/signal"
	"syscall"
)

// SignalHandler is an internal function to capture and handle OS signals (eg SIGTERM).
func SignalHandler(interactive bool) {
	c := make(chan os.Signal, 1)

	if Interactive {
		// Interactive, so we will handle suspend
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGTSTP)
	} else {
		// Non-interactive, so lets ignore the suspend signal and let the OS / calling shell manage that for us
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	}

	go func() {
		for {
			sig := <-c
			switch sig.String() {

			case syscall.SIGINT.String():
				sigint(interactive)

			case syscall.SIGTERM.String():
				sigterm(interactive)

			case syscall.SIGQUIT.String():
				sigquit(interactive)

			case syscall.SIGTSTP.String():
				sigtstp()

			default:
				os.Stderr.WriteString("Unhandled signal: " + sig.String())
			}
		}
	}()
}
