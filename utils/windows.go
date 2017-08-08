// +build windows

package utils

var (
	// NewLineByte is a new line in Windows (carriage return, line feed) as []byte's
	NewLineByte []byte = []byte{'\r', '\n'}
	// NewLineString is a new line in Windows (carriage return, line feed) as a string
	NewLineString string = "\r\n"
)
