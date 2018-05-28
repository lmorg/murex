package json

import (
	"github.com/lmorg/murex/lang/proc/streams/stdio"
	"github.com/lmorg/murex/lang/types/define"
	"github.com/lmorg/murex/utils/json"
)

func readArray(read stdio.Io, callback func([]byte)) error {
	marshaller := func(v interface{}) ([]byte, error) {
		return json.Marshal(v, read.IsTTY())
	}

	return define.ArrayTemplate(marshaller, json.Unmarshal, read, callback)
}
