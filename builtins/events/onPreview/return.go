package onpreview

import (
	"fmt"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
)

const (
	retCacheCmdLine = "CacheCmdLine"
	retCacheTTL     = "CacheTTL"
	retDisplay      = "Display"
)

type eventReturn struct {
	CacheCmdLine bool
	CacheTTL     int
	Display      bool
}

func createReturn() map[string]any {
	return map[string]any{
		retCacheCmdLine: false,
		retCacheTTL:     60 * 60 * 24 * 30, // 30 days
		retDisplay:      true,
	}
}

func errInvalidReturn(property, dataType string, v any) (*eventReturn, error) {
	return nil, fmt.Errorf("$%s variable property '%s' is invalid: expecting %s instead got %T",
		consts.EventReturn, property, dataType, v)
}

func validateReturn(v any) (*eventReturn, error) {
	m, ok := v.(map[string]any)
	if !ok {
		return errInvalidReturn(consts.EventReturn, "map", v)
	}

	evtReturn := new(eventReturn)

	for property, value := range m {
		switch property {
		case retCacheCmdLine:
			incCmdLine, ok := value.(bool)
			if !ok {
				return errInvalidReturn(property, types.Boolean, value)
			}
			evtReturn.CacheCmdLine = incCmdLine

		case retCacheTTL:
			cacheTTL, ok := value.(int)
			if !ok {
				return errInvalidReturn(property, types.Integer, value)
			}
			evtReturn.CacheTTL = cacheTTL

		case retDisplay:
			display, ok := value.(bool)
			if !ok {
				return errInvalidReturn(property, types.Boolean, value)
			}
			evtReturn.Display = display

		default:
			return nil, fmt.Errorf("invalid %s property: '%s'", consts.EventReturn, property)
		}
	}

	return evtReturn, nil
}
