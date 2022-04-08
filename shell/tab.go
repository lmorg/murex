package shell

import (
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/shell/hintsummary"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/dedup"
	"github.com/lmorg/murex/utils/parser"
	"github.com/lmorg/murex/utils/readline"
)

func errCallback(err error) {
	s := err.Error()
	if ansi.IsAllowed() {
		s = ansi.Reset + ansi.FgRed + s
	}
	Prompt.SetHintText(s)
}

var rxHistTag = regexp.MustCompile(`^[-_a-zA-Z0-9]+$`)

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

	rows, err := lang.ShellProcess.Config.Get("shell", "max-suggestions", types.Integer)
	if err != nil {
		rows = 8
	}
	Prompt.MaxTabCompleterRows = rows.(int)

	switch {
	case pt.Variable != "":
		if pt.VarLoc < len(line) {
			prefix = strings.TrimSpace(string(line[pt.VarLoc:]))
		}
		prefix = pt.Variable + prefix
		act.Items = autocomplete.MatchVars(prefix)

	case pt.Comment && pt.FuncName == "^":
		for h := Prompt.History.Len() - 1; h > -1; h-- {
			line, _ := Prompt.History.GetLine(h)
			linePt, _ := parser.Parse([]rune(line), 0)
			if linePt.Comment && rxHistTag.MatchString(linePt.CommentMsg) && strings.HasPrefix(linePt.CommentMsg, pt.CommentMsg) {
				suggestion := linePt.CommentMsg[len(pt.CommentMsg):]
				act.Items = append(act.Items, suggestion)
				act.Definitions[suggestion] = line
				prefix = pt.CommentMsg
			}
		}

	case pt.ExpectFunc:
		if pt.Loc < len(line) {
			prefix = strings.TrimSpace(string(line[pt.Loc:]))
		}

		switch pt.PipeToken {
		case parser.PipeTokenPosix:
			act.Items = autocomplete.MatchFunction(prefix, &act)
		case parser.PipeTokenArrow:
			act.TabDisplayType = readline.TabDisplayList
			if lang.MethodStdout.Exists(pt.LastFuncName, types.Any) {
				// match everything
				dump := lang.MethodStdout.Dump()

				if len(prefix) == 0 {
					for dt := range dump {
						act.Items = append(act.Items, dump[dt]...)
						for i := range dump[dt] {
							act.Definitions[dump[dt][i]] = string(hintsummary.Get(dump[dt][i], autocomplete.GlobalExes[dump[dt][i]]))
						}
					}

				} else {

					for dt := range dump {
						for i := range dump[dt] {
							if strings.HasPrefix(dump[dt][i], prefix) {
								act.Items = append(act.Items, dump[dt][i][len(prefix):])
								act.Definitions[dump[dt][i][len(prefix):]] = string(hintsummary.Get(dump[dt][i], autocomplete.GlobalExes[dump[dt][i]]))
							}
						}
					}
				}

			} else {
				// match type
				outTypes := lang.MethodStdout.Types(pt.LastFuncName)
				outTypes = append(outTypes, types.Any)
				for i := range outTypes {
					inTypes := lang.MethodStdin.Get(outTypes[i])
					if len(prefix) == 0 {
						act.Items = append(act.Items, inTypes...)
						for j := range inTypes {
							act.Definitions[inTypes[j]] = string(hintsummary.Get(inTypes[j], autocomplete.GlobalExes[inTypes[j]]))
						}
						continue
					}
					for j := range inTypes {
						if strings.HasPrefix(inTypes[j], prefix) {
							act.Items = append(act.Items, inTypes[j][len(prefix):])
							act.Definitions[inTypes[j][len(prefix):]] = string(hintsummary.Get(inTypes[j], autocomplete.GlobalExes[inTypes[j]]))
						}
					}
				}
			}

			// If `->` returns no results then fall back to returning everything
			if len(act.Items) == 0 {
				autocompleteFunctions(&act, prefix)
			}

		default:
			autocompleteFunctions(&act, prefix)
		}

	default:
		autocomplete.InitExeFlags(pt.FuncName)
		if !pt.ExpectParam && len(act.ParsedTokens.Parameters) > 0 {
			prefix = pt.Parameters[len(pt.Parameters)-1]
		}

		autocomplete.MatchFlags(&act)
	}

	Prompt.MinTabItemLength = act.MinTabItemLength

	var i int
	if act.DoNotSort {
		i = len(act.Items)
	} else {
		i = dedup.SortAndDedupString(act.Items)
	}
	autocomplete.FormatSuggestions(&act)

	return prefix, act.Items[:i], act.Definitions, act.TabDisplayType
}

func autocompleteFunctions(act *autocomplete.AutoCompleteT, prefix string) {
	act.TabDisplayType = readline.TabDisplayGrid

	act.Items = autocomplete.MatchFunction(prefix, act)

	/*sort.Strings(act.Items)
	for i := 0; i < Prompt.MaxTabCompleterRows && i <= len(act.Items)-1; i++ {
		cmd := prefix + act.Items[i]
		if len(cmd) > 1 && cmd[len(cmd)-1] == ':' {
			cmd = cmd[:len(cmd)-1]
		}
		act.Definitions[act.Items[i]] = string(hintsummary.Get(cmd, autocomplete.GlobalExes[cmd]))
	}*/
}
