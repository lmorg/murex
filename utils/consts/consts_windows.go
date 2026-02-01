//go:build windows
// +build windows

package consts

const (
	// PathSlash is an OS specific directory separator.
	// Normally in Windows this would be a \ but lets standardize everything in murex to be /
	PathSlash = "/"

	// _TMP_DIR is the location of temp directory if it cannot be automatically determined
	_TMP_DIR = `c:/temp/murex/`
)
