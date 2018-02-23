package utils

import (
	"testing"
)

// TestJsonMap tests the the JSON wrapper can marshal interface{} maps which the
// core library cannot
func TestJsonMap(t *testing.T) {
	obj := make(map[interface{}]interface{})
	obj["a"] = "b"
	obj[1] = 2

	b, err := JsonMarshal(obj, false)
	if err != nil {
		t.Error("Error marshalling: " + err.Error())
	}

	if string(b) != `{"a":"b","1":2}` && string(b) != `{"1":2,"a":"b"}` {
		t.Error("Unexpected JSON returned: " + string(b))
	}
}
