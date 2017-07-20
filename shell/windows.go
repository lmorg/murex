// +build windows

package shell

import (
	"io/ioutil"
	"os/user"
	"strings"
)

var HomeDirectory string

func init() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	HomeDirectory = usr.HomeDir + `\`
}

func splitPath(envPath string) []string {
	split := strings.Split(envPath, ";")
	return split
}

func listExes(path string, exes *map[string]bool) {
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		name := strings.ToLower(f.Name())
		if name == ".exe" || name == ".com" || name == ".bat" || name == ".cmd" || name == ".scr" {
			(*exes)[f.Name()] = true
		}
	}
	return
}

func matchExes(s string, exes *map[string]bool, includeColon bool) (items []string) {
	var colon string
	if includeColon {
		colon = ": "
	}

	for name := range *exes {
		lc := strings.ToLower(s)
		if strings.HasPrefix(strings.ToLower(name), lc) {
			switch name {
			case ">", ">>", "[", "=":
				items = append(items, name[len(s):]+" ")
			case "<read-pipe>":
			default:
				items = append(items, name[len(s):]+": ")
			}
		}
	}
	sort.Strings(items)
	return
}

func matchDirs(path string, dirs []string) (items []string) {
	/*items = []string{"../"}
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			items = append(items, f.Name()+`\`)
		}
	}

	for i := range dirs {
		lc := strings.ToLower(path)
		if strings.HasPrefix(strings.ToLower(dirs[path]), lc) {
			items = append(items, name[len(path):])
		}
	}
	return*/
	return []string{"[TODO: write me]"}
}

func matchFiles(path string, dirs []string) (items []string) {
	return []string{"[TODO: write me]"}
}

func isLocal(s string) bool {
	return strings.HasPrefix(s, `.\`) || strings.HasPrefix(s, `\`) || strings.HasPrefix(s[1:], `:\`)
}

func partialPath(loc string) (path, partial string) {
	split := strings.Split(loc, `\`)
	path = strings.Join(split[:len(split)-1], `\`)
	partial = split[len(split)-1]
	return
}

func matchLocal(loc string, includeColon bool) (items []string) {
	return []string{"[TODO: write me]"}
}
