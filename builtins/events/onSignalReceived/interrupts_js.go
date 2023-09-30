//go:build js
// +build js

package signaltrap

import "syscall"

var interrupts = map[string]syscall.Signal{}
