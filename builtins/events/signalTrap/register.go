//go:build !plan9 && !js
// +build !plan9,!js

package signaltrap

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var signalChan chan os.Signal = make(chan os.Signal, 1)

func register(sig string) error {
	for name := range interrupts {
		if name == sig {
			signal.Notify(signalChan, interrupts[name])
			return nil
		}
	}

	return fmt.Errorf("no signal found named '%s'", sig)
}

func deregister(sig syscall.Signal) {
	signal.Reset(sig)
}
