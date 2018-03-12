package shell

import (
	"github.com/lmorg/murex/shell/autocomplete"
	"strings"
)

func tabCompleter(line []rune, pos int) (prefix string, items []string) {
	if len(line) > pos-1 {
		line = line[:pos]
	}

	pt, _ := parse(line)

	switch {
	case pt.Variable != "":
		var s string
		if pt.VarLoc < len(line) {
			s = strings.TrimSpace(string(line[pt.VarLoc:]))
		}
		s = pt.Variable + s
		//retPos = len(s)
		prefix = s

		items = autocomplete.MatchVars(s)

	case pt.ExpectFunc:
		var s string
		if pt.Loc < len(line) {
			s = strings.TrimSpace(string(line[pt.Loc:]))
		}
		//retPos = len(s)
		prefix = s
		items = autocomplete.MatchFunction(s)

	default:
		var s string
		if len(pt.Parameters) > 0 {
			s = pt.Parameters[len(pt.Parameters)-1]
		}
		//retPos = len(s)
		prefix = s

		autocomplete.InitExeFlags(pt.FuncName)

		pIndex := 0
		items = autocomplete.MatchFlags(autocomplete.ExesFlags[pt.FuncName], s, pt.FuncName, pt.Parameters, &pIndex)
	}

	/*v, err := proc.ShellProcess.Config.Get("shell", "max-suggestions", types.Integer)
	if err != nil {
		v = -1
	}

	limitSuggestions := v.(int)
	if len(items) < limitSuggestions || limitSuggestions < 0 {
		limitSuggestions = len(items)
	}
	Instance.Config.MaxCompleteLines = limitSuggestions

	suggest = make([][]rune, len(items))
	for i := range items {
		if len(items[i]) == 0 {
			continue
		}

		if !pt.QuoteSingle && !pt.QuoteDouble && len(items[i]) > 1 && strings.Contains(items[i][:len(items[i])-1], " ") {
			items[i] = strings.Replace(items[i], " ", `\ `, -1)
		}

		if items[i][len(items[i])-1] == '/' || items[i][len(items[i])-1] == '=' {
			suggest[i] = []rune(items[i])
		} else {
			suggest[i] = []rune(items[i] + " ")
		}
	}*/

	for i := range items {
		if len(items[i]) == 0 {
			items[i] = " "
		}
		if items[i][len(items[i])-1] != ' ' && items[i][len(items[i])-1] != '=' && items[i][len(items[i])-1] != ' ' {
			items[i] += " "
		}
	}

	return
}
