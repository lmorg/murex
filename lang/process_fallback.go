//go:build windows || plan9 || js
// +build windows plan9 js

package lang

import (
	"os"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/types"
)

func ttys(p *Process) {
	p.ttyin = tty.Stdin

	p.Stdout, p.CCOut = streams.NewTee(p.Stdout)
	p.ttyout = tty.Stdout

	p.Stderr, p.CCErr = streams.NewTee(p.Stderr)
	p.CCErr.SetDataType(types.Generic)
}
