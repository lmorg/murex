package cache

import (
	"fmt"

	"github.com/lmorg/murex/lang/types"
)

var configCacheDisabled bool

// ReadMimes returns boolean
// This is only intended to be used by `config.Properties.GoFunc.Read()`
func ReadStatus() (any, error) {
	return !configCacheDisabled && !disabled, nil
}

// WriteStatus takes a bool
// This is only intended to be used by `config.Properties.GoFunc.Write()`
func WriteStatus(v any) error {
	v, err := types.ConvertGoType(v, types.Boolean)

	if err != nil {
		return err
	}

	boolean, ok := v.(bool)

	if !ok {
		return fmt.Errorf("cannot set cache enabled value because value is not a boolean")
	}

	configCacheDisabled = !boolean

	return nil
}
