//go:build plan9
// +build plan9

package signaltrap

type syscall string

func (s syscall) String() string { return "" }

var interrupts = map[string]syscall{}
>>>>>>> lambda
