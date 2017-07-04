// +build windows

package shell

import "os/user"

var HomeDirectory string

func init() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	HomeDirectory = usr.HomeDir + "\\"
}
