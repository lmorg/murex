package which

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
	}

	// cmd is not explicitly a path, just search $PATH
	path, err := exec.LookPath(cmd)
	if err != nil {
		return ""
	}
	return path
}

// WhichIgnoreFail will always return a best guess of the executable
func WhichIgnoreFail(cmd string) string {
	path := Which(cmd)
	if path == "" {
		return cmd
	}

	return path
}
