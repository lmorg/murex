//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package lang

import (
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/creack/pty"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/tty"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/readline"
)

func ttys(p *Process) {
	p.ttyout = tty.Stdout

	if p.CCExists != nil && p.CCExists(p.Name.String()) {
		p.Stderr, p.CCErr = streams.NewTee(p.Stderr)
		p.CCErr.SetDataType(types.Generic)

		p.Stdout, p.CCOut = streams.NewTee(p.Stdout)
		if p.Stdout.IsTTY() {
			ptyout, ttyout, err := pty.Open()
			if err != nil {
				return
			}

			_ = pty.InheritSize(tty.Stdout, ptyout)
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, syscall.SIGWINCH)
			go func() {
				for range ch {
					_ = pty.InheritSize(tty.Stdout, ptyout)
				}
			}()

			_, err = readline.MakeRaw(int(ptyout.Fd()))
			if err != nil {
				return
			}

			p.ttyout = ptyout
			go func() {
				_, _ = io.Copy(p.Stdout, ttyout)
				signal.Stop(ch)
				close(ch)
			}()
		}

	}
}
