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
		if f.IsDir() {
			continue
		}
		perm := permbits.FileMode(f.Mode())
		if perm.OtherExecute() /*|| perm.GroupExecute()||perm.UserExecute() need to check what user and groups you are in first */ {
			(*exes)[f.Name()] = true
		}
	}
}

func matchExes(s string, exes *map[string]bool) (items []string) {
	for name := range *exes {
		if strings.HasPrefix(name, s) {
			items = append(items, name[len(s):])
		}
	}
	sort.Strings(items)
	for i := range items {
		items[i] += ": "
	}
	return
}

func isLocal(s string) bool {
	return strings.HasPrefix(s, "./") || strings.HasPrefix(s, "/")
}

func partialPath(loc string) (path, partial string) {
	split := strings.Split(loc, "/")
	path = strings.Join(split[:len(split)-1], "/")
	partial = split[len(split)-1]
	if path == "" {
		path = "."
	}
	return
}

func matchLocal(loc string) (items []string) {
	path, file := partialPath(loc)
	exes := make(map[string]bool)
	listExes(path, &exes)
	items = matchExes(file, &exes)
	return
}

func matchDirs(loc string) (items []string) {
	path, partial := partialPath(loc)

	dirs := []string{"../"}
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			dirs = append(dirs, f.Name()+"/")
		}
	}

	for i := range dirs {
		if strings.HasPrefix(dirs[i], partial) {
			items = append(items, dirs[i][len(partial):])
		}
	}
	return
}
