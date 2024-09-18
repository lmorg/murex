package lists

import (
	"errors"
	"fmt"
)

func splitRegexParams(regex []rune) ([]string, error) {
	if len(regex) < 2 {
		return nil, fmt.Errorf("invalid regexp (too few characters) in: `%s`", string(regex))
	}

	switch regex[1] {
	default:
		return splitRegexDefault(regex)

	case '{':
		return splitRegexBraces(regex)

	case '\\':
		return nil, fmt.Errorf("the `\\` character is not valid for separating regex parameters in: `%s`", string(regex))
	}
}

func splitRegexDefault(regex []rune) (s []string, _ error) {
	var (
		param   []rune
		escaped bool
		token   = regex[1]
	)

	for _, c := range regex {
		switch c {
		default:
			if escaped {
				param = append(param, '\\', c)
				escaped = false
				continue
			}
			param = append(param, c)

		case '\\':
			if escaped {
				param = append(param, '\\', c)
				escaped = false
				continue
			}
			escaped = true

		case token:
			if escaped {
				escaped = false
				param = append(param, c)
				continue
			}

			s = append(s, string(param))
			param = []rune{}
		}
	}
	s = append(s, string(param))

	return
}

func splitRegexBraces(regex []rune) ([]string, error) {
	return nil, errors.New("TODO")
}
