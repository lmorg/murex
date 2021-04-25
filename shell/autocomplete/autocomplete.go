package autocomplete

import (
	"sort"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/parser"
	"github.com/lmorg/murex/utils/readline"
)

// AutoCompleteT is a struct designed for ease to pass common values around the
// many functions for autocompletion. It's passed as a pointer and is only
// intended for use by murex internal functions (ie not called by other funcs
// external to the murex codebase)
type AutoCompleteT struct {
	Items             []string
	Definitions       map[string]string
	MinTabItemLength  int
	TabDisplayType    readline.TabDisplayType
	ErrCallback       func(error)
	DelayedTabContext readline.DelayedTabContext
	ParsedTokens      parser.ParsedTokens
	CacheDynamic      bool
}

func (act *AutoCompleteT) append(items ...string) {
	/*// Dedup
	for _, item := range items {
		for i := range act.Items {
			if act.Items[i] == item {
				goto next
			}
		}

		act.Items = append(act.Items, item)
	next:
	}*/

	act.Items = append(act.Items, items...)
}

func (act *AutoCompleteT) appendDef(item, def string) {
	act.Definitions[item] = def
	act.append(item)
}

func (act *AutoCompleteT) largeMin() {
	width := readline.GetTermWidth()
	switch {
	case width < 40:
		act.MinTabItemLength = 10
	case width < 80:
		act.MinTabItemLength = 15
	case width < 120:
		act.MinTabItemLength = 20
	case width < 160:
		act.MinTabItemLength = 30
	default:
		act.MinTabItemLength = 40
	}
}

func (act *AutoCompleteT) disposable() *AutoCompleteT {
	return &AutoCompleteT{
		Items:             []string{},
		Definitions:       make(map[string]string),
		ErrCallback:       act.ErrCallback,
		DelayedTabContext: act.DelayedTabContext,
		ParsedTokens:      act.ParsedTokens,
	}
}

// MatchFunction returns autocomplete suggestions for functions / executables
// based on a partial string
func MatchFunction(partial string, act *AutoCompleteT) (items []string) {
	switch {
	case pathIsLocal(partial):
		items = matchLocal(partial, true)
		items = append(items, matchDirs(partial, act)...)
	default:
		exes := allExecutables(true)
		items = matchExes(partial, exes, true)
	}
	return
}

// MatchVars returns autocomplete suggestions for variables based on a partial
// string
func MatchVars(partial string) (items []string) {
	vars := lang.DumpVariables(lang.ShellProcess)

	for name := range vars {
		if strings.HasPrefix(name, partial[1:]) {
			items = append(items, name[len(partial)-1:])
		}
	}

	sort.Strings(items)
	return
}

// MatchFlags is the entry point for murex's complex system of flag matching
func MatchFlags(flags []Flags, partial, exe string, params []string, pIndex *int, act *AutoCompleteT) int {
	args := dynamicArgs{
		exe:    exe,
		params: params,
	}

	return matchFlags(flags, partial, exe, params, pIndex, args, act)
}
