//go:build windows && plan9 && js
// +build windows,plan9,js

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
	p.Stdout, p.CCOut = streams.NewTee(p.Stdout)
	p.ttyout = os.Stdout

	p.Stderr, p.CCErr = streams.NewTee(p.Stderr)
	p.CCErr.SetDataType(types.Generic)
}
