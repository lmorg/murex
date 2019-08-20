# _murex_ Development Guide

## API Reference: `define.MarshalData()` 

> Converts structured memory into a _murex_ data-type (eg for stdio)

### Description



### Usage

```go
b, err := define.MarshalData(p, dataType, data)
```

### Examples

```go
func exampleCommand(p *lang.Process) error {
    data := map[string]string {
        "foo": "hello foo",
        "bar": "hello bar",
    }

    dataType := "json"

    b, err := define.MarshalData(p, dataType, data)
    if err != nil {
        return err
    }

    _, err := p.Stdout.Write(b)
    return err
}
```

### Detail

Go source file:

```go
package define

import (
	"errors"

	"github.com/lmorg/murex/lang"
)

// MarshalData is a global marshaller which should be called from within murex
// builtin commands (etc).
// See docs/apis/marshaldata.md for more details
func MarshalData(p *lang.Process, dataType string, data interface{}) (b []byte, err error) {

	if Marshallers[dataType] == nil {
		return nil, errors.New("I don't know how to marshal `" + dataType + "`.")
	}

	b, err = Marshallers[dataType](p, data)
	if err != nil {
		return nil, errors.New("[" + dataType + " marshaller] " + err.Error())
	}

	return
}
```

### Parameters

1. `*lang.Process`: Process's runtime state. Typically expressed as the variable `p
2. `string`: _murex_ data type
3. `interface{}`: data you wish to marshal

### See Also

* [`Marshal()` ](../apis/marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [`Unmarshal()` ](../apis/unmarshal.md):
  Converts a structured file format into structured memory
* [`define.UnmarshalData()` ](../apis/unmarshaldata.md):
  Converts a _murex_ data-type into structured memory