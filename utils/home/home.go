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
	}

	MyDir = usr.HomeDir + slash
}
