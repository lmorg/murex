package parameters

import (
	"errors"
	"github.com/lmorg/murex/lang/types"
	"strconv"
	"strings"
)

func (p Parameters) Byte(pos int) ([]byte, error) {
	if p.Len() <= pos {
		return []byte{}, errors.New("Too few parameters.")
	}
	return []byte(p.Params[pos]), nil
}

func (p Parameters) ByteAll() []byte {
	return []byte(strings.Join(p.Params, " "))
}

func (p Parameters) ByteAllRange(start, end int) []byte {
	if end == -1 {
		return []byte(strings.Join(p.Params[start:], " "))
	}
	return []byte(strings.Join(p.Params[start:end], " "))
}

func (p Parameters) String(pos int) (string, error) {
	if p.Len() <= pos {
		return "", errors.New("Too few parameters.")
	}
	return p.Params[pos], nil
}

func (p Parameters) StringArray() []string {
	return p.Params
}

func (p Parameters) StringAll() string {
	return strings.Join(p.Params, " ")
}

// `start` is first point in array. `end` is last. Set `end` to `-1` if you want `[n:]`.
func (p Parameters) StringAllRange(start, end int) string {
	if end == -1 {
		return strings.Join(p.Params[start:], " ")
	}
	return strings.Join(p.Params[start:end], " ")
}

func (p Parameters) Int(pos int) (int, error) {
	if p.Len() <= pos {
		return 0, errors.New("Too few parameters.")
	}
	return strconv.Atoi(p.Params[pos])
}

func (p Parameters) Bool(pos int) (bool, error) {
	if p.Len() <= pos {
		return false, errors.New("Too few parameters.")
	}
	return types.IsTrue([]byte(p.Params[pos]), 0), nil
}

func (p Parameters) Block(pos int) ([]rune, error) {
	switch {
	case p.Len() <= pos:
		return []rune{}, errors.New("Too few parameters.")

	case len(p.Params[pos]) < 2:
		return []rune{}, errors.New("Not a valid code block. Too few characters.")

	case p.Params[pos][0] != '{':
		return []rune{}, errors.New("Not a valid code block. Missing opening curly brace. Found `" + string([]byte{p.Params[pos][0]}) + "` instead.")

	case p.Params[pos][len(p.Params[pos])-1] != '}':
		return []rune{}, errors.New("Not a valid code block. Missing closing curly brace. Found `" + string([]byte{p.Params[pos][len(p.Params[pos])-1]}) + "` instead.")

	}

	return []rune(p.Params[pos][1 : len(p.Params[pos])-1]), nil
}

func (p Parameters) Len() int {
	return len(p.Params)
}
