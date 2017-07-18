// +build !windows

package shell

import (
	"github.com/lmorg/murex/utils/permbits"
	"io/ioutil"
	"os/user"
	"sort"
	"strings"
)

var HomeDirectory string

func init() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	HomeDirectory = usr.HomeDir + "/"
}

func splitPath(envPath string) []string {
	split := strings.Split(envPath, ":")
	return split
}

func listExes(path string, exes *map[string]bool) {
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		perm := permbits.FileMode(f.Mode())
		if perm.OtherExecute() /*|| perm.GroupExecute()||perm.UserExecute() need to check what user and groups you are in first */ {
			(*exes)[f.Name()] = true
		}
	}
}

func matchExes(s string) (items []string) {
	for name := range allExecutables() {
		if strings.HasPrefix(name, s) {
			items = append(items, name[len(s):]+": ")
		}
	}
	sort.Strings(items)
	return
}
