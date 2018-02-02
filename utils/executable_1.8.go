// +build go1.8

package utils

import "os"

// Executable is a wrapper around os.Executable so I can retain compatibility with earlier versions of Go
func Executable() (string, error) {
	return os.Executable()
}
