package json

import (
	"context"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func readArrayWithType(ctx context.Context, read stdio.Io, callback func(any, string)) error {
	// Create a marshaller function to pass to ArrayWithTypeTemplate
	marshaller := func(v any) ([]byte, error) {
		return json.Marshal(v, read.IsTTY())
	}

	return lang.ArrayWithTypeTemplate(ctx, types.Json, marshaller, json.Unmarshal, read, callback)
}
