package consts

import (
	"io/ioutil"

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
}
