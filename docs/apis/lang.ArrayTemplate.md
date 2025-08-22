# `lang.ArrayTemplate()` (template API)

> Unmarshals a data type into a Go struct and returns the results as an array

## Description

This is a template API you can use for your custom data types to wrap around an
existing Go marshaller and return a Murex array which is consistent with
other structures such as nested JSON or YAML documents.

It should only be called from `ReadArray()` functions.

Because `lang.ArrayTemplate()` relies on a marshaller, it means any types that
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
	"github.com/lmorg/murex/utils"
)

// ArrayTemplate is a template function for reading arrays from marshalled data
func ArrayTemplate(ctx context.Context, marshal func(any) ([]byte, error), unmarshal func([]byte, any) error, read stdio.Io, callback func([]byte)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	if len(utils.CrLfTrim(b)) == 0 {
		return nil
	}

	var v interface{}
	err = unmarshal(b, &v)
	if err != nil {
		return err
	}

	return ArrayDataTemplate(ctx, marshal, unmarshal, v, callback)
}

func ArrayDataTemplate(ctx context.Context, marshal func(any) ([]byte, error), unmarshal func([]byte, any) error, data any, callback func([]byte)) error {
	switch v := data.(type) {
	case string:
		return readArrayByString(v, callback)

	case []string:
		return readArrayBySliceString(ctx, v, callback)

	case []interface{}:
		return readArrayBySliceInterface(ctx, marshal, v, callback)

	case map[string]string:
		return readArrayByMapStrStr(ctx, v, callback)

	case map[string]interface{}:
		return readArrayByMapStrIface(ctx, marshal, v, callback)

	case map[interface{}]string:
		return readArrayByMapIfaceStr(ctx, v, callback)

	case map[interface{}]interface{}:
		return readArrayByMapIfaceIface(ctx, marshal, v, callback)

	default:
		jBytes, err := marshal(v)
		if err != nil {

			return err
		}

		callback(jBytes)

		return nil
	}
}

func readArrayByString(v string, callback func([]byte)) error {
	callback([]byte(v))

	return nil
}

func readArrayBySliceString(ctx context.Context, v []string, callback func([]byte)) error {
	for i := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			callback([]byte(v[i]))
		}
	}

	return nil
}

func readArrayBySliceInterface(ctx context.Context, marshal func(interface{}) ([]byte, error), v []interface{}, callback func([]byte)) error {
	if len(v) == 0 {
		return nil
	}

	for i := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			switch v := v[i].(type) {
			case string:
				callback([]byte(v))

			case []byte:
				callback(v)

			default:
				jBytes, err := marshal(v)
				if err != nil {
					return err
				}
				callback(jBytes)
			}
		}
	}

	return nil
}

func readArrayByMapIfaceIface(ctx context.Context, marshal func(interface{}) ([]byte, error), v map[interface{}]interface{}, callback func([]byte)) error {
	for key, val := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			bKey := []byte(fmt.Sprint(key) + ": ")
			b, err := marshal(val)
			if err != nil {
				return err
			}
			callback(append(bKey, b...))
		}
	}

	return nil
}

func readArrayByMapStrStr(ctx context.Context, v map[string]string, callback func([]byte)) error {
	for key, val := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			callback([]byte(key + ": " + val))
		}
	}

	return nil
}

func readArrayByMapStrIface(ctx context.Context, marshal func(interface{}) ([]byte, error), v map[string]interface{}, callback func([]byte)) error {
	for key, val := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			bKey := []byte(key + ": ")
			b, err := marshal(val)
			if err != nil {
				return err
			}
			callback(append(bKey, b...))
		}
	}

	return nil
}

func readArrayByMapIfaceStr(ctx context.Context, v map[interface{}]string, callback func([]byte)) error {
	for key, val := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			callback([]byte(fmt.Sprint(key) + ": " + val))
		}
	}

	return nil
}
```

## Parameters

1. `func(interface{}) ([]byte, error)`: data type's marshaller
2. `func([]byte, interface{}) error`: data type's unmarshaller
3. `stdio.Io`: stream to read from (eg stdin)
4. `func([]byte)`: callback function to write each array element

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