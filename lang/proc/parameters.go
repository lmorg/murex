package proc

import (
	"errors"
	"github.com/lmorg/murex/lang/types"
	"strconv"
	"strings"
)

type Parameters []string

func (p Parameters) AllByte() []byte {
	return []byte(strings.Join(p, " "))
}

func (p Parameters) AllString() string {
	return strings.Join(p, " ")
}

func (p Parameters) Int(pos int) (int, error) {
	if p.Len() <= pos {
		return 0, errors.New("Too few parameters.")
	}
	return strconv.Atoi(p[pos])
}

func (p Parameters) String(pos int) (string, error) {
	if p.Len() <= pos {
		return "", errors.New("Too few parameters.")
	}
	return p[pos], nil
}

func (p Parameters) Bool(pos int) (bool, error) {
	if p.Len() <= pos {
		return false, errors.New("Too few parameters.")
	}
	return types.IsTrue([]byte(p[pos]), 0), nil
}

func (p Parameters) Block(pos int) ([]rune, error) {
	switch {
	case p.Len() <= pos:
		return []rune{}, errors.New("Too few parameters.")

	case len(p[pos]) < 2:
		return []rune{}, errors.New("Not a valid code block. Too few characters.")

	case p[pos][0] != '{':
		return []rune{}, errors.New("Not a valid code block. Missing opening curly brace. Found " + string([]byte{p[pos][0]}) + " instead.")

	case p[pos][len(p[pos])-1] != '}':
		return []rune{}, errors.New("Not a valid code block. Missing closing curly brace. Found " + string([]byte{p[pos][len(p[pos])-1]}) + " instead.")

	}

	return []rune(p[pos][1 : len(p[pos])-1]), nil
}

func (p Parameters) Len() int {
	/*if len(p) == 1 && len(p[0]) == 0 {
		return 0
	}*/
	return len(p)
}

/*
func (p Parameters) Last() Parameter {
	return Parameter(p[len(p)-1])
}
*/
