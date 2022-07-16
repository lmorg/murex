//go:build !windows
// +build !windows

package consts

const (
	// PathSlash is an OS specific directory separator
	PathSlash = "/"

	// tempDir is the location of temp directory if it cannot be automatically determind
	tempDir = "/tmp/murex/"
)
