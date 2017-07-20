// +build windows

package shell

import (
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
	colon := " "
	if includeColon {
		colon = ": "
	}

	for name := range *exes {
		lc := strings.ToLower(s)
		if strings.HasPrefix(strings.ToLower(name), lc) {
			switch name {
			case ">", ">>", "[", "=":
				items = append(items, name[len(s):])
			case "<read-pipe>":
			default:
				items = append(items, name[len(s):])
			}
		}
	}
	sort.Strings(items)
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
	return strings.HasPrefix(s, `.\`) || strings.HasPrefix(s, `\`) || (len(s) > 2 && strings.HasPrefix(s[1:], `:\`))
}

func partialPath(loc string) (path, partial string) {
	split := strings.Split(loc, `\`)
	path = strings.Join(split[:len(split)-1], `\`)
	partial = split[len(split)-1]
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

	dirs := []string{`..\`}
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			dirs = append(dirs, f.Name()+`\`)
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

	item := []string{`..\`}
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			item = append(item, f.Name()+`\`)
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
