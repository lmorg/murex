package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/lmorg/murex/utils/mxjson"
)

// Unmarshal is a wrapper around the standard json.Unmarshal function. This is
// done this way so that murex can swap out the JSON unmarshaller from the
// standard libraries with a 3rd party decoder that might run more efficiently.
func Unmarshal(data []byte, v any) (err error) {
	//err = gojay.Unmarshal(data, v)
	//if err == nil {
	//	return
	//}

	return json.Unmarshal(data, v)
}

// UnmarshalMurex is a wrapper around Go's JSON unmarshaller to support nested
// brace quotes (which allows for a cleaner syntax when embedding Murex code as
// JSON strings) and line comments via the hash, `#`, prefix.
func UnmarshalMurex(data []byte, v any) error {
	err := unmarshalMurex(data, v)
	if err == nil {
		return nil
	}

	_, mxerr := mxjson.Parse(data)
	return fmt.Errorf("mxjson parse error: %s\n%s", err, mxerr)
}

func unmarshalMurex(data []byte, v any) error {
	var (
		escape   bool
		comment  bool
		comments = make([]string, 1)
		single   bool
		double   bool
		brace    int
		replace  []string
		pop      []byte
	)

	for i := range data {
		if brace > 0 {
			pop = append(pop, data[i])
		}

		switch data[i] {
		case '\\':
			switch {
			case comment:
				comments[len(comments)-1] += string(data[i])
			default:
				escape = !escape
			}

		case '#':
			switch {
			case escape:
				escape = false
			case single, double, brace > 0:
				// do nothing
			default:
				comment = true
				comments[len(comments)-1] += string(data[i])
			}

		case '\'':
			switch {
			case comment:
				comments[len(comments)-1] += string(data[i])
			case escape:
				escape = false
			case double, brace > 0:
				// do nothing
			case single:
				single = false
			default:
				single = true
			}

		case '"':
			switch {
			case comment:
				comments[len(comments)-1] += string(data[i])
			case escape:
				escape = false
			case single, brace > 0:
				// do nothing
			case double:
				double = false
			default:
				double = true
			}

		case '(':
			switch {
			case comment:
				comments[len(comments)-1] += string(data[i])
			case escape:
				escape = false
			case single, double:
				// do nothing
			case brace == 0:
				//pop = append(pop, data[i])
				brace++
			default:
				brace++
			}

		case ')':
			switch {
			case comment:
				comments[len(comments)-1] += string(data[i])
			case escape:
				escape = false
			case single, double:
				// do nothing
			case brace == 1:
				replace = append(replace, string(pop[:len(pop)-1]))
				pop = []byte{}
				brace--
			default:
				brace--
			}

		case '\n':
			switch {
			case comment:
				comment = false
				comments = append(comments, "")
			case escape:
				escape = false
			}

		default:
			switch {
			case comment:
				comments[len(comments)-1] += string(data[i])
			default:
				escape = false
			}
		}

	}

	switch {
	case comment:
		comments = append(comments, "")
	case single:
		return errors.New("unterminated single quotes")
	case double:
		return errors.New("unterminated double quotes")
	case brace > 0:
		return fmt.Errorf("%d more open brace(s) than closed", brace)
	case brace < 0:
		return fmt.Errorf("%d more closed brace(s) than opened", brace)
	}

	// nothing to do so might as well just forward the params on without editing
	if len(replace) == 0 && len(comments) == 1 {
		return json.Unmarshal(data, v)
	}

	s := string(data)

	for i := range comments {
		s = strings.Replace(s, comments[i], "", 1)
	}

	for i := range replace {
		if len(replace[i]) > 0 {
			s = strings.Replace(s, "("+replace[i]+")", strconv.Quote(replace[i]), 1)
		}
	}

	return Unmarshal([]byte(s), v)
}
