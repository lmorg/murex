package lang

/*
import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi/codes"
	"github.com/lmorg/murex/utils/escape"
	"github.com/lmorg/murex/utils/home"
	"github.com/lmorg/murex/utils/lists"
	"github.com/lmorg/murex/utils/readline"
)

var (
	rxTokenIndex   = regexp.MustCompile(`(.*?)\[(.*?)\]`)
	rxTokenElement = regexp.MustCompile(`(.*?)\[\[(.*?)\]\]`)
	rxTokenRange   = regexp.MustCompile(`(.*?)\[(.*?)\]([bt8ernsi]*)`)
	rlMutex        sync.Mutex
)

const (
	errEmptyArray = "array '@%s' is empty"
	errEmptyRange = "range '@%s' is empty"
)

// PPConfigDefaults returns the following config settings:
//   - strictArrays
//   - expandGlob
func PPConfigDefaults(p *Process) (bool, bool) {
	strictArrays, err := p.Config.Get("proc", "strict-arrays", "bool")
	if err != nil {
		strictArrays = true
	}

	expandGlob, err := p.Config.Get("shell", "expand-glob", "bool")
	if err != nil {
		expandGlob = true
	}
	expandGlob = expandGlob.(bool) && p.Scope.Id == 0 && !lists.Match(GetNoGlobCmds(), p.Name.String())

	return strictArrays.(bool), expandGlob.(bool)
}

// ParseParameters is an internal function to parse parameters
func ParseParameters(p *Process, paramT *parameters.Parameters) error {
	var namedPipeIsParam bool
	params := []string{}

	strictArrays, expandGlob := PPConfigDefaults(p)

	for i := range paramT.Tokens {
		params = append(params, "")

		var tCount bool
		for j := range paramT.Tokens[i] {
			switch paramT.Tokens[i][j].Type {


			case parameters.TokenTypeGlob:
				if !expandGlob || p.Parent.Id != ShellProcess.Id || p.Background.Get() || !Interactive {
					params[len(params)-1] += paramT.Tokens[i][j].Key

				} else {
					match, globErr := filepath.Glob(paramT.Tokens[i][j].Key)
					glob, err := autoGlobPrompt(p.Name.String(), paramT.Tokens[i][j].Key, match)
					if err != nil {
						return err
					}
					if glob {
						if globErr != nil {
							return fmt.Errorf("invalid glob: '%s'\n%s", paramT.Tokens[i][j].Key, err.Error())
						}
						if len(match) == 0 {
							return fmt.Errorf("glob returned zero results.\nglob: '%s'", paramT.Tokens[i][j].Key)
						}
						if !tCount {
							params = params[:len(params)-1]
						}
						params = append(params, match...)
					} else {
						params[len(params)-1] += paramT.Tokens[i][j].Key
					}
				}
				tCount = true
				namedPipeIsParam = true





}

func autoGlobPrompt(cmd string, before string, match []string) (bool, error) {
	rlMutex.Lock()
	defer rlMutex.Unlock() // performance doesn't matter here

	rl := readline.NewInstance()
	prompt := fmt.Sprintf("(%s) Do you wish to expand '%s'? [yN]: ", cmd, before)
	rl.SetPrompt(prompt)
	rl.HintText = func(_ []rune, _ int) []rune { return autoGlobPromptHintText(rl, match) }
	rl.History = new(readline.NullHistory)

	for {
		line, err := rl.Readline()
		if err != nil {
			return false, err
		}

		switch strings.ToLower(line) {
		case "y", "yes":
			return true, nil
		case "", "n", "no":
			return false, nil
		}
	}
}

const (
	warningNoGlobMatch = "Warning: no files match that pattern"
	globExpandsTo      = "Glob expands to: "
)

func autoGlobPromptHintText(rl *readline.Instance, match []string) []rune {
	if len(match) == 0 {
		rl.HintFormatting = codes.FgRed
		return []rune(warningNoGlobMatch)
	}

	slice := make([]string, len(match))
	copy(slice, match)
	escape.CommandLine(slice)
	after := globExpandsTo + strings.Join(slice, ", ")
	return []rune(after)
}
*/
