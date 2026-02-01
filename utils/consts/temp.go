package consts

import (
	"os"

	"github.com/lmorg/murex/app"
)

// temporaryDirectory is the location of tmp directory
var temporaryDirectory string

func init() {
	var err error

	temporaryDirectory, err = os.MkdirTemp("", app.Name)
	if err != nil || temporaryDirectory == "" {
		temporaryDirectory = _TMP_DIR
	}

	if temporaryDirectory[len(temporaryDirectory)-1:] != PathSlash {
		temporaryDirectory += PathSlash
	}

	createDirIfNotExist(temporaryDirectory)
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			_, err = os.Stderr.WriteString("!!! WARNING: temp directory doesn't exist and unable to create it. This might cause problems.\nTemp directory: " + dir)

			if err != nil {
				panic("Unable to create tmp directories, unable to write to STDERR. Something is amiss")
			}
		}
	}
}

func TmpDir() string {
	return temporaryDirectory
}
