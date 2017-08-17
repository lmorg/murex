package utils

import (
	"encoding/json"
	"errors"
)

// JsonNoData is a custom default error message when JSON marshaller returns nil
const JsonNoData = "No data returned."

// Wrapper around Go's JSON marshaller to prettify output depending on whether the target is a terminal or not.
// This is so that the output is human readable when output for a human but a single line machine readable formatting
// for better support with iteration / concatenation when output to system functions.
func JsonMarshal(obj interface{}, isTTY bool) (b []byte, err error) {
	if isTTY {
		b, err = json.MarshalIndent(obj, "", "    ")
		if err != nil {
			return
		}

	} else {
		b, err = json.Marshal(obj)
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
