# _murex_ Shell Guide

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
	// This is one of the very few maps in Murex which isn't hidden behind a sync
	// lock of one description or other. The rational is that even mutexes can
	// add a noticeable overhead on the performance of tight loops and I expect
	// this function to be called _a lot_ while also only needing to be written
	// to via code residing in within builtin types init() function (ie while
	// murex is effectively single threaded). So there shouldn't be any data-
	// races -- PROVIDING developers strictly follow the pattern of only writing
	// to this map within init() func's.
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

1. `*lang.Process`: Process's runtime state. Typically expressed as the variable `p` 
2. `string`: _murex_ data type
3. `interface{}`: data you wish to marshal

### See Also

* apis/[`Marshal()` ](../apis/marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* apis/[`Unmarshal()` ](../apis/unmarshal.md):
  Converts a structured file format into structured memory
* apis/[`define.UnmarshalData()` ](../apis/unmarshaldata.md):
  Converts a _murex_ data-type into structured memory