package pathsplit

import (
	"errors"
	"strings"
)

func Split(path string) ([]string, error) {
	switch len(path) {
	case 0:
		return nil, errors.New("empty path and missing separator")

	case 1:
		return nil, errors.New("path separator supplied but empty path followed")

	case 2:
		if path[0] == path[1] {
			return []string{""}, nil
		}
		fallthrough

	default:
		return strings.Split(path[1:], path[0:1]), nil
	}
}
