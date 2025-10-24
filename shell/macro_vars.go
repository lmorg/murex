package shell

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/cache"
	"github.com/lmorg/murex/utils/lists"
	"github.com/lmorg/readline/v4"
)

var rxMacroVar = regexp.MustCompile(`(\^\$[-_a-zA-Z0-9]+)`)

func getMacroVars(cmdline string) ([]string, []string, error) {
	var err error

	if !rxMacroVar.MatchString(cmdline) {
		return nil, nil, nil
	}

	history := make(map[string]macroVarHistory)
	cache.Read(cache.MACRO_VAR_HISTORY, cmdline, &history)

	assigned := make(map[string]bool)
	match := rxMacroVar.FindAllString(cmdline, -1)
	vars := make([]string, len(match))

	for i := range match {
		if assigned[match[i]] {
			continue
		}

		for {
			rl := readline.NewInstance()
			rl.SetPrompt(ansi.ExpandConsts(fmt.Sprintf(
				"{YELLOW}Enter value for: {RED}%s{YELLOW}? {RESET}", match[i][2:],
			)))
			rl.History = history[match[i]]
			rl.TabCompleter = history[match[i]].TabCompleter
			vars[i], err = rl.Readline()
			if err != nil {
				return nil, nil, err
			}
			if vars[i] != "" {
				break
			}
			os.Stderr.WriteString(ansi.ExpandConsts("{RED}Cannot use zero length strings. Please enter a value or press CTRL+C to cancel.{RESET}\n"))
		}
		assigned[match[i]] = true
		if matched := lists.MatchIndexString(history[match[i]], vars[i]); matched != -1 {
			new, err := lists.RemoveOrdered(history[match[i]], matched)
			if err != nil {
				history[match[i]] = new
			}
		}
		history[match[i]] = append(history[match[i]], vars[i])
	}

	cache.Write(cache.MACRO_VAR_HISTORY, cmdline, history, time.Now().AddDate(0, 6, 0))
	return match, vars, nil
}

func expandMacroVars(s string, match, vars []string) string {
	for i := range match {
		s = strings.ReplaceAll(s, match[i], vars[i])
	}

	return s
}

type macroVarHistory []string

func (h macroVarHistory) Write(s string) (int, error) { return len(h), nil }
func (h macroVarHistory) Len() int                    { return len(h) }
func (h macroVarHistory) Dump() any                   { return h }

func (h macroVarHistory) GetLine(i int) (string, error) {
	switch {
	case i < 0:
		return "", errors.New("requested history item out of bounds: < 0")
	case i > h.Len()-1:
		return "", errors.New("requested history item out of bounds: > Len()")
	default:
		return (h)[i], nil
	}
}

func (h macroVarHistory) TabCompleter(r []rune, _ int, _ readline.DelayedTabContext) *readline.TabCompleterReturnT {
	tcr := &readline.TabCompleterReturnT{
		Prefix: string(r),
	}

	for i := range h {
		if strings.HasPrefix(h[i], tcr.Prefix) {
			tcr.Suggestions = append(tcr.Suggestions, h[i][len(r):])
		}
	}

	return tcr
}
