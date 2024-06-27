package ansi

import (
	"bytes"
	"fmt"
)

func GetConsts(p []byte) string {
	switch len(p) {
	case 0:
		return ""
	case 1:
		return getConsts(rune(p[0]))
	}

	for constant, code := range constants {
		if bytes.Equal(code, p) {
			return fmt.Sprintf("{%s}", constant)
		}
	}

	var s string
	for _, r := range string(p) {
		s += getConsts(r)
	}
	return s
}

func getConsts(r rune) string {
	switch r {
	case 4:
		return "{EOF}"
	case 7:
		return "{BELL}"
	case 10:
		return "{CR}"
	case 13:
		return "{LF}"
	case 27:
		return "{ESC}"

	default:
		for constant, code := range constants {
			if len(code) == 1 && rune(code[0]) == r {
				return fmt.Sprintf("{%s}", constant)
			}
		}

		return string(r)
	}
}
