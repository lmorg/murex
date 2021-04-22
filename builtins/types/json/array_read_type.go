package json

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func readArrayWithType(read stdio.Io, callback func([]byte, string)) error {
	// Create a marshaller function to pass to ArrayWithTypeTemplate
	marshaller := func(v interface{}) ([]byte, error) {
		return json.Marshal(v, read.IsTTY())
	}

	return lang.ArrayWithTypeTemplate(types.Json, marshaller, json.Unmarshal, read, callback)
}
