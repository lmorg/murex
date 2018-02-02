package autocomplete

import (
	"github.com/lmorg/murex/shell/vars"
	"github.com/lmorg/murex/utils/consts"
	"io/ioutil"
	"strings"
)

func partialPath(s string) (path, partial string) {
	expanded := vars.ExpandVariables([]rune(s))
	split := strings.Split(string(expanded), consts.PathSlash)
	path = strings.Join(split[:len(split)-1], consts.PathSlash)
	partial = split[len(split)-1]

	if len(s) > 0 && s[0] == consts.PathSlash[0] {
		path = consts.PathSlash + path
	}

	if path == "" {
		path = "."
	}
	return
}

func matchLocal(s string, includeColon bool) (items []string) {
	path, file := partialPath(s)
	exes := make(map[string]bool)
	listExes(path, exes)
	items = matchExes(file, exes, includeColon)
	return
}

func matchFilesAndDirs(s string) (items []string) {
	s = vars.ExpandVariablesString(s)
	path, partial := partialPath(s)

	var item []string

	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			item = append(item, f.Name()+consts.PathSlash)
		} else {
			item = append(item, f.Name())
		}
	}

	item = append(item, ".."+consts.PathSlash)

	for i := range item {
		if strings.HasPrefix(item[i], partial) {
			items = append(items, item[i][len(partial):])
		}
	}
	return
}
