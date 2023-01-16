package expressions

import (
	"fmt"
	"path/filepath"
	"strings"
	"sync"

	"github.com/lmorg/murex/utils/ansi/codes"
	"github.com/lmorg/murex/utils/escape"
	"github.com/lmorg/murex/utils/readline"
)

func (tree *ParserT) parseGlob(glob []rune) ([]string, error) {
	globS := string(glob)

	match, globErr := filepath.Glob(globS)
	expand, err := expandGlobPrompt(tree.statement.String(), globS, match)
	if err != nil {
		return nil, err
	}
	if !expand {
		return nil, nil
	}

	if globErr != nil {
		return nil, fmt.Errorf("invalid glob: '%s'\n%s", globS, err.Error())
	}
	if len(match) == 0 {
		return nil, fmt.Errorf("glob returned zero results.\nglob: '%s'", globS)
	}

	return match, nil
}

var rlMutex sync.Mutex

func expandGlobPrompt(cmd string, before string, match []string) (bool, error) {
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
