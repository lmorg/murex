# `lang.ArrayWithTypeTemplate()` (template API)

> Unmarshals a data type into a Go struct and returns the results as an array with data type included

## Description

This is a template API you can use for your custom data types to wrap around an
existing Go marshaller and return a Murex array which is consistent with
other structures such as nested JSON or YAML documents.

It should only be called from `ReadArrayWithType()` functions.

Because `lang.ArrayTemplateWithType()` relies on a marshaller, it means any types that
rely on this API are not going to be stream-able.



## Examples

Example calling `lang.ArrayTemplate()` function:

```go
package json

import (
	"context"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/utils/json"
)

func readArray(ctx context.Context, read stdio.Io, callback func([]byte)) error {
	// Create a marshaller function to pass to ArrayTemplate
	marshaller := func(v interface{}) ([]byte, error) {
		return json.Marshal(v, read.IsTTY())
	}

	return lang.ArrayTemplate(ctx, marshaller, json.Unmarshal, read, callback)
}
```

## Detail

### API Source:

```go
package lang

import (
	"context"
	"fmt"

	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/consts"
)

// ArrayWithTypeTemplate is a template function for reading arrays from marshalled data
func ArrayWithTypeTemplate(ctx context.Context, dataType string, marshal func(any) ([]byte, error), unmarshal func([]byte, any) error, read stdio.Io, callback func(any, string)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	if len(utils.CrLfTrim(b)) == 0 {
		return nil
	}

	var v any
	err = unmarshal(b, &v)
	if err != nil {
		return err
	}

	return ArrayDataWithTypeTemplate(ctx, dataType, marshal, unmarshal, v, callback)
}

func ArrayDataWithTypeTemplate(ctx context.Context, dataType string, marshal func(any) ([]byte, error), unmarshal func([]byte, any) error, data any, callback func(any, string)) error {
	switch v := data.(type) {
	case []any:
		return readArrayWithTypeBySliceInterface(ctx, dataType, marshal, v, callback)

	case []string:
		return readArrayWithTypeBySliceString(ctx, v, callback)

	case []float64:
		return readArrayWithTypeBySliceFloat(ctx, v, callback)

	case []int:
		return readArrayWithTypeBySliceInt(ctx, v, callback)

	case []bool:
		return readArrayWithTypeBySliceBool(ctx, v, callback)

	case string:
		return readArrayWithTypeByString(v, callback)

	case []byte:
		return readArrayWithTypeByString(string(v), callback)

	case []rune:
		return readArrayWithTypeByString(string(v), callback)

	case map[string]string:
		return readArrayWithTypeByMap(ctx, dataType, marshal, v, callback)

	case map[string]any:
		return readArrayWithTypeByMap(ctx, dataType, marshal, v, callback)

	case map[any]string:
		return readArrayWithTypeByMap(ctx, dataType, marshal, v, callback)

	case map[any]any:
		return readArrayWithTypeByMap(ctx, dataType, marshal, v, callback)

	case map[int]string:
		return readArrayWithTypeByMap(ctx, dataType, marshal, v, callback)

	case map[int]any:
		return readArrayWithTypeByMap(ctx, dataType, marshal, v, callback)

	case map[float64]string:
		return readArrayWithTypeByMap(ctx, dataType, marshal, v, callback)

	case map[float64]any:
		return readArrayWithTypeByMap(ctx, dataType, marshal, v, callback)

	default:
		return fmt.Errorf("cannot turn %T into an array\n%s", v, consts.IssueTrackerURL)
	}
}

func readArrayWithTypeByString(v string, callback func(any, string)) error {
	callback(v, types.String)

	return nil
}

func readArrayWithTypeBySliceInt(ctx context.Context, v []int, callback func(any, string)) error {
	for i := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			callback(v[i], types.Integer)
		}
	}

	return nil
}

func readArrayWithTypeBySliceFloat(ctx context.Context, v []float64, callback func(any, string)) error {
	for i := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			callback(v[i], types.Number)
		}
	}

	return nil
}

func readArrayWithTypeBySliceBool(ctx context.Context, v []bool, callback func(any, string)) error {
	for i := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			callback(v[i], types.Boolean)

		}
	}

	return nil
}

func readArrayWithTypeBySliceString(ctx context.Context, v []string, callback func(any, string)) error {
	for i := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			callback(v[i], types.String)
		}
	}

	return nil
}

func readArrayWithTypeBySliceInterface(ctx context.Context, dataType string, marshal func(any) ([]byte, error), v []any, callback func(any, string)) error {
	if len(v) == 0 {
		return nil
	}

	for i := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			switch v[i].(type) {

			case string:
				callback((v[i].(string)), types.String)

			case float64:
				callback(v[i].(float64), types.Number)

			case int:
				callback(v[i].(int), types.Integer)

			case bool:
				if v[i].(bool) {
					callback(true, types.Boolean)
				} else {
					callback(false, types.Boolean)
				}

			case []byte:
				callback(string(v[i].([]byte)), types.String)

			case []rune:
				callback(string(v[i].([]rune)), types.String)

			case nil:
				callback(nil, types.Null)

			default:
				jBytes, err := marshal(v[i])
				if err != nil {
					return err
				}
				callback(jBytes, dataType)
			}
		}
	}

	return nil
}

func readArrayWithTypeByMap[K comparable, V any](ctx context.Context, dataType string, marshal func(any) ([]byte, error), v map[K]V, callback func(any, string)) error {
	for key, val := range v {
		select {
		case <-ctx.Done():
			return nil
		default:
			m := map[K]any{key: val}
			b, err := marshal(m)
			if err != nil {
				return err
			}
			callback(string(b), dataType)
		}
	}

	return nil
}
```

## Parameters

1. `func(interface{}) ([]byte, error)`: data type's marshaller
2. `func([]byte, interface{}) error`: data type's unmarshaller
3. `stdio.Io`: stream to read from (eg stdin)
4. `func(interface{}, string)`: callback function to write each array element, with data type

## See Also

* [apis/`ReadArray()` (type)](../apis/ReadArray.md):
  Read from a data type one array element at a time
* [apis/`ReadArrayWithType()` (type)](../apis/ReadArrayWithType.md):
  Read from a data type one array element at a time and return the elements contents and data type
* [apis/`ReadIndex()` (type)](../apis/ReadIndex.md):
  Data type handler for the index, `[`, builtin
* [apis/`ReadMap()` (type)](../apis/ReadMap.md):
  Treat data type as a key/value structure and read its contents
* [apis/`ReadNotIndex()` (type)](../apis/ReadNotIndex.md):
  Data type handler for the bang-prefixed index, `![`, builtin
* [apis/`WriteArray()` (type)](../apis/WriteArray.md):
  Write a data type, one array element at a time
* [apis/`lang.IndexTemplateObject()` (template API)](../apis/lang.IndexTemplateObject.md):
  Returns element(s) from a data structure
* [apis/`lang.IndexTemplateTable()` (template API)](../apis/lang.IndexTemplateTable.md):
  Returns element(s) from a table

<hr/>

This document was generated from [lang/stdio/interface_doc.yaml](https://github.com/lmorg/murex/blob/master/lang/stdio/interface_doc.yaml).