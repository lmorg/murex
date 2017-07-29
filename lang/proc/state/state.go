package state

//go:generate stringer -type=FunctionStates

type FunctionStates int

const (
	Unknown FunctionStates = iota
	Created
	Running
	Terminating
	AwaitingGC
)
