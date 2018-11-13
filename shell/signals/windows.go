// +build windows

package signals

import (
	"os"
	"os/signal"
	"syscall"
)

// Handler is an internal function to capture and handle OS signals (eg SIGTERM).
func Handler(interactive bool) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
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

			default:
				os.Stderr.WriteString("Unhandled signal: " + sig.String())
			}
		}
	}()
}
