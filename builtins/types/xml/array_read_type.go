package xml

import (
	"context"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
)

func readArrayWithType(ctx context.Context, read stdio.Io, callback func(interface{}, string)) error {
	// Create a marshaller function to pass to ArrayWithTypeTemplate
	marshaller := func(v interface{}) ([]byte, error) {
		return MarshalTTY(v, read.IsTTY())
	}

	return lang.ArrayWithTypeTemplate(ctx, typeName, marshaller, unmarshaller, read, callback)
}
