package readline

import (
	"errors"
)

const (
	_CtrlC = "Ctrl+C"
	_EOF   = "EOF"
)

var (
	// CtrlC is returned when ctrl+c is pressed
	ErrCtrlC = errors.New(_CtrlC)

	// EOF is returned when ctrl+d is pressed.
	// (this is actually the same value as io.EOF)
	ErrEOF = errors.New(_EOF)
)
