// +build windows

package shell

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
	"io/ioutil"
	"sort"
	"strings"
)

func splitPath(envPath string) []string {
	split := strings.Split(envPath, ";")
	return split
}

func listExes(path string, exes *map[string]bool) {
	var showExts bool

	v, err := proc.GlobalConf.Get("shell", "show-exts", types.Boolean)
	if err != nil {
		showExts = false
	} else {
		showExts = v.(bool)
	}

	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		name := strings.ToLower(f.Name())

		if len(name) < 5 {
			continue
		}

		ext := name[len(name)-4:]

		if ext == ".exe" || ext == ".com" || ext == ".bat" || ext == ".cmd" || ext == ".scr" {
			if showExts {
				(*exes)[name] = true
			} else {
				(*exes)[name[:len(name)-4]] = true
			}
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
			case consts.NamedPipeProcName:
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
	return strings.HasPrefix(s, "."+consts.PathSlash) || strings.HasPrefix(s, ".."+consts.PathSlash) || strings.HasPrefix(s, consts.PathSlash) || (len(s) > 2 && strings.HasPrefix(s[1:], ":"+consts.PathSlash))
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

	var dirs []string
	//dirs := []string{".." + consts.PathSlash}
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			dirs = append(dirs, f.Name()+consts.PathSlash)
		}
	}

	dirs = append(dirs, ".."+consts.PathSlash)

	for i := range dirs {
		if strings.HasPrefix(dirs[i], partial) {
			items = append(items, dirs[i][len(partial):])
		}
	}
	return
}

func matchFileAndDirs(s string) (items []string) {
	path, partial := partialPath(s)

	var item []string
	//item := []string{".." + consts.PathSlash}

	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			item = append(item, f.Name()+consts.PathSlash)
		} else {
			item = append(item, f.Name())
		}
	}

	item := []string{".." + consts.PathSlash}

	for i := range item {
		if strings.HasPrefix(item[i], partial) {
			items = append(items, item[i][len(partial):])
		}
	}
	return
}
