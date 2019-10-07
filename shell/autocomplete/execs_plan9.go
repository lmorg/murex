// +build plan9

package autocomplete

import (
	"io/ioutil"
	"sort"
	"strings"

	"github.com/lmorg/murex/utils/consts"
)

// SplitPath takes a $PATH (or similar) string and splits it into an array of paths
func SplitPath(envPath string) []string {
	split := strings.Split(envPath, ":")
	return split
}

func listExes(path string, exes map[string]bool) {
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		// TODO: There is a log of logic missing in here that appears in the unix source
		exes[f.Name()] = true
	}
}

func matchExes(s string, exes map[string]bool, includeColon bool) (items []string) {
	colon := ""
	// We only want a colon added if the exe is the function call rather than a
	// functions parameter (eg `some-exec` vs `sudo some-exec`).
	if includeColon {
		colon = ":"
	}

	for name := range exes {
		if strings.HasPrefix(name, s) {
			if name != consts.NamedPipeProcName {
				items = append(items, name[len(s):])
			}
		}
	}
	sort.Strings(items)

	// I know it seems weird and inefficient to cycle through the array after
	// it has been created just to append a couple of characters (that easily
	// could have been appended in the former for loop) but this is so that the
	// colon isn't included as part of the sorting algorithm (eg otherwise
	// `manpath:` would precede `man:`). Ideally I would write my own sorting
	// function to take this into account but that can be part of the
	// optimisation stage - whenever I get there.
	for i := range items {
		if !isSpecialBuiltin(items[i]) {
			items[i] += colon
		}
	}
	return
}
