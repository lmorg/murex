package onkeypress

import (
	"fmt"

	"github.com/lmorg/murex/utils/readline"
)

const (
	metaHotKeyActons = "Actions"
	metaSetLine      = "SetLine"
	metaSetPos       = "SetPos"
	metaClose        = "CloseReadline"
	metaHintText     = "SetHintText"
	metaContinue     = "Continue"
)

func createMeta(line []rune, pos int) map[string]any {
	return map[string]any{
		metaHotKeyActons: []string{},
		metaSetLine:      string(line),
		metaSetPos:       pos,
		metaClose:        false,
		metaHintText:     "",
		metaContinue:     false,
	}
}

type metaT struct {
	Actions       []string
	SetLine       string
	SetPos        int
	CloseReadline bool
	HintText      string
	Continue      bool
}

func compileActionSlice(actions []string) ([]func(*readline.Instance), error) {
	functions := make([]func(*readline.Instance), len(actions))

	for i := 0; i < len(actions); i++ {
		fn, ok := fnLookup[actions[i]]
		if !ok {
			return nil, fmt.Errorf("invalid action name: '%s'", actions[i])
		}

		functions[i] = fn
	}

	return functions, nil
}
