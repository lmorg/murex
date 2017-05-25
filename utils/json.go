package utils

import (
	"encoding/json"
	"errors"
)

const JsonNoData = "No data returned."

func JsonMarshal(obj interface{}) (b []byte, err error) {
	b, err = json.MarshalIndent(obj, "", "\t")
	if err != nil {
		return
	}

	if string(b) == "null" {
		b = make([]byte, 0)
		err = errors.New(JsonNoData)
	}

	return
}
