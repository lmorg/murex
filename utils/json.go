package utils

import (
	"encoding/json"
	"errors"
)

const JsonNoData = "No data returned."

func JsonMarshal(obj interface{}, isTTY bool) (b []byte, err error) {
	if isTTY {
		b, err = json.MarshalIndent(obj, "", "\t")
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
