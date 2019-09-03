package consts

import (
	"io/ioutil"
	"os"

	"github.com/lmorg/murex/config"
)

// TempDir is the location of temp directory
var TempDir string

func init() {
	var err error

	TempDir, err = ioutil.TempDir("", config.AppName)
	if err != nil || TempDir == "" {
		TempDir = tempDir
	}

	if TempDir[len(TempDir)-1:] != PathSlash {
		TempDir += PathSlash
	}

	createDirIfNotExist(TempDir)
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		/*err = */ os.MkdirAll(dir, 0755)
		//if err != nil {
		//	panic(err)
		//}
	}
}
