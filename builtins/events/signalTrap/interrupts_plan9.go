//go:build plan9
// +build plan9

package signaltrap

import "syscall"

var interrupts = []syscall.Signal{}
