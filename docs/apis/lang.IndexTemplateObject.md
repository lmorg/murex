# `lang.IndexTemplateObject()` (template API)

> Returns element(s) from a data structure

## Description

This is a template API you can use for your custom data types.

It should only be called from `ReadIndex()` and `ReadNotIndex()` functions.

This function ensures consistency with the index, `[`, builtin when used with
different Murex data types. Thus making indexing a data type agnostic
capability.



## Examples

Example calling `lang.IndexTemplateObject()` function:

```go
package json

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/json"
)

func index(p *lang.Process, params []string) error {
	var jInterface interface{}

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &jInterface)
	if err != nil {
		return err
	}

	marshaller := func(iface interface{}) ([]byte, error) {
		return json.Marshal(iface, p.Stdout.IsTTY())
	}

	return lang.IndexTemplateObject(p, params, &jInterface, marshaller)
}
```

## Detail

### API Source:

```go
package lang

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/lmorg/murex/lang/types"
)

type indexValueT interface {
	~string | ~bool | ~int | ~float64 | any
}

// IndexTemplateObject is a handy standard indexer you can use in your custom data types for structured object types.
// The point of this is to minimize code rewriting and standardising the behavior of the indexer.
func IndexTemplateObject(p *Process, params []string, object *any, marshaller func(any) ([]byte, error)) error {
	if p.IsNot {
		return itoNot(p, params, object, marshaller)
	}
	return itoIndex(p, params, object, marshaller)
}

// itoIndex allow
func itoIndex(p *Process, params []string, object *any, marshaller func(any) ([]byte, error)) error {
	switch v := (*object).(type) {
	case []any:
		return itoIndexArray(p, params, v, marshaller)

	case map[string]any:
		return itoIndexMap(p, params, v, marshaller)
	case map[any]any:
		return itoIndexMap(p, params, v, marshaller)

	default:
		return errors.New("object cannot be indexed")
	}
}

func itoIndexArray[V indexValueT](p *Process, params []string, v []V, marshaller func(any) ([]byte, error)) error {
	var objArray []V
	for _, key := range params {
		i, err := strconv.Atoi(key)
		if err != nil {
			return err
		}
		if i < 0 {
			i += len(v)
		}
		if i >= len(v) {
			return fmt.Errorf("key '%s' greater than number of items in array", key)
		}

		if len(params) > 1 {
			objArray = append(objArray, v[i])
			continue
		}

		switch value := any(v[i]).(type) {
		case nil:
			p.Stdout.SetDataType(types.Null)
		case bool:
			p.Stdout.SetDataType(types.Boolean)
			if value {
				p.Stdout.Write(types.TrueByte)
			} else {
				p.Stdout.Write(types.FalseByte)
			}
		case int:
			p.Stdout.SetDataType(types.Integer)
			s := strconv.Itoa(value)
			p.Stdout.Write([]byte(s))
		case float64:
			p.Stdout.SetDataType(types.Number)
			s := types.FloatToString(value)
			p.Stdout.Write([]byte(s))
		case string:
			p.Stdout.SetDataType(types.String)
			p.Stdout.Write([]byte(value))
		default:
			b, err := marshaller(value)
			if err != nil {
				return err
			}
			p.Stdout.Writeln(b)
		}
	}
	if len(objArray) > 0 {
		b, err := marshaller(objArray)
		if err != nil {
			return err
		}
		p.Stdout.Writeln(b)
	}
	return nil
}

func itoIndexMap[K comparable, V indexValueT](p *Process, params []string, v map[K]V, marshaller func(any) ([]byte, error)) error {
	var (
		objArray []any
		obj      any
		err      error
	)

	for i := range params {
		if len(params[i]) > 2 && params[i][0] == '[' && params[i][len(params[i])-1] == ']' {
			obj, err = ElementLookup(v, params[i][1:len(params[i])-1])
			if err != nil {
				return err
			}

		} else {
			var (
				iString int
				key     any
				ok      bool
			)

			for {
				switch iString {
				case 0:
					key = params[i]
				case 1:
					key = strings.Title(params[i])
				case 2:
					key = strings.ToLower(params[i])
				case 3:
					key = strings.ToUpper(params[i])
				default:
					return fmt.Errorf("key '%s' not found", params[i])
				}

				obj, ok = v[key.(K)]
				if ok {
					break
				}
				iString++
			}
		}

		if len(params) > 1 {
			objArray = append(objArray, obj)

		} else {
			switch obj := obj.(type) {
			case nil:
				p.Stdout.SetDataType(types.Null)
			case bool:
				p.Stdout.SetDataType(types.Boolean)
				if obj {
					p.Stdout.Write(types.TrueByte)
				} else {
					p.Stdout.Write(types.FalseByte)
				}
			case int:
				p.Stdout.SetDataType(types.Integer)
				s := strconv.Itoa(obj)
				p.Stdout.Write([]byte(s))
			case float64:
				p.Stdout.SetDataType(types.Number)
				s := types.FloatToString(obj)
				p.Stdout.Write([]byte(s))
			case string:
				p.Stdout.SetDataType(types.String)
				p.Stdout.Write([]byte(obj))
			default:
				b, err := marshaller(obj)
				if err != nil {
					return err
				}
				p.Stdout.Writeln(b)
			}
		}
	}
	if len(objArray) > 0 {
		b, err := marshaller(objArray)
		if err != nil {
			return err
		}
		p.Stdout.Writeln(b)
	}
	return nil

}

// itoNot requires the indexes to be explicit
func itoNot(p *Process, params []string, object *any, marshaller func(any) ([]byte, error)) error {
	switch v := (*object).(type) {
	case []any:
		return itoNotArray(p, params, v, marshaller)

	case map[string]any:
		return itoNotMap(p, params, v, marshaller)
	case map[any]any:
		return itoNotMap(p, params, v, marshaller)

	default:
		return errors.New("object cannot be !indexed")
	}
}

func itoNotArray[V indexValueT](p *Process, params []string, v []V, marshaller func(any) ([]byte, error)) error {
	var objArray []any
	not := make(map[int]bool)
	for _, key := range params {
		i, err := strconv.Atoi(key)
		if err != nil {
			return err
		}
		if i < 0 {
			return errors.New("cannot have negative keys in array")
		}
		if i >= len(v) {
			return fmt.Errorf("key '%s' greater than number of items in array", key)
		}

		not[i] = true
	}

	for i := range v {
		if !not[i] {
			objArray = append(objArray, v[i])
		}
	}

	b, err := marshaller(objArray)
	if err != nil {
		return err
	}
	_, err = p.Stdout.Writeln(b)

	return err
}

func itoNotMap[K comparable, V indexValueT](p *Process, params []string, v map[K]V, marshaller func(any) ([]byte, error)) error {
	objMap := make(map[K]any)
	not := make(map[K]bool)
	var key any

	for _, key = range params {
		not[key.(K)] = true

		key = strings.Title(key.(string))
		not[key.(K)] = true

		key = strings.ToLower(key.(string))
		not[key.(K)] = true

		key = strings.ToUpper(key.(string))
		not[key.(K)] = true
	}

	for s := range v {
		if !not[s] {
			objMap[s] = v[s]
		}
	}

	b, err := marshaller(objMap)
	if err != nil {
		return err
	}
	p.Stdout.Writeln(b)

	return nil
}
```

## Parameters

1. `*lang.Process`: Process's runtime state. Typically expressed as the variable `p` 
2. `[]string`: slice of parameters used in `[` / `![` 
3. `*interface{}`: a pointer to the data structure being indexed
4. `func(interface{}) ([]byte, error)`: data type marshaller function

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
* [apis/`lang.IndexTemplateTable()` (template API)](../apis/lang.IndexTemplateTable.md):
  Returns element(s) from a table
* [parser/index](../parser/item-index.md):
  Outputs an element from an array, map or table

<hr/>

This document was generated from [lang/stdio/interface_doc.yaml](https://github.com/lmorg/murex/blob/master/lang/stdio/interface_doc.yaml).