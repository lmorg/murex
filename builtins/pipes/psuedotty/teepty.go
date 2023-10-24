package psuedotty

import (
	"github.com/lmorg/murex/builtins/pipes/streams"
)

func NewTeePTY(width int, height int) (pty *PTY, primary *streams.Tee, secondary *streams.Stdin, err error) {
	pty, err = NewPTY(width, height)
	if err != nil {
		return nil, nil, nil, err
	}

	primary, secondary = streams.NewTee(pty)

	return pty, primary, secondary, nil
}
