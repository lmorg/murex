package runmode

//go:generate stringer -type=RunMode

// RunMode is the type for defining the murex interpreters run mode
type RunMode int

// These are the different supported run modes
const (
	Normal RunMode = iota
	Try
	TryPipe
	Shell
)
