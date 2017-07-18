// +build windows

package shell

import (
	"os/user"
	"strings"
)

var HomeDirectory string

func init() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	HomeDirectory = usr.HomeDir + "\\"
}

func splitPath(envPath string) []string {
	split := strings.Split(envPath, ";")
	return split
}

func listExes(path string, exes *map[string]bool) {
	for _, f := range files {
		name := strings.ToLower(f.Name())
		if name == ".exe" || name == ".com" || name == ".bat" || name == ".cmd" || name == ".scr" {
			(*exes)[f.Name()] = true
		}
	}
	return
}

func matchExes(s string) (items []string) {
	for name := range allExecutables() {
		lc := strings.ToLower(s)
		if strings.HasPrefix(strings.ToLower(name), lc) {
			items = append(items, name[len(s):]+": ")
		}
	}
	sort.Strings(items)
	return
}
