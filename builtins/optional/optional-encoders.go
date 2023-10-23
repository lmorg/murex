//go:build windows || opt_cmd_encoders
// +build windows opt_cmd_encoders

package optional

import (
	_ "github.com/lmorg/murex/builtins/optional/encoders" // base64, file archives, etc
)
