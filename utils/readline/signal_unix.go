//go:build !windows && !js
// +build !windows,!js

package readline

import (
	"os"
	"os/signal"
	"syscall"
)

func (rl *Instance) sigwinch() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			print("\r" + seqUp + seqClearScreenBelow + seqDown)
			print(rl.prompt + string(rl.line))
			rl.updateHelpers()
		}
	}()

	rl.closeSigwinch = func() {
		signal.Stop(ch)
		close(ch)
	}
}
