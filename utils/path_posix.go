//go:build !windows
// +build !windows

package utils

import (
	"os"
	"path/filepath"

	"github.com/lmorg/murex/utils/consts"
)

// NormalisePath takes a string and returns an absolute path if the string is a
// relative path
func NormalisePath(path string) string {
	// bugfix: https://github.com/lmorg/murex/issues/789
	if len(path) == 0 {
		path = "."
	}

	pwd, err := os.Getwd()
	if err == nil && path[0] != consts.PathSlash[0] {
		path = pwd + consts.PathSlash + path
	}

	path = filepath.Clean(path)

	return path
}
