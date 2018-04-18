package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/lmorg/murex/debug"
)

// JsonNoData is a custom default error message when JSON marshaller returns nil
const JsonNoData = "No data returned."

// JsonMarshal is a wrapper around Go's JSON marshaller to prettify output
// depending on whether the target is a terminal or not. This is so that the
// output is human readable when output for a human but a single line machine
// readable formatting for better support with iteration / concatenation when
// output to system functions.
func JsonMarshal(v interface{}, isTTY bool) (b []byte, err error) {
	b, err = marshal(v, isTTY)
	if err != nil && strings.Contains(err.Error(), "unsupported type: map[interface {}]interface {}") {
		b, err = marshal(deinterface(v), isTTY)
	}

	if err != nil {
		return
	}

	if string(b) == "null" {
		b = make([]byte, 0)
		return b, errors.New(JsonNoData)
	}

	return
}

// marshal is a JSON marshaller which auto indents if output is a TTY
func marshal(v interface{}, isTTY bool) (b []byte, err error) {
	if isTTY {
		b, err = json.MarshalIndent(v, "", "    ")
		return
	}

	b, err = json.Marshal(v)
	return
}

// deinterface is used to fudge around the lack of support for
// `map[interface{}]interface{}` in Go's JSON marshaller.
func deinterface(v interface{}) interface{} {
	switch t := v.(type) {
	case map[interface{}]interface{}:
		newV := make(map[string]interface{})
		for key := range t {
			newV[fmt.Sprint(key)] = deinterface(t[key])
		}
		//debug.Log(fmt.Sprintf("Deinterface: %T\n", t))
		return newV

	case []interface{}:
		newV := make([]interface{}, 0)
		for i := range t {
			newV = append(newV, deinterface(t[i]))
		}
		return newV

	default:
		//fmt.Printf("%T\n", t)
		return v
	}
}

// JsonUnmarshal is a wrapper around Go's JSON unmarshaller to support nested
// brace quotes (which allows for a cleaner syntax when embedding Murex code as
// JSON strings) and line comments via the hash, `#`, prefix.
func JsonUnmarshal(data []byte, v interface{}) error {
	var (
		escape   bool
		comment  bool
		comments []string = make([]string, 1)
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
		return errors.New("Unterminated single quotes")
	case double:
		return errors.New("Unterminated double quotes")
	case brace > 0:
		return errors.New("More open braces than closed")
	case brace < 0:
		return errors.New("More closed braces than opened")
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

	debug.Json("Murex JSON comments", comments)
	debug.Json("Murex JSON braces", replace)
	debug.Log(string(data))
	debug.Log(s)

	return json.Unmarshal([]byte(s), v)
}
