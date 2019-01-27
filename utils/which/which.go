package which

import (
	"os"

	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils/consts"
)

// Which works similarly to the UNIX command with the same name
func Which(cmd string) string {
	envPath := os.Getenv("PATH")

	for _, path := range autocomplete.SplitPath(envPath) {
		filepath := path + consts.PathSlash + cmd
		_, err := os.Stat(filepath)
		if !os.IsNotExist(err) {
			return filepath
		}
	}

	return ""
}
