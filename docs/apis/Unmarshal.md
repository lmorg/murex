# `Unmarshal()` (type)

> Converts a structured file format into structured memory

## Description

This is a function you would write when programming a Murex data-type.
The unmarshal function takes in a byte slice and returns a Go (golang)
`type` or `struct` or an error.

This unmarshaller is then registered to Murex inside an `init()` function
and Murex builtins can use that unmarshaller via the `UnmarshalData()`
API.

## Usage

Registering `Unmarshal()` (for writing builtin data-types)

```go
// To avoid data races, this should only happen inside func init()
lang.Unmarshallers[ /* your type name */ ] = /* your readIndex func */
```

Using an existing unmarshaller (eg inside a builtin command)

```go
// See documentation on lang.UnmarshalData for more details
v, err := lang.UnmarshalData(p *lang.Process, dataType string)
```

## Examples

Defining a marshaller for a murex data-type

```go
package example

import (
	"encoding/json"

	"github.com/lmorg/murex/lang"
)

func init() {
	// Register data-type
	lang.Unmarshallers["example"] = unmarshal
}

// Describe unmarshaller
func unmarshal(p *lang.Process) (interface{}, error) {
	// Read data from STDIN. Because JSON expects closing tokens, we should
	// read the entire stream before unmarshalling it. For formats like CSV or
	// jsonlines which are more line based, we might want to read STDIN line by
	// line. However given there is just one data return, you still effectively
	// head to read the entire file before returning the structure. There are
	// other APIs for iterative returns for streaming data - more akin to the
	// traditional way UNIX pipes would work.
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return nil, err
	}

	var v interface{}
	err = json.Unmarshal(b, &v)

	// Return the Go data structure or error
	return v, err
}
```

## Parameters

1. `*lang.Process`: Process's runtime state. Typically expressed as the variable `p`

## See Also

- [apis/`Marshal()` (type)](/apis/Marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
- [apis/`lang.MarshalData()` (system API)](/apis/lang.MarshalData.md):
  Converts structured memory into a Murex data-type (eg for stdio)
- [apis/`lang.UnmarshalData()` (system API)](/apis/lang.UnmarshalData.md):
  Converts a Murex data-type into structured memory
