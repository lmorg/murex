//go:generate ./actions.mx

package onkeypress

import (
	"fmt"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/readline"
)

const (
	metaHotKeyActions = "Actions"
	metaSetLine       = "SetLine"
	metaSetPos        = "SetPos"
	metaClose         = "CloseReadline"
	metaHintText      = "SetHintText"
	metaContinue      = "Continue"
)

func createMeta(line []rune, pos int) map[string]any {
	return map[string]any{
		metaHotKeyActions: []string{},
		metaSetLine:       string(line),
		metaSetPos:        pos,
		metaClose:         false,
		metaHintText:      "",
		metaContinue:      false,
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

func errInvalidMeta(property, dataType string, v any) ([]func(*readline.Instance), error) {
	return nil, fmt.Errorf("meta variable property '%s' is invalid: expecting %s instead got %T")
}

func validateMeta(v any) (functions []func(*readline.Instance), err error) {
	m, ok := v.(map[string]any)
	if !ok {
		return errInvalidMeta("$.", "map", v)
	}

	for property, value := range m {
		switch property {
		case metaHotKeyActions:
			actions, ok := value.([]string)
			if !ok {
				return errInvalidMeta(property, "array", value)
			}
			functions, err = compileActionSlice(actions)
			if err != nil {
				return nil, err
			}

		case metaSetLine, metaHintText:
			_, ok := value.(string)
			if !ok {
				return errInvalidMeta(property, types.String, value)
			}

		case metaSetPos:
			_, ok := value.(int)
			if !ok {
				return errInvalidMeta(property, types.Integer, value)
			}

		case metaClose, metaContinue:
			_, ok := value.(bool)
			if !ok {
				return errInvalidMeta(property, types.Boolean, value)
			}

		default:
			return nil, fmt.Errorf("invalid meta variable property: '$.%s'", property)
		}
	}

	return
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
