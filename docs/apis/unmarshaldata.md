# _murex_ Shell Guide

## API Reference: `define.UnmarshalData()` 

> Converts a _murex_ data-type into structured memory

### Description



### Usage

```go
data, err := define.UnmarshalData(p, dataType)
```

### Examples

```go
func exampleCommand(p *lang.Process) error {
    data := string `{ "foo": "hello foo", "bar": "hello bar" }`

    dataType := "json"

    v, err := define.UnmarshalData(p, dataType)
    if err != nil {
        return err
    }

    s := fmt.Sprint(v)
    _, err := p.Stdout.Write([]byte(s))
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

// UnmarshalData is a global unmarshaller which should be called from within
// murex builtin commands (etc).
// See docs/apis/marshaldata.md for more details
func UnmarshalData(p *lang.Process, dataType string) (v interface{}, err error) {
	// This is one of the very few maps in Murex which isn't hidden behind a sync
	// lock of one description or other. The rational is that even mutexes can
	// add a noticeable overhead on the performance of tight loops and I expect
	// this function to be called _a lot_ while also only needing to be written
	// to via code residing in within builtin types init() function (ie while
	// murex is effectively single threaded). So there shouldn't be any data-
	// races -- PROVIDING developers strictly follow the pattern of only writing
	// to this map within init() func's.
	if Unmarshallers[dataType] == nil {
		return nil, errors.New("I don't know how to unmarshal `" + dataType + "`.")
	}

	v, err = Unmarshallers[dataType](p)
	if err != nil {
		return nil, errors.New("[" + dataType + " unmarshaller] " + err.Error())
	}

	return v, nil
}
```

### Parameters

1. `*lang.Process`: Process's runtime state. Typically expressed as the variable `p
2. `string`: _murex_ data type

### See Also

* [`Marshal()` ](../apis/marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [`Unmarshal()` ](../apis/unmarshal.md):
  Converts a structured file format into structured memory
* [`define.MarshalData()` ](../apis/marshaldata.md):
  Converts structured memory into a _murex_ data-type (eg for stdio)