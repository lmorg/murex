# _murex_ Shell Docs

## API Reference: `ReadMap()` (type)

> Treat data type as a key/value structure and read its contents

## Description

This is a function you would write when programming a _murex_ data-type.

It's called by builtins to allow them to read data structures one key/value
pair at a time.

The purpose of this function is to allow builtins to support sequential reads
(where possible) and also create a standard interface for builtins, thus
allowing them to be data-type agnostic.

## Usage

Registering your `ReadMap()`

```go
// To avoid confusion, this should only happen inside func init()
stdio.RegisterReadMap(/* your type name */, /* your readMap func */)
```

## Examples

Example `ReadMap()` function:

```go
package json

import (
	"strconv"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/utils/json"
)

func readMap(read stdio.Io, _ *config.Config, callback func(key, value string, last bool)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	var jObj interface{}
	err = json.Unmarshal(b, &jObj)
	if err == nil {

		switch v := jObj.(type) {
		case []interface{}:
			for i := range jObj.([]interface{}) {
				j, err := json.Marshal(jObj.([]interface{})[i], false)
				if err != nil {
					return err
				}
				callback(strconv.Itoa(i), string(j), i != len(jObj.([]interface{}))-1)
			}

		case map[string]interface{}:
			i := 1
			for key := range jObj.(map[string]interface{}) {
				switch jObj.(map[string]interface{})[key].(type) {
				case string:
					callback(key, jObj.(map[string]interface{})[key].(string), i != len(jObj.(map[string]interface{})))

				default:
					j, err := json.Marshal(jObj.(map[string]interface{})[key], false)
					if err != nil {
						return err
					}
					callback(key, string(j), i != len(jObj.(map[string]interface{})))
				}
				i++
			}
			return nil

		default:
			if debug.Enabled {
				panic(v)
			}
		}
		return nil
	}
	return err
}
```

## Detail

There isn't (yet) a template read function for types to call. However that
might follow in a future release of _murex_.

## Parameters

1. `stdio.Io`: stream to read from (eg STDIN)
2. `*config.Config`: scoped config (eg your data type might have configurable parsing rules)
3. `func(key, value string, last bool)`: callback function: key and value of map plus boolean which is true if last element in row (eg reading from tables rather than key/values)

## See Also

* [apis/`ReadArray()` (type)](../apis/ReadArray.md):
  Read from a data type one array element at a time
* [apis/`ReadIndex()` (type)](../apis/ReadIndex.md):
  Data type handler for the index, `[`, builtin
* [apis/`ReadNotIndex()` (type)](../apis/ReadNotIndex.md):
  Data type handler for the bang-prefixed index, `![`, builtin
* [apis/`WriteArray()` (type)](../apis/WriteArray.md):
  Write a data type, one array element at a time