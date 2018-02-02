package autocomplete

import (
	"github.com/lmorg/murex/lang/proc"
	"os"
	"sort"
	"strings"
)

// MatchFunction returns autocomplete suggestions for functions / executables based on a partial string
func MatchFunction(partial string) (items []string) {
	switch {
	case pathIsLocal(partial):
		items = matchLocal(partial, true)
		items = append(items, matchDirs(partial)...)
	default:
		exes := allExecutables(true)
		items = matchExes(partial, exes, true)
	}
	return
}

// MatchVars returns autocomplete suggestions for variables based on a partial string
func MatchVars(partial string) (items []string) {
	vars := proc.GlobalVars.DumpMap()

	envVars := os.Environ()
	for i := range envVars {
		v := strings.Split(envVars[i], "=")
		vars[v[0]] = true
	}

	for name := range vars {
		if strings.HasPrefix(name, partial[1:]) {
			items = append(items, name[len(partial)-1:])
		}
	}
	sort.Strings(items)
	return
}
