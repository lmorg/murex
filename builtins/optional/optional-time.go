//go:build windows || opt_time
// +build windows opt_time

package optional

import (
	_ "github.com/lmorg/murex/builtins/optional/time" // sleep
)
