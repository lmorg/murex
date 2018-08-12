package streams

// template functions for stdio.Io methods to call
// (saves reinventing the wheel lots of times)

import (
	"io"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
	"github.com/lmorg/murex/lang/types"
)

func readArray(read stdio.Io, callback func([]byte)) error {
	dt := read.GetDataType()

	if ReadArray[dt] != nil {
		return ReadArray[dt](read, callback)
	}

	return ReadArray[types.Generic](read, callback)
}

func readMap(read stdio.Io, config *config.Config, callback func(key, value string, last bool)) error {
	dt := read.GetDataType()

	if ReadMap[dt] != nil {
		return ReadMap[dt](read, config, callback)
	}

	return ReadMap[types.Generic](read, config, callback)
}

// writeTo reads from the stream.Io interface and writes to a destination
// io.Writer interface
func writeTo(std stdio.Io, w io.Writer) (int64, error) {
	var (
		total int64
		i, n  int
		p     []byte = make([]byte, 1024*10)
		err   error
	)

	for {
		i, err = std.Read(p)

		if err == io.EOF {
			return total, nil
		}

		if err != nil {
			return total, err
		}

		n, err = w.Write(p[:i])
		total += int64(n)

		if err != nil {
			return total, err
		}

	}
}
