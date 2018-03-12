package shell

import (
	"fmt"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/shell/history"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/home"
	"github.com/lmorg/murex/utils/readline"
)

var Interactive bool

// Start the interactive shell
func Start() {
	var err error

	/*Instance, err = readline.NewEx(&readline.Config{
		InterruptPrompt:        interruptPrompt,
		AutoComplete:           murexCompleter,
		FuncFilterInputRune:    filterInput,
		DisableAutoSaveHistory: true,
		NoEofOnEmptyDelete:     true,
	})

	if err != nil {
		panic(err)
	}*/

	Interactive = true
	readline.TabCompleter = tabCompleter
	readline.HistoryAutoWrite = false

	h, err := history.New(home.MyDir + consts.PathSlash + ".murex_history")
	if err != nil {
		ansi.Stderrln(ansi.FgRed, "Error opening history file: "+err.Error())
	}

	readline.History = h

	/*Instance.Config.SetListener(listener)
	defer Instance.Close()
	SigHandler()*/
	go autocomplete.UpdateGlobalExeList()

	prompt()
}

func prompt() {
	nLines := 1
	var merged string
	var block []rune

	for {
		getSyntaxHighlighting()

		if nLines > 1 {
			getMultilinePrompt(nLines)
		} else {
			getPrompt()
		}

		/*line, err := Instance.Readline()
		if err == readline.ErrInterrupt {
			merged = ""
			nLines = 1
			continue

		} else if err == io.EOF {
			break
		}*/

		line, err := readline.Readline()
		if err != nil {
			switch err.Error() {
			case readline.ErrCtrlC:
				merged = ""
				nLines = 1
				fmt.Println("^C")
				continue
			case readline.ErrEOF:
				return
			default:
				panic(err)
			}
		}

		if nLines > 1 {
			block = append(block, []rune(utils.NewLineString+line)...)
		} else {
			block = []rune(line)
		}

		expanded, err := history.ExpandVariables(block)
		if err != nil {
			ansi.Stderrln(ansi.FgRed, err.Error())
			merged = ""
			nLines = 1
			continue
		}

		if string(expanded) != string(block) {
			ansi.Stderrln(ansi.FgGreen, string(expanded))
		}

		pt, _ := parse(block)
		switch {
		case pt.Bracket > 0:
			nLines++
			merged += line + `^\n`

		case pt.Escaped:
			nLines++
			merged += line[:len(line)-1] + `^\n`

		case pt.QuoteSingle:
			nLines++
			merged += line + `^\n`

		case pt.QuoteDouble:
			nLines++
			merged += line + `\n`

		case len(block) == 0:
			continue

		default:
			merged += line
			mergedExp, err := history.ExpandVariables([]rune(merged))
			if err == nil {
				merged = string(mergedExp)
			}
			/*Instance.SaveHistory(merged)
			if History.Last != merged {
				History.Last = merged
				History.Write(merged)
			}*/
			readline.History.Write(merged)

			nLines = 1
			merged = ""

			lang.ShellExitNum, _ = lang.RunBlockShellNamespace(expanded, nil, nil, nil)
			//streams.CrLf.Write()
		}
	}
}

/*func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, true
	case readline.CharForward:
		forward++
		return r, true
	}
	forward = 0
	return r, true
}*/

func getSyntaxHighlighting() {
	highlight, err := proc.ShellProcess.Config.Get("shell", "syntax-highlighting", types.Boolean)
	if err != nil {
		highlight = false
	}
	if highlight.(bool) == true {
		readline.SyntaxHighlight = syntaxHighlight
	} else {
		readline.SyntaxHighlight = nil
	}
}

/*func syntaxHighlight(input string) (output string) {
	_, output = parse([]rune(input))
	return
}*/
