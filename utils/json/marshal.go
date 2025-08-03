package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// NoData is a custom default error message when JSON marshaller returns nil
const NoData = "no data returned"

// Marshal is a wrapper around Go's JSON marshaller to prettify output
// depending on whether the target is a terminal or not. This is so that the
// output is human readable when output for a human but a single line machine
// readable formatting for better support with iteration / concatenation when
// output to system functions.
func Marshal(v any, isTTY bool) (b []byte, err error) {
	b, err = marshal(v, isTTY)
	if err != nil && strings.Contains(err.Error(), "unsupported type: map[interface {}]interface {}") {
		b, err = marshal(deinterface(v), isTTY)
	}

	if err != nil {
		return
	}

	if string(b) == "null" {
		b = make([]byte, 0)
		return b, errors.New(NoData)
	}

	return
}

// marshal is a JSON marshaller which auto indents if output is a TTY
func marshal(v any, isTTY bool) (b []byte, err error) {
	//b, err = gojay.Marshal(v)
	//if err == nil {
	//	return
	//}

	if isTTY {
		b, err = json.MarshalIndent(v, "", "    ")
		return
	}

	b, err = json.Marshal(v)
	return
}

// deinterface is used to fudge around the lack of support for
// `map[any]any` in Go's JSON marshaller.
func deinterface(v any) any {
	switch t := v.(type) {
	case map[any]any:
		newV := make(map[string]any)
		for key := range t {
			newV[fmt.Sprint(key)] = deinterface(t[key])
		}
		//debug.Log(fmt.Sprintf("Deinterface: %T\n", t))
		return newV

	case []any:
		newV := make([]any, 0)
		for i := range t {
			newV = append(newV, deinterface(t[i]))
		}
		return newV

	default:
		//fmt.Printf("%T\n", t)
		return v
	}
}
