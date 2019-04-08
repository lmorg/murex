package shell

import (
	"fmt"
	"os"

	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/history"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/counter"
	"github.com/lmorg/murex/utils/home"
	"github.com/lmorg/readline"
)

var (
	// Interactive describes whether murex is running as an interactive shell
	// or not
	Interactive bool

	// Prompt is the readline instance
	Prompt = readline.NewInstance()

	// PromptId is an custom defined ID for each prompt Goprocess so we don't
	// accidentally end up with multiple prompts running
	PromptId = new(counter.MutexCounter)
)

// Start the interactive shell
func Start() {
	/*defer func() {
		if debug.Enabled {
			return
		}
		if r := recover(); r != nil {
			os.Stderr.WriteString(fmt.Sprintln("Panic caught:", r))
			Start()
		}
	}()*/

	var err error

	Interactive = true
	Prompt.TempDirectory = consts.TempDir
	Prompt.TabCompleter = tabCompletion
	Prompt.SyntaxCompleter = syntaxCompletion
	Prompt.HistoryAutoWrite = false

	h, err := history.New(home.MyDir + consts.PathSlash + ".murex_history")
	if err != nil {
		lang.ShellProcess.Stderr.Writeln([]byte("Error opening history file: " + err.Error()))
	} else {
		Prompt.History = h
	}

	SignalHandler(true)

	v, err := lang.ShellProcess.Config.Get("shell", "max-suggestions", types.Integer)
	if err != nil {
		v = 4
	}
	Prompt.MaxTabCompleterRows = v.(int)

	ShowPrompt()

	noQuit := make(chan int)
	<-noQuit
}

// ShowPrompt display's the shell command line prompt
func ShowPrompt() {
	if !Interactive {
		panic("shell.ShowPrompt() called before initialising prompt with shell.Start()")
	}

	thisProc := PromptId.Add()

	nLines := 1
	var merged string
	var block []rune
	Prompt.GetMultiLine = func(r []rune) []rune {
		var multiLine []rune
		if len(block) == 0 {
			multiLine = r
		} else {
			multiLine = append(append(block, []rune(utils.NewLineString)...), r...)
		}

		expanded, err := history.ExpandVariables(multiLine, Prompt)
		if err != nil {
			expanded = multiLine
		}
		return expanded
	}

	for {
		getSyntaxHighlighting()
		getHintTextEnabled()
		getHintTextFormatting()
		cachedHintText = []rune{}

		if nLines > 1 {
			getMultilinePrompt(nLines)
		} else {
			block = []rune{}
			getPrompt()
		}

		line, err := Prompt.Readline()
		if err != nil {
			switch err {
			case readline.CtrlC:
				merged = ""
				nLines = 1
				fmt.Println(PromptSIGINT)
				continue

			case readline.EOF:
				fmt.Println(utils.NewLineString)
				//return
				os.Exit(0)

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
			//ansi.Stderrln(lang.ShellProcess, ansi.FgRed, err.Error())
			lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
			merged = ""
			nLines = 1
			continue
		}

		if string(expanded) != string(block) {
			os.Stdout.WriteString(ansi.ExpandConsts("{GREEN}") + string(expanded) + ansi.ExpandConsts("{RESET}") + utils.NewLineString)
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
			mergedExp, err := history.ExpandVariablesInLine([]rune(merged), Prompt)
			if err == nil {
				merged = string(mergedExp)
			}

			Prompt.History.Write(merged)

			nLines = 1
			merged = ""

			//lang.ShellExitNum, _ = lang.RunBlockShellConfigSpaceWithPrompt(expanded, nil, new(term.Out), term.NewErr(ansi.IsAllowed()), thisProc)
			fork := lang.ShellProcess.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN)
			fork.Stderr = term.NewErr(ansi.IsAllowed())
			fork.PromptId = thisProc
			lang.ShellExitNum, _ = fork.Execute(expanded)

			term.CrLf.Write()

			if PromptId.NotEqual(thisProc) {
				return
			}
		}
	}
}

func getSyntaxHighlighting() {
	highlight, err := lang.ShellProcess.Config.Get("shell", "syntax-highlighting", types.Boolean)
	if err != nil {
		highlight = false
	}
	if highlight.(bool) {
		Prompt.SyntaxHighlighter = syntaxHighlight
	} else {
		Prompt.SyntaxHighlighter = nil
	}
}

func getHintTextEnabled() {
	showHintText, err := lang.ShellProcess.Config.Get("shell", "hint-text-enabled", types.Boolean)
	if err != nil {
		showHintText = false
	}
	if showHintText.(bool) {
		Prompt.HintText = hintText
	} else {
		Prompt.HintText = nil
	}
}

func getHintTextFormatting() {
	formatting, err := lang.ShellProcess.Config.Get("shell", "hint-text-formatting", types.String)
	if err != nil {
		formatting = ""
	}
	Prompt.HintFormatting = ansi.ExpandConsts(formatting.(string))
}
