package json

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

// TestJsonMap tests the the JSON wrapper can marshal interface{} maps which the
// core library cannot
func TestJsonMap(t *testing.T) {
	count.Tests(t, 1, "TestJsonMap")

	obj := make(map[interface{}]interface{})
	obj["a"] = "b"
	obj[1] = 2

	b, err := Marshal(obj, false)
	if err != nil {
		t.Error("Error marshalling: " + err.Error())
	}

	if string(b) != `{"a":"b","1":2}` && string(b) != `{"1":2,"a":"b"}` {
		t.Error("Unexpected JSON returned: " + string(b))
	}
}
