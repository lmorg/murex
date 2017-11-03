package utils

import (
	"encoding/json"
	"errors"
)

// JsonNoData is a custom default error message when JSON marshaller returns nil
const JsonNoData = "No data returned."

// JsonMarshal is a wrapper around Go's JSON marshaller to prettify output depending on whether the target is a terminal
// or not. This is so that the output is human readable when output for a human but a single line machine readable
// formatting for better support with iteration / concatenation when output to system functions.
func JsonMarshal(v interface{}, isTTY bool) (b []byte, err error) {
	//v := deinterface(obj)
	if isTTY {
		b, err = json.MarshalIndent(v, "", "    ")
		if err != nil {
			return
		}

	} else {
		b, err = json.Marshal(v)
		if err != nil {
			return
		}
	}

	if string(b) == "null" {
		b = make([]byte, 0)
		return b, errors.New(JsonNoData)
	}

	return
}

// deinterface is used to fudge around the lack of support for `map[interface{}]interface{}` in Go's JSON marshaller.
/*func deinterface(v interface{}) interface{} {
	switch t := v.(type) {
	case map[interface{}]interface{}:
		newV := make(map[string]interface{})
		for key := range t {
			newV[fmt.Sprint(key)] = deinterface(t[key])
		}
		fmt.Printf("--> %T\n", t)
		return newV

	case []map[interface{}]interface{}:
		newA := make([]map[string]interface{}, 0)
		for m := range t {
			newM := make(map[string]interface{})
			for key := range t[m] {
				newM[fmt.Sprint(key)] = deinterface(t[m])
			}
			newA = append(newA, newM)
		}
		fmt.Printf("==> %T\n", t)
		return newA

	default:
		//fmt.Printf("%T\n", t)
		return v
	}
}*/
