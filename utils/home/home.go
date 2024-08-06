//go:build !js
// +build !js

package home

import (
	"os"
	"os/user"

	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/consts"
)

// MyDir is the $USER directory.
// Typically /home/$USER/ on non-Windows systems, or \users\$USER on Windows.
var MyDir string

func init() {
	usr, err := user.Current()
	if err != nil {
		os.Stderr.WriteString(err.Error() + utils.NewLineString)
		MyDir = consts.PathSlash
		return
	}

	MyDir = usr.HomeDir
}
