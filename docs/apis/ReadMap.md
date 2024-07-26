# `ReadMap()` (type)

> Treat data type as a key/value structure and read its contents

## Description

This is a function you would write when programming a Murex data-type.

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
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func readMap(read stdio.Io, _ *config.Config, callback func(*stdio.Map)) error {
	// Create a marshaller function to pass to ArrayWithTypeTemplate
	marshaller := func(v interface{}) ([]byte, error) {
		return json.Marshal(v, read.IsTTY())
	}

	return lang.MapTemplate(types.Json, marshaller, json.Unmarshal, read, callback)
}
```

## Detail

There isn't (yet) a template read function for types to call. However that
might follow in a future release of Murex.

## Parameters

1. `stdio.Io`: stream to read from (eg stdin)
2. `*config.Config`: scoped config (eg your data type might have configurable parsing rules)
3. `func(key, value string, last bool)`: callback function: key and value of map plus boolean which is true if last element in row (eg reading from tables rather than key/values)

## See Also

* [apis/`ReadArray()` (type)](../apis/ReadArray.md):
  Read from a data type one array element at a time
* [apis/`ReadArrayWithType()` (type)](../apis/ReadArrayWithType.md):
  Read from a data type one array element at a time and return the elements contents and data type
* [apis/`ReadIndex()` (type)](../apis/ReadIndex.md):
  Data type handler for the index, `[`, builtin
* [apis/`ReadNotIndex()` (type)](../apis/ReadNotIndex.md):
  Data type handler for the bang-prefixed index, `![`, builtin
* [apis/`WriteArray()` (type)](../apis/WriteArray.md):
  Write a data type, one array element at a time

<hr/>

This document was generated from [lang/stdio/interface_doc.yaml](https://github.com/lmorg/murex/blob/master/lang/stdio/interface_doc.yaml).