# _murex_ Shell Docs

## API Reference: `Marshal()` 

> Converts structured memory into a structured file format (eg for stdio)

### Description

This is a function you would write when programming a _murex_ data-type.
The marshal function takes in a Go (golang) `type` or `struct` and returns
a byte slice of a "string" representation of that format (eg in JSON) or an
error.

This marshaller is then registered to _murex_ inside an `init()` function
and _murex_ builtins can use that marshaller via the `MarshalData()` API.

### Usage

Registering marshal (for writing builtin data-types)

```go
// To avoid data races, this should only happen inside func init()
lang.Marshallers["json"] = marshal
```

Using an existing marshaller (eg inside a builtin command)

```go
// See documentation on lang.MarshalData for more details
b, err := lang.MarshalData(p, dataType, data)
```

### Examples

Defining a marshaller for a murex data-type

```go
package example

import (
	"encoding/json"

	"github.com/lmorg/murex/lang"
)

func init() {
	// Register data-type
	lang.Marshallers["json"] = marshal
}

// Describe marshaller
func marshal(p *lang.Process, v interface{}) ([]byte, error) {
	if p.Stdout.IsTTY() {
		// If STDOUT is a TTY (ie not pipe, text file or other destination other
		// than a terminal) then output JSON in an indented, human readable,
		// format....
		return json.MarshalIndent(v, "", "    ")
	}

	// ....otherwise we might as well output it in a minified format
	return json.Marshal(v)
}
```

### Parameters

1. `*lang.Process`: Process's runtime state. Typically expressed as the variable `p` 
2. `interface{}`: data you wish to marshal

### See Also

* [apis/`Unmarshal()` ](../apis/unmarshal.md):
  Converts a structured file format into structured memory
* [apis/`lang.MarshalData()` ](../apis/marshaldata.md):
  Converts structured memory into a _murex_ data-type (eg for stdio)
* [apis/`lang.UnmarshalData()` ](../apis/unmarshaldata.md):
  Converts a _murex_ data-type into structured memory