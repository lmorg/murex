package readline

import (
	"golang.org/x/crypto/ssh/terminal"
)

// IsTerminal returns true if the given file descriptor is a terminal.
func IsTerminal(fd int) bool {
	return terminal.IsTerminal(fd)
}

// MakeRaw put the terminal connected to the given file descriptor into raw
// mode and returns the previous state of the terminal so that it can be
// restored.
func MakeRaw(fd int) (*terminal.State, error) {
	return terminal.MakeRaw(fd)
}

// GetState returns the current state of a terminal which may be useful to
// restore the terminal after a signal.
func GetState(fd int) (*terminal.State, error) {
	return terminal.GetState(fd)
}

// Restore restores the terminal connected to the given file descriptor to a
// previous state.
func Restore(fd int, state *terminal.State) error {
	return terminal.Restore(fd, state)
}

// GetSize returns the dimensions of the given terminal.
func GetSize(fd int) (width, height int, err error) {
	return terminal.GetSize(fd)
}
