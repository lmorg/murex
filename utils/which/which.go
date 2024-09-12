package which

import (
	"os"

	"github.com/lmorg/murex/utils/consts"
)

// Which works similarly to the UNIX command with the same name
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

func WhichIgnoreFail(cmd string) string {
	path := Which(cmd)
	if path == "" {
		return cmd
	}

	return path
}
