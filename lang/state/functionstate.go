package state

//go:generate stringer -type=FunctionState

// FunctionState is what the point along the murex pipeline a lang.Process is at
type FunctionState int32

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
	Stopped
)
