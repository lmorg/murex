package lang

import (
	"errors"
	"fmt"
)

// UnmarshalData is a global unmarshaller which should be called from within
// murex builtin commands (etc).
// See docs/apis/marshaldata.md for more details
func UnmarshalData(p *Process, dataType string) (v interface{}, err error) {
	// This is one of the very few maps in Murex which isn't hidden behind a sync
	// lock of one description or other. The rational is that even mutexes can
	// add a noticeable overhead on the performance of tight loops and I expect
	// this function to be called _a lot_ while also only needing to be written
	// to via code residing in within builtin types init() function (ie while
	// murex is effectively single threaded). So there shouldn't be any data-
	// races -- PROVIDING developers strictly follow the pattern of only writing
	// to this map within init() func's.
	if Unmarshallers[dataType] == nil {
		return nil, errors.New("I don't know how to unmarshal `" + dataType + "`")
	}

	v, err = Unmarshallers[dataType](p)
	if err != nil {
		return nil, errors.New("[" + dataType + " unmarshaller] " + err.Error())
	}

	return v, nil
}

func UnmarshalDataBuffered(parent *Process, b []byte, dataType string) (interface{}, error) {
	fork := parent.Fork(F_CREATE_STDIN | F_NO_STDOUT | F_NO_STDERR)
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
