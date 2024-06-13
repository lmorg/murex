//go:generate ./actions.mx

package onkeypress

import (
	"encoding/json"
	"fmt"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/readline"
)

const (
	metaHotKeyActions = "Actions"
	metaSetLine       = "SetLine"
	metaSetPos        = "SetCursorPosition"
	metaHintText      = "SetHintText"
	metaNextEvent     = "NextEvent"
)

func createMeta(line []rune, pos int) map[string]any {
	return map[string]any{
		metaHotKeyActions: []string{},
		metaSetLine:       string(line),
		metaSetPos:        pos,
		metaHintText:      "",
		metaNextEvent:     false,
	}
}

func errInvalidMeta(property, dataType string, v any) (*readline.EventReturn, error) {
	return nil, fmt.Errorf("meta variable property '%s' is invalid: expecting %s instead got %T",
		property, dataType, v)
}

func validateMeta(v any) (*readline.EventReturn, error) {
	m, ok := v.(map[string]any)
	if !ok {
		return errInvalidMeta("$.", "map", v)
	}

	var (
		evtReturn = new(readline.EventReturn)
		err       error
	)

	for property, value := range m {
		switch property {
		case metaHotKeyActions:
			var actions []string
			switch t := value.(type) {
			case string:
				err = json.Unmarshal([]byte(t), &actions)
				if err != nil {
					return nil, err
				}
			case []string:
				actions = t
			default:
				return errInvalidMeta(property, "array", value)
			}

			evtReturn.Actions, err = compileActionSlice(actions)
			if err != nil {
				return nil, err
			}

		case metaSetLine:
			s, ok := value.(string)
			if !ok {
				return errInvalidMeta(property, types.String, value)
			}
			evtReturn.SetLine = []rune(s)

		case metaHintText:
			s, ok := value.(string)
			if !ok {
				return errInvalidMeta(property, types.String, value)
			}
			evtReturn.HintText = []rune(s)

		case metaSetPos:
			i, ok := value.(int)
			if !ok {
				return errInvalidMeta(property, types.Integer, value)
			}
			evtReturn.SetPos = i

		case /*metaClose,*/ metaNextEvent:
			b, ok := value.(bool)
			if !ok {
				return errInvalidMeta(property, types.Boolean, value)
			}
			evtReturn.NextEvent = b

		default:
			return nil, fmt.Errorf("invalid meta variable property: '$.%s'", property)
		}
	}

	return evtReturn, nil
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
