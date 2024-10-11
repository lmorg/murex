# `lang.MarshalData()` (system API)

> Converts structured memory into a Murex data-type (eg for stdio)

## Description



## Usage

```go
b, err := lang.MarshalData(p, dataType, data)
```

## Examples

```go
func exampleCommand(p *lang.Process) error {
    data := map[string]string {
        "foo": "hello foo",
        "bar": "hello bar",
    }

    dataType := "json"

    b, err := lang.MarshalData(p, dataType, data)
    if err != nil {
        return err
    }

    _, err := p.Stdout.Write(b)
    return err
}
```

## Detail

Go source file:

```go
package lang

import (
	"errors"
	"fmt"
)

// type DeMetaT func(any) error
type MarshallerT func(*Process, any) ([]byte, error)

var (
	//_demetaers = make(map[string]DeMetaT)

	// _marshallers defines the Go functions for converting a Go interface into a murex data type
	_marshallers = make(map[string]MarshallerT)
)

/*func RegisterDeMetaer(dataType string, demetaer DeMetaT) {
	if _demetaers[dataType] != nil {
		panic(fmt.Sprintf("premarshaller already exists for %s", dataType))
	}
	_demetaers[dataType] = demetaer
}

func DeMetaData(dataType string, data any) error {
	if _demetaers[dataType] == nil {
		return nil // we shouldn't assume there is a demetaer
	}

	return _demetaers[dataType](data)
}*/

func RegisterMarshaller(dataType string, marshaller MarshallerT) {
	if _marshallers[dataType] != nil {
		panic(fmt.Sprintf("marshaller already exists for %s", dataType))
	}
	_marshallers[dataType] = marshaller
}

// MarshalData is a global marshaller which should be called from within murex
// builtin commands (etc).
// See docs/apis/marshaldata.md for more details
func MarshalData(p *Process, dataType string, data any) ([]byte, error) {
	// This is one of the very few maps in Murex which isn't hidden behind a sync
	// lock of one description or other. The rational is that even mutexes can
	// add a noticeable overhead on the performance of tight loops and I expect
	// this function to be called _a lot_ while also only needing to be written
	// to via code residing in within builtin types init() function (ie while
	// murex is effectively single threaded). So there shouldn't be any data-
	// races -- PROVIDING developers strictly follow the pattern of only writing
	// to this map within init() func's.
	if _marshallers[dataType] == nil {
		return nil, errors.New("I don't know how to marshal `" + dataType + "`.")
	}

	b, err := _marshallers[dataType](p, data)
	if err != nil {
		return nil, errors.New("[" + dataType + " marshaller] " + err.Error())
	}

	return b, nil
}
```

## Parameters

1. `*lang.Process`: Process's runtime state. Typically expressed as the variable `p` 
2. `string`: Murex data type
3. `interface{}`: data you wish to marshal

## See Also

* [apis/`Marshal()` (type)](../apis/Marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [apis/`Unmarshal()` (type)](../apis/Unmarshal.md):
  Converts a structured file format into structured memory
* [apis/`lang.UnmarshalData()` (system API)](../apis/lang.UnmarshalData.md):
  Converts a Murex data-type into structured memory

<hr/>

This document was generated from [lang/define_marshal_doc.yaml](https://github.com/lmorg/murex/blob/master/lang/define_marshal_doc.yaml).