package home

import (
	"github.com/lmorg/murex/utils"
	"os"
	"os/user"
)

// MyDir is the $USER directory. Typically /home/$USER/ on non-Windows systems, or \users\$USER on Windows.
var MyDir string

func init() {
	usr, err := user.Current()
	if err != nil {
		os.Stderr.WriteString(err.Error() + utils.NewLineString)
		MyDir = PathSlash
		return
	}

	MyDir = usr.HomeDir
}

// UserDir is the home directory of a `username`.
func UserDir(username string) string {
	usr, err := user.Lookup(username)
	if err != nil {
		return PathSlash
	}

	return usr.HomeDir
}
