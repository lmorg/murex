// +build windows

package utils

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/lmorg/murex/utils/consts"
)

var rxPathPrefix *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z]:[\\/]`)

// NormalisePath takes a string and returns an absolute path if the string is a
// relative path
func NormalisePath(path string) string {
	pwd, err := os.Getwd()
	if err == nil && path[0] != consts.PathSlash[0] && !rxPathPrefix.MatchString(path) {
		path = pwd + consts.PathSlash + path
	}

	path = filepath.Clean(path)

	return path
}
