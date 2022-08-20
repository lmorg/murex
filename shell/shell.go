package shell

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/history"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/cd/cache"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/counter"
	"github.com/lmorg/murex/utils/readline"
	"github.com/lmorg/murex/utils/spellcheck"
)

var (
	// Prompt is the readline instance
	Prompt = readline.NewInstance()

	// PromptId is an custom defined ID for each prompt Goprocess so we don't
	// accidentally end up with multiple prompts running
	PromptId = new(counter.MutexCounter)

	rxHashTag = regexp.MustCompile(`#[-_a-zA-Z0-9]+$`)
)

// Start the interactive shell
func Start() {
	cache.GatherFileCompletions(".")

	if debug.Enabled {
		defer func() {
			if r := recover(); r != nil {
				os.Stderr.WriteString(fmt.Sprintln("Panic caught:", r))
				Start()
			}
		}()
	}

	var err error

	lang.Interactive = true
	Prompt.TempDirectory = consts.TempDir
	Prompt.TabCompleter = tabCompletion
	Prompt.SyntaxCompleter = syntaxCompletion
	Prompt.HistoryAutoWrite = false

	setPromptHistory()

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
	if !lang.Interactive {
		panic("shell.ShowPrompt() called before initialising prompt with shell.Start()")
	}

	var (
		thisProc = PromptId.Add()
		nLines   = 1
		merged   string
		block    []rune
	)

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

	Prompt.DelayedSyntaxWorker = Spellchecker

	for {
		//debug.Log("ShowPrompt (for{})")

		getSyntaxHighlighting()
		getHintTextEnabled()
		getHintTextFormatting()
		cachedHintText = []rune{}

		if nLines > 1 {
			getMultilinePrompt(nLines)
		} else {
			block = []rune{}
			getPrompt()
			writeTitlebar()
		}

		//debug.Log("ShowPrompt (Prompt.Readline())")
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
				lang.Exit(0)

			default:
				panic(err)
			}
		}

		switch {
		case nLines > 1:
			block = append(block, []rune(utils.NewLineString+line)...)
		default:
			block = []rune(line)
		}

		expanded, err := history.ExpandVariables(block, Prompt)
		if err != nil {
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
			merged += strings.TrimSpace(line[:len(line)-1]) + " "

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

			macroFind, macroReplace, err := getMacroVars(merged)
			if err != nil {
				nLines = 1
				merged = ""
				continue
			}

			if len(macroFind) > 0 {
				if !rxHashTag.MatchString(merged) {
					merged = expandMacroVars(merged, macroFind, macroReplace)
				}
				expanded = []rune(expandMacroVars(string(expanded), macroFind, macroReplace))
			}

			Prompt.History.Write(merged)

			nLines = 1
			merged = ""

			fork := lang.ShellProcess.Fork(lang.F_PARENT_VARTABLE | lang.F_NEW_MODULE | lang.F_NO_STDIN)
			fork.FileRef = &ref.File{Source: &ref.Source{Module: app.ShellModule}}
			fork.Stderr = term.NewErr(ansi.IsAllowed())
			fork.PromptId = thisProc
			fork.CCEvent = lang.ShellProcess.CCEvent
			fork.CCExists = lang.ShellProcess.CCExists
			lang.ShellExitNum, _ = fork.Execute(expanded)

			if PromptId.NotEqual(thisProc) {
				return
			}
		}
	}
}

var rxMacroVar = regexp.MustCompile(`(\^\$[-_a-zA-Z0-9]+)`)

func getMacroVars(s string) ([]string, []string, error) {
	var err error

	if !rxMacroVar.MatchString(s) {
		return nil, nil, nil
	}

	match := rxMacroVar.FindAllString(s, -1)
	vars := make([]string, len(match))
	for i := range match {
		for {
			rl := readline.NewInstance()
			rl.SetPrompt(ansi.ExpandConsts(fmt.Sprintf(
				"{YELLOW}Enter value for: {RED}%s{YELLOW}? {RESET}", match[i][2:],
			)))
			rl.History = new(readline.NullHistory)
			vars[i], err = rl.Readline()
			if err != nil {
				return nil, nil, err
			}
			if vars[i] != "" {
				break
			}
			os.Stderr.WriteString(ansi.ExpandConsts("{RED}Cannot use zero length strings. Please enter a value or press CTRL+C to cancel.{RESET}\n"))
		}
	}

	return match, vars, nil
}

func expandMacroVars(s string, match, vars []string) string {
	for i := range match {
		s = strings.ReplaceAll(s, match[i], vars[i])
	}

	return s
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

var ignoreSpellCheckErr bool

func Spellchecker(r []rune) []rune {
	s := string(r)
	new, err := spellcheck.String(s)
	if err != nil && !ignoreSpellCheckErr {
		ignoreSpellCheckErr = true
		hint := fmt.Sprintf("{RED}Spellchecker error: %s{RESET} {BLUE}https://murex.rocks/docs/user-guide/spellcheck.html{RESET}", err.Error())
		Prompt.ForceHintTextUpdate(ansi.ExpandConsts(hint))
	}

	ignoreSpellCheckErr = false // reset ignore status

	return []rune(new)
}
