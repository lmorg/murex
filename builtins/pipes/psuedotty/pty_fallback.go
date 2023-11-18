//go:build windows || js || plan9 || no_pty
// +build windows js plan9 no_pty

package psuedotty

import (
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/stdio"
)

func init() {
	stdio.RegisterPipe("pty", registerPipe)
}

func registerPipe(_ string) (stdio.Io, error) {
	return streams.NewStdin(), nil
}

type PTY struct {
	streams.Stdin
}

func NewPTY(width, height int) (*PTY, error) {
	pty := PTY{*streams.NewStdin()}
	return &pty, nil
}
