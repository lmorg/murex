// +build windows

package autocomplete

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/variables"
	"github.com/lmorg/murex/utils/consts"
	"io/ioutil"
	"sort"
	"strings"
)

func splitPath(envPath string) []string {
	split := strings.Split(envPath, ";")
	return split
}

func pathIsLocal(s string) bool {
	return strings.HasPrefix(s, "."+consts.PathSlash) || strings.HasPrefix(s, ".."+consts.PathSlash) || strings.HasPrefix(s, consts.PathSlash) || (len(s) > 2 && strings.HasPrefix(s[1:], ":"+consts.PathSlash))
}

func matchDirs(s string) (items []string) {
	s = variables.ExpandString(s)
	path, partial := partialPath(s)

	var dirs []string

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
