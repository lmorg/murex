package runmode

//go:generate stringer -type=RunMode

// RunMode is the type for defining the murex interpreters run mode
type RunMode int

// These are the different supported run modes
const (
	Default RunMode = iota

	Normal

	BlockUnsafe
	FunctionUnsafe
	ModuleUnsafe

	BlockTry
	BlockTryPipe
	BlockTryErr
	BlockTryPipeErr

	FunctionTry
	FunctionTryPipe
	FunctionTryErr
	FunctionTryPipeErr

	ModuleTry
	ModuleTryPipe
	ModuleTryErr
	ModuleTryPipeErr
)

// IsStrict checks if RunMode is a Try or TryPipe block
func (i RunMode) IsStrict() bool {
	if i > ModuleUnsafe {
		return true
	}

	return false
}
