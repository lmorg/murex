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
	var colon string
	if includeColon {
		colon = ": "
	}

	for name := range *exes {
		if strings.HasPrefix(name, s) {
			switch name {
			case ">", ">>", "[", "=":
				items = append(items, name[len(s):]+" ")
			case "<read-pipe>":
			default:
				items = append(items, name[len(s):]+colon)
			}
		}
	}
	sort.Strings(items)
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

func matchFileAndDirs(loc string) (items []string) {
	path, partial := partialPath(loc)

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
