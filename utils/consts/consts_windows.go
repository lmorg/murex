// +build windows

package consts

const (
	// PathSlash is an OS specific directory separator.
	// Normally in Windows this would be a \ but lets standardise everything in murex to be /
	PathSlash = "/"

	// tempDir is the location of temp directory if it cannot be automatically determind
	tempDir = `c:/temp/murex/`
)
