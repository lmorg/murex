//go:generate ./actions.mx

package onkeypress

import (
	"encoding/json"
	"fmt"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/readline/v4"
)

const (
	retHotKeyActions = "Actions"
	retSetLine       = "SetLine"
	retSetPos        = "SetCursorPos"
	retHintText      = "SetHintText"
	retContinue      = "Continue"
)

func createReturn(state *readline.EventState) map[string]any {
	return map[string]any{
		retHotKeyActions: []string{},
		retSetLine:       state.Line,
		retSetPos:        state.CursorPos,
		retHintText:      "",
		retContinue:      false,
	}
}

func errInvalidReturn(property, dataType string, v any) (*readline.EventReturn, error) {
	return nil, fmt.Errorf("$%s variable property '%s' is invalid: expecting %s instead got %T",
		consts.EventReturn, property, dataType, v)
}

func validateReturn(v any) (*readline.EventReturn, error) {
	m, ok := v.(map[string]any)
	if !ok {
		return errInvalidReturn(consts.EventReturn, "map", v)
	}

	var (
		evtReturn = new(readline.EventReturn)
		err       error
	)

	for property, value := range m {
		switch property {
		case retHotKeyActions:
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
				return errInvalidReturn(property, "array", value)
			}

			evtReturn.Actions, err = compileActionSlice(actions)
			if err != nil {
				return nil, err
			}

		case retSetLine:
			s, ok := value.(string)
			if !ok {
				return errInvalidReturn(property, types.String, value)
			}
			evtReturn.SetLine = []rune(s)

		case retHintText:
			s, ok := value.(string)
			if !ok {
				return errInvalidReturn(property, types.String, value)
			}
			evtReturn.HintText = []rune(s)

		case retSetPos:
			i, ok := value.(int)
			if !ok {
				return errInvalidReturn(property, types.Integer, value)
			}
			evtReturn.SetPos = i

		case retContinue:
			cont, ok := value.(bool)
			if !ok {
				return errInvalidReturn(property, types.Boolean, value)
			}
			evtReturn.Continue = cont

		default:
			return nil, fmt.Errorf("invalid %s property: '%s'", consts.EventReturn, property)
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
