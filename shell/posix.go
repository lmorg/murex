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

func matchExes(s string, exes *map[string]bool, includeColon bool) (items []string) {
	colon := " "
	if includeColon {
		colon = ": "
	}

	for name := range *exes {
		if strings.HasPrefix(name, s) {
			if name != "<read-pipe>" {
				items = append(items, name[len(s):])
			}
		}
	}
	sort.Strings(items)

	// I know it seems weird and inefficient to cycle through the array after it has been created just to append a
	// couple of characters (that easily could have been appended in the former for loop) but this is so that the
	// colon isn't included as part of the sorting algo (otherwise `manpath:` would precede `man:`). Ideally I would
	// write my own sorting algo to take this into account but that can be part of the optimisation stage - whenever
	// I get there.
	for i := range items {
		switch items[i] {
		case ">", ">>", "[", "=":
			items[i] += " "
		default:
			items[i] += colon
		}
	}
	return
}

func isLocal(s string) bool {
	return strings.HasPrefix(s, "./") || strings.HasPrefix(s, "/")
}

func partialPath(s string) (path, partial string) {
	split := strings.Split(s, "/")
	path = strings.Join(split[:len(split)-1], "/")
	partial = split[len(split)-1]

	if len(s) > 0 && s[0] == '/' {
		path = "/" + path
	}

	if path == "" {
		path = "."
	}
	return
}

func matchLocal(s string, includeColon bool) (items []string) {
	path, file := partialPath(s)
	exes := make(map[string]bool)
	listExes(path, &exes)
	items = matchExes(file, &exes, includeColon)
	return
}

func matchDirs(s string) (items []string) {
	path, partial := partialPath(s)

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

func matchFileAndDirs(s string) (items []string) {
	path, partial := partialPath(s)

	item := []string{"../"}
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			item = append(item, f.Name()+"/")
		} else {
			item = append(item, f.Name())
		}
	}

	for i := range item {
		if strings.HasPrefix(item[i], partial) {
			items = append(items, item[i][len(partial):])
		}
	}
	return
}
