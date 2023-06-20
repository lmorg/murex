package jsonconcat

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func readMap(read stdio.Io, _ *config.Config, callback func(*stdio.Map)) error {
	// Create a marshaller function to pass to ArrayWithTypeTemplate
	marshaller := func(v interface{}) ([]byte, error) {
		return json.Marshal(v, read.IsTTY())
	}

	return lang.MapTemplate(types.Json, marshaller, json.Unmarshal, read, callback)
}
