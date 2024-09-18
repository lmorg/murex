package types

import (
	"fmt"
	"strings"
)

// Flatten takes anything that should look like a string, and turns it into a
// single string type.
func Flatten(v any) (string, error) {
	switch t := v.(type) {
	case string:
		return t, nil

	case []byte:
		return string(t), nil

	case []rune:
		return string(t), nil

	case []string:
		return strings.Join(t, "\n"), nil

	case [][]string:
		var s string
		for i := range t {
			s += strings.Join(t[i], "\t") + "\n"
		}
		return s, nil

	default:
		return "", fmt.Errorf("i don't know how to flatten %T", v)
	}
}
