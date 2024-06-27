package onpreview

import (
	"fmt"
	"strings"

	ops "github.com/lmorg/murex/builtins/events/onPreview/previewops"
)

var interrupts = []string{
	ops.Begin,
	ops.End,
	ops.Function,
	ops.Builtin,
	ops.Exec,
}

func isValidInterrupt(interrupt string) error {
	for i := range interrupts {
		if interrupts[i] == interrupt {
			return nil
		}
	}

	return fmt.Errorf("invalid interrupt. Expecting one of the following: %s", strings.Join(interrupts, ", "))
}
