//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package lang

import (
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/creack/pty"
	"github.com/lmorg/murex/builtins/pipes/psuedotty"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/readline/v4"
)

func ttys(p *Process) {
	if p.CCExists != nil && p.CCExists(p.Name.String()) {
		p.Stderr, p.CCErr = streams.NewTee(p.Stderr)
		p.CCErr.SetDataType(types.Generic)

		if p.Stdout.IsTTY() {
			width, height, err := readline.GetSize(int(p.Stdout.File().Fd()))
			if err != nil {
				width, height = 80, 25
			}

			var tee stdio.Io
			p.Stdout, tee, p.CCOut, err = psuedotty.NewTeePTY(width, height)
			if err != nil {
				p.Stdout, p.CCOut = streams.NewTee(p.Stdout)
				return
			}

			ch := make(chan os.Signal, 1)
			signal.Notify(ch, syscall.SIGWINCH)
			go func() {
				for range ch {
					_ = pty.InheritSize(os.Stdout, p.Stdout.File())
				}
			}()

			go func() {
				_, _ = io.Copy(os.Stdout, tee)
				signal.Stop(ch)
				close(ch)
			}()

		} else {
			p.Stdout, p.CCOut = streams.NewTee(p.Stdout)
		}
	}
}
