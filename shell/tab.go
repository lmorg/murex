package shell

import (
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/readline"
)

func errCallback(err error) {
	s := err.Error()
	if ansi.IsAllowed() {
		s = ansi.Reset + ansi.FgRed + s
	}
	Prompt.SetHintText(s)
}

func tabCompletion(line []rune, pos int, dtc readline.DelayedTabContext) (string, []string, map[string]string, readline.TabDisplayType) {
	var prefix string

	if len(line) > pos-1 {
		line = line[:pos]
	}

	pt, _ := parse(line)

	act := autocomplete.AutoCompleteT{
		Definitions:       make(map[string]string),
		ErrCallback:       errCallback,
		DelayedTabContext: dtc,
		ParsedTokens:      pt,
	}

	switch {
	case pt.Variable != "":
		if pt.VarLoc < len(line) {
			prefix = strings.TrimSpace(string(line[pt.VarLoc:]))
		}
		prefix = pt.Variable + prefix
		act.Items = autocomplete.MatchVars(prefix)

	case pt.ExpectFunc:
		if pt.Loc < len(line) {
			prefix = strings.TrimSpace(string(line[pt.Loc:]))
		}
		act.Items = autocomplete.MatchFunction(prefix, &act)

	default:
		if len(pt.Parameters) > 0 {
			prefix = pt.Parameters[len(pt.Parameters)-1]
		}
		autocomplete.InitExeFlags(pt.FuncName)

		pIndex := 0
		autocomplete.MatchFlags(autocomplete.ExesFlags[pt.FuncName], prefix, pt.FuncName, pt.Parameters, &pIndex, &act)
	}

	v, err := lang.ShellProcess.Config.Get("shell", "max-suggestions", types.Integer)
	if err != nil {
		v = 4
	}
	Prompt.MaxTabCompleterRows = v.(int)

	autocomplete.FormatSuggestions(&act)
	return prefix, act.Items, act.Definitions, act.TabDisplayType
}
