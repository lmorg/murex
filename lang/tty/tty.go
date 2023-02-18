package tty

import (
	"os"
)

var (
	Stdin  *os.File = os.Stdin
	Stdout *os.File = os.Stdout
	Stderr *os.File = os.Stderr

	NewLine = []byte("↵\r\n")
)
