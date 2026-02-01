//go:build !windows
// +build !windows

package consts

const (
	// PathSlash is an OS specific directory separator
	PathSlash = "/"

	// _TMP_DIR is the location of temp directory if it cannot be automatically determined
	_TMP_DIR = "/tmp/murex/"
)
