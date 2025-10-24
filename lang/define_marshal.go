package lang

import (
	"errors"
	"fmt"
)

type MarshallerT func(*Process, any) ([]byte, error)

var (
	// _marshallers defines the Go functions for converting a Go interface into a murex data type
	_marshallers = make(map[string]MarshallerT)
)

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
