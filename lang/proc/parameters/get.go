package parameters

import (
	"errors"
	"strconv"
	"strings"

	"github.com/lmorg/murex/lang/types"
)

// Byte gets a single parameter as a []byte slice
func (p Parameters) Byte(pos int) ([]byte, error) {
	if p.Len() <= pos {
		return []byte{}, errors.New("Too few parameters")
	}
	return []byte(p.Params[pos]), nil
}

// ByteAll returns all parameters as one space-delimited []byte slice
func (p Parameters) ByteAll() []byte {
	return []byte(strings.Join(p.Params, " "))
}

// ByteAllRange returns all parameters within range as one space-delimited []byte slice
// `start` is first point in array. `end` is last. Set `end` to `-1` if you want `[n:]`.
func (p Parameters) ByteAllRange(start, end int) []byte {
	if end == -1 {
		return []byte(strings.Join(p.Params[start:], " "))
	}
	return []byte(strings.Join(p.Params[start:end], " "))
}

// String gets a single parameter as string
func (p Parameters) String(pos int) (string, error) {
	if p.Len() <= pos {
		return "", errors.New("Too few parameters")
	}
	return p.Params[pos], nil
}

// StringArray returns all parameters as a slice of strings
func (p Parameters) StringArray() []string {
	return p.Params
}

// StringAll returns all parameters as one space-delimited string
func (p Parameters) StringAll() string {
	return strings.Join(p.Params, " ")
}

// StringAllRange returns all parameters within range as one space-delimited string.
// `start` is first point in array. `end` is last. Set `end` to `-1` if you want `[n:]`.
func (p Parameters) StringAllRange(start, end int) string {
	switch {
	case len(p.Params) == 0:
		return ""
	case end == -1:
		return strings.Join(p.Params[start:], " ")
	default:
		return strings.Join(p.Params[start:end], " ")
	}
}

// Int gets parameter as integer
func (p Parameters) Int(pos int) (int, error) {
	if p.Len() <= pos {
		return 0, errors.New("Too few parameters")
	}
	return strconv.Atoi(p.Params[pos])
}

// Uint32 gets parameter as Uint32
func (p Parameters) Uint32(pos int) (uint32, error) {
	if p.Len() <= pos {
		return 0, errors.New("Too few parameters")
	}
	i, err := strconv.ParseUint(p.Params[pos], 10, 32)
	return uint32(i), err
}

// Bool gets parameter as boolean
func (p Parameters) Bool(pos int) (bool, error) {
	if p.Len() <= pos {
		return false, errors.New("Too few parameters")
	}
	return types.IsTrue([]byte(p.Params[pos]), 0), nil
}

// Block get parameter as a code block or JSON block
func (p Parameters) Block(pos int) ([]rune, error) {
	switch {
	case p.Len() <= pos:
		return []rune{}, errors.New("Too few parameters")

	case len(p.Params[pos]) < 2:
		return []rune{}, errors.New("Not a valid code block. Too few characters")

	case p.Params[pos][0] != '{':
		return []rune{}, errors.New("Not a valid code block. Missing opening curly brace. Found `" + string([]byte{p.Params[pos][0]}) + "` instead.")

	case p.Params[pos][len(p.Params[pos])-1] != '}':
		return []rune{}, errors.New("Not a valid code block. Missing closing curly brace. Found `" + string([]byte{p.Params[pos][len(p.Params[pos])-1]}) + "` instead.")

	}

	return []rune(p.Params[pos][1 : len(p.Params[pos])-1]), nil
}

// Len returns the number of parameters
func (p Parameters) Len() int {
	return len(p.Params)
}

// TokenLen returns the number of parameters from ParamTokens. This method is
// not recommended for casual use - instead use Len() - however if you are
// monitoring other processes and those processes might be in a state prior to
// execution then you might want to check the parameter length before it has
// been parsed. This is this methods _only_ use case.
func (p Parameters) TokenLen() int {
	var count int
	for _, tokens := range p.Tokens {
		for i := range tokens {
			if tokens[i].Type != TokenTypeNil && tokens[i].Type != TokenTypeNamedPipe {
				count++
				break
			}
		}
	}

	return count
}
