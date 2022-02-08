package alter

import (
	"errors"
	"strings"
)

// SplitPath takes a string with a prefixed delimiter and separates it into a slice of path elements
func SplitPath(path string) ([]string, error) {
	split := strings.Split(path, string(path[0]))
	if len(split) == 0 || (len(split) == 1 && split[0] == "") {
		return nil, errors.New("Empty path")
	}

	if split[0] == "" {
		split = split[1:]
	}

	return split, nil
}
