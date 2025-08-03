package xml

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func readMap(read stdio.Io, _ *config.Config, callback func(*stdio.Map)) error {
	// Create a marshaller function to pass to ArrayWithTypeTemplate
	marshaller := func(v any) ([]byte, error) {
		return MarshalTTY(v, read.IsTTY())
	}

	return lang.MapTemplate(types.Json, marshaller, unmarshaller, read, callback)
}
