package runmode

//go:generate stringer -type=RunMode

// RunMode is the type for defining the murex interpreters run mode
type RunMode int

// These are the different supported run modes
const (
	Normal RunMode = iota
	Evil
	BlockTry
	BlockTryPipe
	FunctionTry
	FunctionTryPipe
	ModuleTry
	ModuleTryPipe
)

// IsStrict checks if RunMode is a Try or TryPipe block
func (i RunMode) IsStrict() bool {
	if i > Evil {
		return true
	}

	return false
}

func (i RunMode) IsScopeFunction() bool {
	if i > BlockTryPipe {
		return true
	}

	return false
}
