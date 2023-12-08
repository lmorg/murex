package cache

import (
	"fmt"

	"github.com/lmorg/murex/lang/types"
)

var configCache bool

// ReadMimes returns boolean
// This is only intended to be used by `config.Properties.GoFunc.Read()`
func ReadStatus() (interface{}, error) {
	return configCache && !disabled, nil
}

// WriteStatus takes a bool
// This is only intended to be used by `config.Properties.GoFunc.Write()`
func WriteStatus(v interface{}) error {
	switch v := v.(type) {
	case bool:
		configCache = v
		disabled = !v
		return nil

	default:
		return fmt.Errorf("invalid data-type. Expecting a %s", types.Boolean)
	}
}
