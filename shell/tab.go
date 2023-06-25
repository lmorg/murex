package shell

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/shell/hintsummary"
	"github.com/lmorg/murex/shell/history"
	"github.com/lmorg/murex/shell/preview"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/ansi/codes"
	"github.com/lmorg/murex/utils/dedup"
	"github.com/lmorg/murex/utils/objectkeys"
	"github.com/lmorg/murex/utils/parser"
	"github.com/lmorg/murex/utils/readline"
)

func errCallback(err error) {
	s := err.Error()
	if ansi.IsAllowed() {
		s = codes.Reset + codes.FgRed + s
	}
	Prompt.ForceHintTextUpdate(s)
}

var (
	rxHistTag       = regexp.MustCompile(`^[-_a-zA-Z0-9]+$`)
	rxValidVariable = regexp.MustCompile(`^[._a-zA-Z0-9]+$`)
)

func tabCompletion(line []rune, pos int, dtc readline.DelayedTabContext) *readline.TabCompleterReturnT {
	r := new(readline.TabCompleterReturnT)

	if pos < 0 {
		return r
	}

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

	rows, _ := lang.ShellProcess.Config.Get("shell", "max-suggestions", types.Integer)
	Prompt.MaxTabCompleterRows = rows.(int)

	switch {
	case pt.VarSigil != "":
		//panic(json.LazyLoggingPretty(pt))
		if pt.VarLoc < len(line) {
			r.Prefix = strings.TrimSpace(string(line[pt.VarLoc:]))
		}
		r.Prefix = pt.VarSigil + r.Prefix
		if strings.Contains(r.Prefix, ".") {
			name := strings.Split(r.Prefix, ".")[0][1:]
			v, err := lang.ShellProcess.Variables.GetValue(name)
			if err != nil {
				act.ErrCallback(err)
			}

			nameLen := len(name) + 1
			prefixLen := len(r.Prefix) - len(name) - 1
			writeString := func(s string) error {
				if len(s) < prefixLen ||
					s[:prefixLen] != r.Prefix[nameLen:] ||
					!rxValidVariable.MatchString(s) {

					return nil
				}

				act.Items = append(act.Items, s[prefixLen:])
				return nil
			}

			err = objectkeys.Recursive(dtc.Context, "", v, ".", writeString, -1)
			if err != nil {
				act.ErrCallback(err)
			}

		} else {
			act.Items = autocomplete.MatchVars(r.Prefix)
		}

	case pt.Comment && pt.FuncName == "^":
		for h := Prompt.History.Len() - 1; h > -1; h-- {
			line, _ := Prompt.History.GetLine(h)
			linePt, _ := parser.Parse([]rune(line), 0)
			if linePt.Comment && rxHistTag.MatchString(linePt.CommentMsg) && strings.HasPrefix(linePt.CommentMsg, pt.CommentMsg) {
				suggestion := linePt.CommentMsg[len(pt.CommentMsg):]
				act.Items = append(act.Items, suggestion)
				act.Definitions[suggestion] = line
				r.Prefix = pt.CommentMsg
			}
		}

	case pt.FuncName == "^":
		autocompleteHistoryHat(&act)

	case pt.ExpectFunc:
		if len(line) == 0 {
			Prompt.ForceHintTextUpdate("Tip: press [ctrl]+[r] to recall previously used command lines")
		}
		go autocomplete.UpdateGlobalExeList()

		if pt.Loc < len(line) {
			r.Prefix = strings.TrimSpace(string(line[pt.Loc:]))
		}

		r.HintCache = cacheHints
		r.Preview = preview.Command

		switch pt.PipeToken {
		case parser.PipeTokenNone:
			autocompleteFunctions(&act, r.Prefix)
			v, _ := lang.ShellProcess.Config.Get("shell", "auto-cd", types.Boolean)
			autoCd, _ := v.(bool)
			if autoCd {
				autocomplete.MatchDirectories(r.Prefix, &act)
			}

		case parser.PipeTokenPosix:
			autocomplete.MatchFunction(r.Prefix, &act)

		case parser.PipeTokenArrow:
			act.TabDisplayType = readline.TabDisplayList
			globalExes := autocomplete.GlobalExes.Get()

			if lang.MethodStdout.Exists(pt.LastFuncName, types.Any) {
				// match everything
				dump := lang.MethodStdout.Dump()

				if len(r.Prefix) == 0 {
					for dt := range dump {
						act.Items = append(act.Items, dump[dt]...)
						for i := range dump[dt] {
							act.Definitions[dump[dt][i]] = string(hintsummary.Get(dump[dt][i], (*globalExes)[dump[dt][i]]))
						}
					}

				} else {

					for dt := range dump {
						for i := range dump[dt] {
							if strings.HasPrefix(dump[dt][i], r.Prefix) {
								act.Items = append(act.Items, dump[dt][i][len(r.Prefix):])
								act.Definitions[dump[dt][i][len(r.Prefix):]] = string(hintsummary.Get(dump[dt][i], (*globalExes)[dump[dt][i]]))
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
					if len(r.Prefix) == 0 {
						act.Items = append(act.Items, inTypes...)
						for j := range inTypes {
							act.Definitions[inTypes[j]] = string(hintsummary.Get(inTypes[j], (*globalExes)[inTypes[j]]))
						}
						continue
					}
					for j := range inTypes {
						if strings.HasPrefix(inTypes[j], r.Prefix) {
							act.Items = append(act.Items, inTypes[j][len(r.Prefix):])
							act.Definitions[inTypes[j][len(r.Prefix):]] = string(hintsummary.Get(inTypes[j], (*globalExes)[inTypes[j]]))
						}
					}
				}
			}

			// If `->` returns no results then fall back to returning everything
			if len(act.Items) == 0 {
				autocompleteFunctions(&act, r.Prefix)
			}

		default:
			autocompleteFunctions(&act, r.Prefix)
		}

	default:
		autocomplete.InitExeFlags(pt.FuncName)
		if !pt.ExpectParam && len(act.ParsedTokens.Parameters) > 0 {
			r.Prefix = pt.Parameters[len(pt.Parameters)-1]
		}

		autocomplete.MatchFlags(&act)
		r.Preview = preview.Parameter
	}

	Prompt.MinTabItemLength = act.MinTabItemLength
	width := readline.GetTermWidth()
	switch {
	case width >= 200:
		Prompt.MaxTabItemLength = width / 5
	case width >= 150:
		Prompt.MaxTabItemLength = width / 4
	case width >= 100:
		Prompt.MaxTabItemLength = width / 3
	case width >= 70:
		Prompt.MaxTabItemLength = width / 2
	}

	var i int
	if act.DoNotSort {
		i = len(act.Items)
	} else {
		i = dedup.SortAndDedupString(act.Items)
	}
	if !act.DoNotEscape {
		autocomplete.FormatSuggestions(&act)
	}

	//return prefix, act.Items[:i], act.Definitions, act.TabDisplayType
	r.Suggestions = act.Items[:i]
	r.Descriptions = act.Definitions
	r.DisplayType = act.TabDisplayType
	return r
}

func autocompleteFunctions(act *autocomplete.AutoCompleteT, prefix string) {
	act.TabDisplayType = readline.TabDisplayGrid

	autocomplete.MatchFunction(prefix, act)
}

func autocompleteHistoryLine(prefix string) ([]string, map[string]string) {
	var (
		items       []string
		definitions = make(map[string]string)
	)

	dump := Prompt.History.Dump().([]history.Item)

	for i := len(dump) - 1; i > -1; i-- {
		if len(dump[i].Block) <= len(prefix) {
			continue
		}

		if !strings.HasPrefix(dump[i].Block, prefix) {
			continue
		}

		item := dump[i].Block[len(prefix):]

		if definitions[item] != "" {
			continue
		}

		dateTime := dump[i].DateTime.Format("02-Jan-06 15:04")
		items = append(items, item)
		definitions[item] = dateTime
	}

	return items, definitions
}

func autocompleteHistoryHat(act *autocomplete.AutoCompleteT) {
	size := Prompt.History.Len()
	act.Items = make([]string, size)
	act.Definitions = make(map[string]string, size)
	dump := Prompt.History.Dump().([]history.Item)

	j := len(dump)
	for i := range dump {
		j--
		s := strconv.Itoa(dump[i].Index)
		act.Definitions[s] = dump[i].Block
		act.Items[j] = s
	}

	act.TabDisplayType = readline.TabDisplayList
	act.DoNotSort = true
}

func cacheHints(prefix string, exes []string) []string {
	v, err := lang.ShellProcess.Config.Get("shell", "pre-cache-hint-summaries", types.String)
	if err != nil {
		v = ""
	}
	if v.(string) != "on-tab" {
		return nil
	}

	hints := make([]string, len(exes))
	external := autocomplete.GlobalExes.Get()
	for i := range exes {
		exe := strings.TrimSpace(prefix + exes[i])
		hints[i] = string(hintsummary.Get(exe, (*external)[exe]))
	}

	return hints
}
