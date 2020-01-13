package parser

import (
	"fmt"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

// safeCmds is a list of all the functions considered safe when using tab
// autocomplete with ExecCmdline
var safeCmds = []string{
	"open", "regexp", "match", "cast", "format", "[", "[[", "runtime",
	"cat", "ps", "grep", "ls",
}

func isCmdUnsafe(f string) bool {
	for _, sb := range safeCmds {
		if f == sb {
			return false
		}
	}

	return true
}

// GetSafeCmds returns a slice of the safeCmds
func GetSafeCmds() []string {
	return safeCmds
}

// ReadSafeCmds returns an interface{} of the safeCmds.
// This is only intended to be used by `config.Properties.GoFunc.Read()`
func ReadSafeCmds() (interface{}, error) {
	return safeCmds, nil
}

// WriteSafeCmds takes a JSON-encoded string and writes it to the safeCmds
// slice.
// This is only intended to be used by `config.Properties.GoFunc.Write()`
func WriteSafeCmds(v interface{}) error {
	switch v.(type) {
	case string:
		return json.Unmarshal([]byte(v.(string)), &safeCmds)

	default:
		return fmt.Errorf("Invalid data-type. Expecting a %s encoded string", types.Json)
	}
}
