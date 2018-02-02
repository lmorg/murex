// +build !go1.8

package utils

import (
	"os"
)

// Executable is rubbish stand in function for os.Executable for compatibility with Go versions < 1.8
func Executable() (string, error) {
	return os.Args[0], nil
}
