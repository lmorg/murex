// +build !windows

package signals

import (
	"os"
	"os/signal"
	"syscall"
)

// Handler is an internal function to capture and handle OS signals (eg SIGTERM).
func Handler(interactive bool) {
	/*defer func() {
		if r := recover(); r != nil {
			os.Stderr.WriteString(fmt.Sprintln("Exception caught: ", r))
			Handler(interactive)
		}
	}()*/

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGTSTP)
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
