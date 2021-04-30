# _murex_ Shell Docs

## API Reference: `lang.ArrayTemplate()` (template API)

> Unmarshals a data type into a Go struct and returns the results as an array

## Description

This is a template API you can use for your custom data types to wrap around an
existing Go marshaller and return a _murex_ array which is consistent with
other structures such as nested JSON or YAML documents.

It should only be called from `ReadArray()` functions.

Because `lang.ArrayTemplate()` relies on a marshaller, it means any types that
rely on this API are not going to be stream-able.



## Examples

Example calling `lang.ArrayTemplate()` function:

```go
package json

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/utils/json"
)

func readArray(read stdio.Io, callback func([]byte)) error {
	// Create a marshaller function to pass to ArrayTemplate
	marshaller := func(v interface{}) ([]byte, error) {
		return json.Marshal(v, read.IsTTY())
	}

	return lang.ArrayTemplate(marshaller, json.Unmarshal, read, callback)
}
```

## Detail

### API Source:

```go
package lang

import (
	"fmt"

	"github.com/lmorg/murex/lang/proc/stdio"
)

// ArrayTemplate is a template function for reading arrays from marshalled data
func ArrayTemplate(marshal func(interface{}) ([]byte, error), unmarshal func([]byte, interface{}) error, read stdio.Io, callback func([]byte)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	var v interface{}
	err = unmarshal(b, &v)

	if err != nil {
		return err
	}

	switch v.(type) {
	case string:
		return readArrayByString(v.(string), callback)

	case []string:
		return readArrayBySliceString(v.([]string), callback)

	case []interface{}:
		return readArrayBySliceInterface(marshal, v.([]interface{}), callback)

	case map[string]string:
		return readArrayByMapStrStr(v.(map[string]string), callback)

	case map[string]interface{}:
		return readArrayByMapStrIface(marshal, v.(map[string]interface{}), callback)

	case map[interface{}]string:
		return readArrayByMapIfaceStr(v.(map[interface{}]string), callback)

	case map[interface{}]interface{}:
		return readArrayByMapIfaceIface(marshal, v.(map[interface{}]interface{}), callback)

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

func readArrayBySliceString(v []string, callback func([]byte)) error {
	for i := range v {
		callback([]byte(v[i]))
	}

	return nil
}

func readArrayBySliceInterface(marshal func(interface{}) ([]byte, error), v []interface{}, callback func([]byte)) error {
	if len(v) == 0 {
		return nil
	}

	switch v[0].(type) {
	case string:
		for i := range v {
			callback([]byte(v[i].(string)))
		}

	case []byte:
		for i := range v {
			callback(v[i].([]byte))
		}

	default:
		for i := range v {

			jBytes, err := marshal(v[i])
			if err != nil {
				return err
			}
			callback(jBytes)
		}
	}

	return nil
}

func readArrayByMapIfaceIface(marshal func(interface{}) ([]byte, error), v map[interface{}]interface{}, callback func([]byte)) error {
	for key, val := range v {

		bKey := []byte(fmt.Sprint(key) + ": ")
		b, err := marshal(val)
		if err != nil {
			return err
		}
		callback(append(bKey, b...))
	}

	return nil
}

func readArrayByMapStrStr(v map[string]string, callback func([]byte)) error {
	for key, val := range v {

		callback([]byte(key + ": " + val))
	}

	return nil
}

func readArrayByMapStrIface(marshal func(interface{}) ([]byte, error), v map[string]interface{}, callback func([]byte)) error {
	for key, val := range v {

		bKey := []byte(key + ": ")
		b, err := marshal(val)
		if err != nil {
			return err
		}
		callback(append(bKey, b...))
	}

	return nil
}

func readArrayByMapIfaceStr(v map[interface{}]string, callback func([]byte)) error {
	for key, val := range v {

		callback([]byte(fmt.Sprint(key) + ": " + val))
	}

	return nil
}
```

## Parameters

1. `func(interface{}) ([]byte, error)`: data type's marshaller
2. `func([]byte, interface{}) error`: data type's unmarshaller
3. `stdio.Io`: stream to read from (eg STDIN)
4. `func([]byte)`: callback function to write each array element

## See Also

* [apis/`ReadArray()` (type)](../apis/ReadArray.md):
  Read from a data type one array element at a time
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