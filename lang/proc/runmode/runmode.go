package runmode

//go:generate stringer -type=FunctionStates

// RunMode is the type for defining the murex interpreters run mode
type RunMode int

// These are the different supported run modes
const (
	Normal RunMode = iota
	Try
	TryPipe
	Evil
)
