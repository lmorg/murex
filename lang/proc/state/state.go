package state

//go:generate stringer -type=FunctionStates

type FunctionStates int

const (
	Undefined FunctionStates = iota
	MemAllocated
	Assigned
	Starting
	Executing
	Executed
	Terminating
	AwaitingGC
)
