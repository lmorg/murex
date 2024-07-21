package events

import (
	"fmt"
	"strings"
)

func CompileInterruptKey(interrupt, name string) string {
	return fmt.Sprintf("%s.%s", name, interrupt)
}

type Key struct {
	Name      string
	Interrupt string
}

// GetInterruptFromKey returns: name, interrupt
func GetInterruptFromKey(key string) *Key {
	split := strings.SplitN(key, ".", 2)
	switch len(split) {
	case 2:
		return &Key{split[0], split[1]}
	case 1:
		return &Key{split[0], ""}
	default:
		return &Key{"", ""}
	}
}
