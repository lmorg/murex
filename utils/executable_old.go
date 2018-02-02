// +build !go1.8

package utils

import (
	"github.com/lmorg/murex/config"
	"os"
)

func Executable() (string, error) {
	return os.Args[0]
}
