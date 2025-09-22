package which

import (
	"os"

	"github.com/lmorg/murex/utils/consts"
)

// Which works similarly to the UNIX command with the same name.
// If the executable is now found in $PATH then a zero length string is returned.
func Which(cmd string) string {
	_, err := os.Stat(cmd)
	if !os.IsNotExist(err) {
		return cmd
	}

	envPath := os.Getenv("PATH")

	for _, path := range SplitPath(envPath) {
		filepath := path + consts.PathSlash + cmd
		_, err := os.Stat(filepath)
		if !os.IsNotExist(err) {
			return filepath
		}
	}

	return ""
}

// WhichIgnoreFail will always return a best guess of the executable
func WhichIgnoreFail(cmd string) string {
	path := Which(cmd)
	if path == "" {
		return cmd
	}

	return path
}
