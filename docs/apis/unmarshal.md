# _murex_ Shell Docs

## API Reference: `Unmarshal()` 

> Converts a structured file format into structured memory

## Description

This is a function you would write when programming a _murex_ data-type.
The unmarshal function takes in a byte slice and returns a Go (golang)
`type` or `struct` or an error.

This unmarshaller is then registered to _murex_ inside an `init()` function
and _murex_ builtins can use that unmarshaller via the `UnmarshalData()`
API.

## Usage

Registering unmarshal (for writing builtin data-types)

```go
// To avoid data races, this should only happen inside func init()
lang.Unmarshallers["json"] = unmarshal
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

* [apis/`Unmarshal()` ](../apis/unmarshal.md):
  Converts a structured file format into structured memory
* [apis/`lang.MarshalData()` ](../apis/marshaldata.md):
  Converts structured memory into a _murex_ data-type (eg for stdio)
* [apis/unmashaldata](../apis/unmashaldata.md):
  