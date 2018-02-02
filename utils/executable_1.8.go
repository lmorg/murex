// +build go1.8

package utils

import "os"

func Executable() (string, error) {
	return os.Executable()
}
