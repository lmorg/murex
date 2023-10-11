//go:build js
// +build js

package signaltrap

type syscall string

func (s syscall) String() string { return "" }

var interrupts = map[string]syscall{}
>>>>>>> lambda
