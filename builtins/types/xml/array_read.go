package xml

import (
	"context"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
)

func readArray(ctx context.Context, read stdio.Io, callback func([]byte)) error {
	// Create a marshaller function to pass to ArrayTemplate
	marshaller := func(v interface{}) ([]byte, error) {
		return MarshalTTY(v, read.IsTTY())
	}

	return lang.ArrayTemplate(ctx, marshaller, unmarshaller, read, callback)
}
