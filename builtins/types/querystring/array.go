package string

import (
	"bytes"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
)

func readArray(read stdio.Io, callback func([]byte)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	if len(b) == 0 {
		return nil
	}

	if b[0] == '?' {
		if len(b) == 1 {
			return nil
		}
		b = b[1:]
	}

	split := bytes.Split(b, []byte{'&'})
	for i := range split {
		callback(split[i])
	}

	return nil
}
