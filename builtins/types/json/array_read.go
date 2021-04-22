package json

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func readArray(read stdio.Io, callback func([]byte)) error {
	// Create a marshaller function to pass to ArrayTemplate
	marshaller := func(v interface{}) ([]byte, error) {
		return json.Marshal(v, read.IsTTY())
	}

	return lang.ArrayTemplate(marshaller, json.Unmarshal, read, callback)
}

func readArrayByType(read stdio.Io, callback func([]byte, string)) error {
	// Create a marshaller function to pass to ArrayByTypeTemplate
	marshaller := func(v interface{}) ([]byte, error) {
		return json.Marshal(v, read.IsTTY())
	}

	return lang.ArrayByTypeTemplate(types.Json, marshaller, json.Unmarshal, read, callback)
}
