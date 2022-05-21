package autocomplete

import (
	"strings"
	"time"

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
	DoNotSort         bool
	TimeOut           time.Time
}

func (act *AutoCompleteT) append(items ...string) {
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
	case width < 30:
		act.MinTabItemLength = 15
	case width < 80:
		act.MinTabItemLength = 20
	case width < 120:
		act.MinTabItemLength = 30
	case width < 160:
		act.MinTabItemLength = 40
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

	//sort.Strings(items)
	return
}

// MatchFlags is the entry point for murex's complex system of flag matching
func MatchFlags(act *AutoCompleteT) {
	if act.ParsedTokens.ExpectParam || len(act.ParsedTokens.Parameters) == 0 {
		act.ParsedTokens.Parameters = append(act.ParsedTokens.Parameters, "")
	}

	args := dynamicArgs{
		exe:    act.ParsedTokens.FuncName,
		params: make([]string, len(act.ParsedTokens.Parameters)),
	}
	copy(args.params, act.ParsedTokens.Parameters)

	params := make([]string, len(act.ParsedTokens.Parameters))
	copy(params, act.ParsedTokens.Parameters)

	flags := ExesFlags[act.ParsedTokens.FuncName]

	partial := act.ParsedTokens.Parameters[len(act.ParsedTokens.Parameters)-1]
	exe := act.ParsedTokens.FuncName

	pIndex := 0

	act.TimeOut = time.Now().Add(1 * time.Second)

	occurrences = 0
	matchFlags(flags, 0, partial, exe, params, &pIndex, args, act)
}
