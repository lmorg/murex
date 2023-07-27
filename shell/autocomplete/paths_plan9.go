//go:build plan9
// +build plan9

package autocomplete

import (
	"io/ioutil"
	"strings"

	"github.com/lmorg/murex/utils/consts"
)

func pathIsLocal(s string) bool {
	return strings.HasPrefix(s, consts.PathSlash) ||
		strings.HasPrefix(s, "."+consts.PathSlash) ||
		strings.HasPrefix(s, ".."+consts.PathSlash)
}

func matchDirsOnce(s string) (items []string) {
	//s = variables.ExpandString(s)
	path, partial := partialPath(s)

	var dirs []string

	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() && (f.Name()[0] != '.' ||
			(len(partial) > 0 && partial[0] == '.')) {
			dirs = append(dirs, f.Name()+consts.PathSlash)
			continue
		}

		// TODO: There is a log of logic missing in here that appears in the unix source
	}

	if path != consts.PathSlash {
		dirs = append(dirs, ".."+consts.PathSlash)
	}

	for i := range dirs {
		if strings.HasPrefix(dirs[i], partial) {
			items = append(items, dirs[i][len(partial):])
		}
	}

	return
}
