package onprompt

import (
	"fmt"
	"strings"

	ops "github.com/lmorg/murex/builtins/events/onPrompt/promptops"
)

var interrupts = []string{
	ops.Before,
	ops.After,
	ops.EOF,
	ops.Cancel,
}

func isValidInterrupt(interrupt string) error {
	for i := range interrupts {
		if interrupts[i] == interrupt {
			return nil
		}
	}

	return fmt.Errorf("invalid interrupt. Expecting one of the following: %s", strings.Join(interrupts, ", "))
}
