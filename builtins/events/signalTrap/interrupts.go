package signaltrap

import (
	"fmt"
	"strings"
)

func isValidInterrupt(interrupt string) bool {
	_, ok := interrupts[interrupt]
	return ok
}

func compileInterruptKey(interrupt, name string) string {
	return fmt.Sprintf("%s_%s", interrupt, name)
}

func getInterruptFromKey(key string) []string {
	split := strings.SplitN(key, "_", 2)
	switch len(split) {
	case 2:
		return split
	case 1:
		return []string{"", split[0]}
	default:
		return []string{"", ""}
	}
}
