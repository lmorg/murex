package autocomplete

import (
	"sort"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/readline"
)

type AutoCompleteT struct {
	Items             []string
	Definitions       map[string]string
	TabDisplayType    readline.TabDisplayType
	ErrCallback       func(error)
	DelayedTabContext readline.DelayedTabContext
}

// MatchFunction returns autocomplete suggestions for functions / executables
// based on a partial string
func MatchFunction(partial string, errCallback func(error), dtc *readline.DelayedTabContext) (items []string) {
	switch {
	case pathIsLocal(partial):
		items = matchLocal(partial, true)
		items = append(items, matchDirs(partial, errCallback, dtc)...)
	default:
		exes := allExecutables(true)
		items = matchExes(partial, exes, true)
	}
	return
}

// MatchVars returns autocomplete suggestions for variables based on a partial
// string
func MatchVars(partial string) (items []string) {
	vars := lang.ShellProcess.Variables.DumpMap()

	for name := range vars {
		if strings.HasPrefix(name, partial[1:]) {
			items = append(items, name[len(partial)-1:])
		}
	}

	sort.Strings(items)
	return
}

// MatchFlags is the entry point for murex's complex system of flag matching
func MatchFlags(flags []Flags, partial, exe string, params []string, pIndex *int, defs *map[string]string, tdt *readline.TabDisplayType, errCallback func(error), dtc *readline.DelayedTabContext) (items []string) {
	args := dynamicArgs{
		exe:    exe,
		params: params,
	}

	return matchFlags(flags, partial, exe, params, pIndex, args, defs, tdt, errCallback, dtc)
}
