# `lang.UnmarshalData()` (system API)

> Converts a Murex data-type into structured memory

## Description



## Usage

```go
data, err := lang.UnmarshalData(p, dataType)
```

## Examples

```go
func exampleCommand(p *lang.Process) error {
    data := string `{ "foo": "hello foo", "bar": "hello bar" }`

    dataType := "json"

    v, err := lang.UnmarshalData(p, dataType)
    if err != nil {
        return err
    }

    s := fmt.Sprint(v)
    _, err := p.Stdout.Write([]byte(s))
    return err
}
```

## Detail

Go source file:

```go
package lang

import (
	"fmt"
)


type UnmarshallerT func(*Process) (any, error)

var (
	// _unmarshallers defines the Go functions for converting a murex data type into a Go interface
	_unmarshallers = make(map[string]UnmarshallerT)
)

func RegisterUnmarshaller(dataType string, unmarshaller UnmarshallerT) {
	if _unmarshallers[dataType] != nil {
		panic(fmt.Sprintf("unmarshaller already exists for %s", dataType))
	}
	_unmarshallers[dataType] = unmarshaller
}

// UnmarshalData is a global unmarshaller which should be called from within
// murex builtin commands (etc).
// See docs/apis/marshaldata.md for more details
func UnmarshalData(p *Process, dataType string) (v any, err error) {
	// This is one of the very few maps in Murex which isn't hidden behind a sync
	// lock of one description or other. The rational is that even mutexes can
	// add a noticeable overhead on the performance of tight loops and I expect
	// this function to be called _a lot_ while also only needing to be written
	// to via code residing in within builtin types init() function (ie while
	// murex is effectively single threaded). So there shouldn't be any data-
	// races -- PROVIDING developers strictly follow the pattern of only writing
	// to this map within init() func's.




	if _unmarshallers[dataType] == nil {
		return nil, fmt.Errorf("unknown data type. I don't know how to unmarshal `%s`", dataType)
	}

	v, err = _unmarshallers[dataType](p)
	if err != nil {
		return nil, fmt.Errorf("[%s unmarshaller] %s", dataType, err.Error())
	}

	return v, nil
}

func UnmarshalDataBuffered(parent *Process, b []byte, dataType string) (interface{}, error) {
	fork := parent.Fork(F_BACKGROUND | F_CREATE_STDIN | F_NO_STDOUT | F_NO_STDERR)
	defer fork.Kill()

	_, err := fork.Stdin.Write(b)
	if err != nil {
		return nil, fmt.Errorf("cannot write value to unmarshaller's buffer: %s", err.Error())
	}
	v, err := UnmarshalData(fork.Process, dataType)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal buffer: %s", err.Error())
	}

	return v, nil
}
```

## Parameters

1. `*lang.Process`: Process's runtime state. Typically expressed as the variable `p` 
2. `string`: Murex data type

## See Also

* [apis/`Marshal()` (type)](../apis/Marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [apis/`Unmarshal()` (type)](../apis/Unmarshal.md):
  Converts a structured file format into structured memory
* [apis/`lang.MarshalData()` (system API)](../apis/lang.MarshalData.md):
  Converts structured memory into a Murex data-type (eg for stdio)

<hr/>

This document was generated from [lang/define_unmarshal_doc.yaml](https://github.com/lmorg/murex/blob/master/lang/define_unmarshal_doc.yaml).