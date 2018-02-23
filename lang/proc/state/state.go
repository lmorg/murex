package state

//go:generate stringer -type=FunctionState

// FunctionState is what the point along the murex pipeline a proc.Process is at
type FunctionState int

// The different states available to FunctionState:
const (
	Undefined FunctionState = iota
	MemAllocated
	Assigned
	Starting
	Executing
	Executed
	Terminating
	AwaitingGC
)
