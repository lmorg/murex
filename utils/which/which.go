package which

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/lmorg/murex/utils/consts"
)

// Which works similarly to the UNIX command with the same name.
// If the executable is not found in $PATH then a zero length string is returned.
func Which(cmd string) string {
	if strings.ContainsRune(cmd, '/') || strings.ContainsRune(cmd, filepath.Separator) {
		// cmd is explicitly a path, resolve and return the absolute path if possible
		fi, err := os.Stat(cmd)
		if err == nil && fi.Mode().IsRegular() {
			if abs, err := filepath.Abs(cmd); err == nil {
				return abs
			}
			return cmd
		}
		return ""
	} else {
		// cmd is not explicitly a path, just search $PATH, don't attempt to reoslve absolute path
		for _, path := range SplitPath(os.Getenv("PATH")) {
			fullPath := path + consts.PathSlash + cmd
			fi, err := os.Stat(fullPath)
			if err == nil && fi.Mode().IsRegular() {
				return fullPath
			}
		}

		return ""
	}
}

// WhichIgnoreFail will always return a best guess of the executable
func WhichIgnoreFail(cmd string) string {
	path := Which(cmd)
	if path == "" {
		return cmd
	}

	return path
}
