# `ReadArray()` (type)

> Read from a data type one array element at a time

## Description

This is a function you would write when programming a Murex data-type.

It's called by builtins to allow them to read data structures one array element
at a time.

The purpose of this function is to allow builtins to support sequential reads
(where possible) and also create a standard interface for builtins, thus
allowing them to be data-type agnostic.

## Usage

Registering your `ReadArray()`

```go
// To avoid confusion, this should only happen inside func init()
stdio.RegisterReadArray(/* your type name */, /* your readArray func */)
```

## Examples

Example `ReadArray()` function:

```go
package string

import (
	"bufio"
	"bytes"
	"context"
	"fmt"

	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func readArray(ctx context.Context, read stdio.Io, callback func([]byte)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return scanner.Err()

		default:
			callback(bytes.TrimSpace(scanner.Bytes()))
		}
	}

	err := scanner.Err()
	if err != nil {
		return fmt.Errorf("error while reading a %s array: %s", types.String, err.Error())
	}

	return nil
}
```

## Detail

If your data type is not a stream-able array, it is then recommended that
you pass your array to `lang.ArrayTemplate()` which is a handler to convert Go
structures into Murex arrays. This also makes writing `ReadArray()` handlers
easier since you can just pass `lang.ArrayTemplate()` your marshaller.
For example:

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

The downside of this is that you're then unmarshalling the entire file, which
could be slow on large files and also breaks the streaming nature of UNIX
pipelines.

## Parameters

1. `stdio.Io`: stream to read from (eg STDIN)
2. `func([]byte)`: callback function. Each callback will be a []byte slice containing an array element

## See Also

- [apis/`ReadIndex()` (type)](/apis/ReadIndex.md):
  Data type handler for the index, `[`, builtin
- [apis/`ReadMap()` (type)](/apis/ReadMap.md):
  Treat data type as a key/value structure and read its contents
- [apis/`ReadNotIndex()` (type)](/apis/ReadNotIndex.md):
  Data type handler for the bang-prefixed index, `![`, builtin
- [apis/`WriteArray()` (type)](/apis/WriteArray.md):
  Write a data type, one array element at a time
- [apis/`lang.ArrayTemplate()` (template API)](/apis/lang.ArrayTemplate.md):
  Unmarshals a data type into a Go struct and returns the results as an array
- [apis/`lang.ArrayWithTypeTemplate()` (template API)](/apis/lang.ArrayWithTypeTemplate.md):
  Unmarshals a data type into a Go struct and returns the results as an array with data type included
