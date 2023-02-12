//go:build !js
// +build !js

package home

import (
	"os/user"

	"github.com/lmorg/murex/lang/tty"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/consts"
)

// MyDir is the $USER directory.
// Typically /home/$USER/ on non-Windows systems, or \users\$USER on Windows.
var MyDir string

func init() {
	usr, err := user.Current()
	if err != nil {
		tty.Stderr.WriteString(err.Error() + utils.NewLineString)
		MyDir = consts.PathSlash
		return
	}

	MyDir = usr.HomeDir
}

// UserDir is the home directory of a `username`.
func UserDir(username string) string {
	usr, err := user.Lookup(username)
	if err != nil {
		return consts.PathSlash
	}

	return usr.HomeDir
}
