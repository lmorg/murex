// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build solaris

package tty

import (
	"golang.org/x/sys/unix"
)

// State contains the state of a terminal.
type State struct {
	state *unix.Termios
}

// MakeRaw puts the terminal connected to the given file descriptor into raw
// mode and returns the previous state of the terminal so that it can be
// restored.
// see http://cr.illumos.org/~webrev/andy_js/1060/
func MakeRaw(fd int) (*State, error) {
	oldTermiosPtr, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	if err != nil {
		return nil, err
	}
	oldTermios := *oldTermiosPtr

	newTermios := oldTermios
	newTermios.Cflag &^= unix.ISIG | unix.VQUIT | unix.VINTR | unix.VSUSP

	if err := unix.IoctlSetTermios(fd, unix.TCSETS, &newTermios); err != nil {
		return nil, err
	}

	return &State{
		state: oldTermiosPtr,
	}, nil
}

// Restore restores the terminal connected to the given file descriptor to a
// previous state.
func Restore(fd int, oldState *State) error {
	return unix.IoctlSetTermios(fd, unix.TCSETS, oldState.state)
}
