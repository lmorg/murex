package parameters

import (
	"errors"
	"strconv"
	"strings"

	"github.com/lmorg/murex/lang/types"
)

var errTooFew = errors.New("too few parameters")

// Byte gets a single parameter as a []byte slice
func (p *Parameters) Byte(pos int) ([]byte, error) {
	p.mutex.RLock()

	if p.Len() <= pos {
		p.mutex.RUnlock()
		return []byte{}, errTooFew
	}
	b := []byte(p.params[pos])
	p.mutex.RUnlock()
	return b, nil
}

// ByteAll returns all parameters as one space-delimited []byte slice
func (p *Parameters) ByteAll() []byte {
	p.mutex.RLock()
	b := []byte(strings.Join(p.params, " "))
	p.mutex.RUnlock()
	return b
}

// ByteAllRange returns all parameters within range as one space-delimited []byte slice
// `start` is first point in array. `end` is last. Set `end` to `-1` if you want `[n:]`.
func (p *Parameters) ByteAllRange(start, end int) []byte {
	var b []byte
	p.mutex.RLock()

	if end == -1 {
		b = []byte(strings.Join(p.params[start:], " "))
	}
	b = []byte(strings.Join(p.params[start:end], " "))

	p.mutex.RUnlock()

	return b
}

// RuneArray gets all parameters as a [][]rune slice {parameter, characters}
func (p *Parameters) RuneArray() [][]rune {
	p.mutex.RLock()

	r := make([][]rune, len(p.params))
	for i := range p.params {
		r[i] = []rune(p.params[i])
	}
	p.mutex.RUnlock()
	return r
}

// String gets a single parameter as string
func (p *Parameters) String(pos int) (string, error) {
	p.mutex.RLock()

	if p.Len() <= pos {
		p.mutex.RUnlock()
		return "", errTooFew
	}

	s := p.params[pos]
	p.mutex.RUnlock()
	return s, nil
}

// StringArray returns all parameters as a slice of strings
func (p *Parameters) StringArray() []string {
	p.mutex.RLock()

	params := make([]string, len(p.params))
	copy(params, p.params)

	p.mutex.RUnlock()

	return params
}

// StringAll returns all parameters as one space-delimited string
func (p *Parameters) StringAll() string {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return strings.Join(p.params, " ")
}

// StringAllRange returns all parameters within range as one space-delimited string.
// `start` is first point in array. `end` is last. Set `end` to `-1` if you want `[n:]`.
func (p *Parameters) StringAllRange(start, end int) string {
	p.mutex.RLock()

	var s string

	switch {
	case len(p.params) == 0:
		s = ""
	case end == -1:
		s = strings.Join(p.params[start:], " ")
	default:
		s = strings.Join(p.params[start:end], " ")
	}

	p.mutex.RUnlock()
	return s
}

// Int gets parameter as integer
func (p *Parameters) Int(pos int) (int, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	if p.Len() <= pos {
		return 0, errTooFew
	}
	return strconv.Atoi(p.params[pos])
}

// Uint32 gets parameter as Uint32
func (p *Parameters) Uint32(pos int) (uint32, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	if p.Len() <= pos {
		return 0, errTooFew
	}
	i, err := strconv.ParseUint(p.params[pos], 10, 32)
	return uint32(i), err
}

// Bool gets parameter as boolean
func (p *Parameters) Bool(pos int) (bool, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	if p.Len() <= pos {
		return false, errTooFew
	}
	return types.IsTrue([]byte(p.params[pos]), 0), nil
}

// Block get parameter as a code block or JSON block
func (p *Parameters) Block(pos int) ([]rune, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	switch {
	case p.Len() <= pos:
		return []rune{}, errTooFew

	case len(p.params[pos]) < 2:
		return []rune{}, errors.New("not a valid code block. Too few characters. Code blocks should be surrounded by curly brace, eg `{ code }`")

	case p.params[pos][0] != '{':
		return []rune{}, errors.New("not a valid code block. Missing opening curly brace. Found `" + string([]byte{p.params[pos][0]}) + "` instead.")

	case p.params[pos][len(p.params[pos])-1] != '}':
		return []rune{}, errors.New("not a valid code block. Missing closing curly brace. Found `" + string([]byte{p.params[pos][len(p.params[pos])-1]}) + "` instead.")

	}

	return []rune(p.params[pos][1 : len(p.params[pos])-1]), nil
}

// Len returns the number of parameters
func (p *Parameters) Len() int {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return len(p.params)
}
