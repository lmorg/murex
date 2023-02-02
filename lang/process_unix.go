//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package lang

import (
	"io"
	"os"

	"github.com/creack/pty"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/readline"
)

func ttys(p *Process) {
	p.ttyout = os.Stdout

	if p.CCExists != nil && p.CCExists(p.Name.String()) {
		p.Stderr, p.CCErr = streams.NewTee(p.Stderr)
		p.CCErr.SetDataType(types.Generic)

		p.Stdout, p.CCOut = streams.NewTee(p.Stdout)
		if p.Stdout.IsTTY() {
			ptyout, tty, err := pty.Open()
			if err != nil {
				//panic(err)
				return
			}

			size, err := pty.GetsizeFull(os.Stdout)
			if err == nil {
				_ = pty.Setsize(tty, size)
				_ = pty.Setsize(ptyout, size)
			}

			_, err = readline.MakeRaw(int(ptyout.Fd()))
			if err != nil {
				//panic(err)
				return
			}

			p.ttyout = ptyout
			go func() { _, _ = io.Copy(p.Stdout, tty) }()
		}

	}
}
