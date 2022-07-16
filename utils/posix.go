//go:build !windows
// +build !windows

package utils

var (
	// NewLineByte is a new line in POSIX systems (no carriage return) as []byte's
	NewLineByte []byte = []byte{'\n'}
	// NewLineString is a new line in POSIX systems (no carriage return) as a string
	NewLineString string = "\n"
)
