package shell

import (
	"fmt"
	"os"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/shell/history"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/home"
	"github.com/lmorg/murex/utils/readline"
)

var (
	// Interactive describes whether murex is running as an interactive shell or not
	Interactive bool

	// Prompt is the readline instance
	Prompt *readline.Instance = readline.NewInstance()
)

// Start the interactive shell
func Start() {
	var err error

	Interactive = true
	Prompt.TabCompleter = tabCompletion
	Prompt.SyntaxCompleter = syntaxCompletion
	Prompt.HistoryAutoWrite = false
	Prompt.HintText = hintText

	h, err := history.New(home.MyDir + consts.PathSlash + ".murex_history")
	if err != nil {
		ansi.Stderrln(proc.ShellProcess, ansi.FgRed, "Error opening history file: "+err.Error())
	} else {
		Prompt.History = h
	}

	/*Instance.Config.SetListener(listener)
	defer Instance.Close()*/
	SigHandler()

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

		line, err := Prompt.Readline()
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

		expanded, err := history.ExpandVariables(block, Prompt)
		if err != nil {
			ansi.Stderrln(proc.ShellProcess, ansi.FgRed, err.Error())
			merged = ""
			nLines = 1
			continue
		}

		if string(expanded) != string(block) {
			//ansi.Stderrln(proc.ShellProcess, ansi.FgGreen, string(expanded))
			os.Stderr.WriteString(ansi.FgGreen + string(expanded) + ansi.Reset + utils.NewLineString)
		}

		pt, _ := parse(block)
		switch {
		case pt.NestedBlock > 0:
			nLines++
			merged += line + `^\n`

		case pt.Escaped:
			nLines++
			merged += line[:len(line)-1] + `^\n`

		case pt.QuoteSingle, pt.QuoteBrace > 0:
			nLines++
			merged += line + `^\n`

		case pt.QuoteDouble:
			nLines++
			merged += line + `\n`

		case len(block) == 0:
			continue

		default:
			merged += line
			mergedExp, err := history.ExpandVariables([]rune(merged), Prompt)
			if err == nil {
				merged = string(mergedExp)
			}

			Prompt.History.Write(merged)

			nLines = 1
			merged = ""

			lang.ShellExitNum, _ = lang.RunBlockShellNamespace(expanded, nil, nil, nil)
			streams.CrLf.Write()
		}
	}
}

func getSyntaxHighlighting() {
	highlight, err := proc.ShellProcess.Config.Get("shell", "syntax-highlighting", types.Boolean)
	if err != nil {
		highlight = false
	}
	if highlight.(bool) == true {
		Prompt.SyntaxHighlighter = syntaxHighlight
	} else {
		Prompt.SyntaxHighlighter = nil
	}
}
