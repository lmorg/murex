package runmode

//go:generate stringer -type=RunMode

// RunMode is the type for defining the murex interpreters run mode
type RunMode int

// These are the different supported run modes
const (
	Default RunMode = iota

	Normal
	Evil

	BlockTry
	BlockTryPipe

	ModuleTry
	ModuleTryPipe

	FunctionTry
	FunctionTryPipe
)

// IsStrict checks if RunMode is a Try or TryPipe block
func (i RunMode) IsStrict() bool {
	if i > Evil {
		return true
	}

	return false
}
