package sexp

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

var (
	// Strict(er) R.Rivset 1997 draft token + unicode letter support (hello 1997).
	// Strings and bytes matching this get marshaled as a token in non-canonical form.
	reMarshalToken = regexp.MustCompile(`^[\p{L}][\p{L}\p{N}\-./_:*+=]+$`)
)

func Unmarshal(data []byte) ([]interface{}, error) {
	l := NewLexer(data)
	exp, err := unmarshal(l)
	if err != nil {
		return nil, err
	}
	return exp.([]interface{})[0].([]interface{}), nil
}

func Marshal(data interface{}, canonical bool) ([]byte, error) {
	return marshal(data, canonical)
}

// The iterative way did my head in so I've gone recursive for now.
func unmarshal(l *lexer) (interface{}, error) {
	exp := make([]interface{}, 0)
	for item := l.Next(); item.Type != ItemEOF; item = l.Next() {
		switch item.Type {
		case ItemBracketLeft:
			subexp, err := unmarshal(l)
			if err != nil {
				return nil, err
			}
			exp = append(exp, subexp)
		case ItemBracketRight:
			return exp, nil
		case ItemToken, ItemQuote, ItemVerbatim:
			exp = append(exp, item.Value)
		case ItemError:
			return nil, fmt.Errorf("%s. Error was generated at position %d near '%s'", item.Value, l.pos, l.near())
		default:
			return nil, fmt.Errorf("Unexpected %s at position %d near '%s'", item, l.pos, l.near())
		}
	}
	return exp, nil
}

func marshal(data interface{}, canonical bool) ([]byte, error) {
	if d, ok := data.([]interface{}); ok {
		return marshal_slice(d, canonical)
	} else if d, ok := data.(string); ok {
		return marshal_bytes([]byte(d), canonical)
	} else if d, ok := data.([]byte); ok {
		return marshal_bytes(d, canonical)
	}
	return marshal_bytes([]byte(fmt.Sprintf("%v", data)), canonical)
}

func marshal_slice(data []interface{}, canonical bool) ([]byte, error) {
	exp := new(bytes.Buffer)
	exp.WriteString("(")
	separator := " "
	if canonical {
		separator = ""
	}
	atoms := [][]byte{}
	for _, sexp := range data {
		atom, err := marshal(sexp, canonical)
		if err != nil {
			return atom, err
		}
		atoms = append(atoms, atom)
	}
	exp.Write(bytes.Join(atoms, []byte(separator)))
	exp.WriteString(")")
	return exp.Bytes(), nil
}

func marshal_bytes(data []byte, canonical bool) ([]byte, error) {
	if canonical {
		return []byte(fmt.Sprintf("%d:%s", len(data), data)), nil
	} else if reMarshalToken.Match(data) {
		return data, nil
	}
	return []byte(strconv.Quote(string(data))), nil
}
