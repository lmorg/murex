package noglob

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

// noGlobCmds is a list of all the functions considered unsafe to expand with
// globbing
var noGlobCmds = []string{
	"rx", "!rx", "g", "!g", "cast", "format", "select", "!regexp", "regexp",
	"find", "[", "![", "[[",
	lang.ExpressionFunctionName,
}

func canGlobCmd(f string) bool {
	for _, sb := range noGlobCmds {
		if f == sb {
			return true
		}
	}

	return false
}

// GetNoGlobCmds returns a slice of the noGlobCmds
func GetNoGlobCmds() []string {
	a := make([]string, len(noGlobCmds))
	copy(a, noGlobCmds)
	return a
}

// ReadNoGlobCmds returns an interface{} of the noGlobCmds.
// This is only intended to be used by `config.Properties.GoFunc.Read()`
func ReadNoGlobCmds() (interface{}, error) {
	return GetNoGlobCmds(), nil
}

// WriteNoGlobCmds takes a JSON-encoded string and writes it to the noGlobCmds
// slice.
// This is only intended to be used by `config.Properties.GoFunc.Write()`
func WriteNoGlobCmds(v interface{}) error {
	switch v := v.(type) {
	case string:
		return json.Unmarshal([]byte(v), &noGlobCmds)

	default:
		return fmt.Errorf("invalid data-type. Expecting a %s encoded string", types.Json)
	}
}
