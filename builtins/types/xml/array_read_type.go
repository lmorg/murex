package xml

import (
	"context"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
)

func readArrayWithType(ctx context.Context, read stdio.Io, callback func(any, string)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	// Create a marshaller function to pass to ArrayDataTemplate
	marshaller := func(v any) ([]byte, error) {
		return MarshalTTY(v, read.IsTTY())
	}

	var v any
	unmarshaller(b, &v)

	r, ok := v.(map[string]any)
	if !ok || len(r) != 1 {
		return lang.ArrayDataWithTypeTemplate(ctx, typeName, marshaller, unmarshaller, v, callback)
	}

	var root string
	for root = range r {
	}

	e, ok := r[root].(map[string]any)
	if !ok || len(r) != 1 {
		marshaller = func(v any) ([]byte, error) {
			return marshalTTY(v, read.IsTTY(), root, xmlDefaultElement)
		}

		return lang.ArrayDataWithTypeTemplate(ctx, typeName, marshaller, unmarshaller, r[root], callback)
	}

	var element string
	for element = range e {
	}

	marshaller = func(v any) ([]byte, error) {
		return marshalTTY(v, read.IsTTY(), element, xmlDefaultElement)
	}

	return lang.ArrayDataWithTypeTemplate(ctx, typeName, marshaller, unmarshaller, e[element], callback)
}
