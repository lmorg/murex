package stdio

// template functions for stdio.Io methods to call
// (saves reinventing the wheel lots of times)

import (
	"fmt"
	"io"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
)

// ReadArray is a template function for stdio.Io
func ReadArray(read Io, callback func([]byte)) error {
	dt := read.GetDataType()

	if readArray[dt] != nil {
		return readArray[dt](read, callback)
	}

	return readArray[types.Generic](read, callback)
}

// ReadArrayWithType is a template function for stdio.Io
func ReadArrayWithType(read Io, callback func([]byte, string)) error {
	dt := read.GetDataType()

	if readArrayWithType[dt] != nil {
		return readArrayWithType[dt](read, callback)
	}

	return readArrayWithType[types.Generic](read, callback)
}

// ReadMap is a template function for stdio.Io
func ReadMap(read Io, config *config.Config, callback func(key, value string, last bool)) error {
	dt := read.GetDataType()

	if readMap[dt] != nil {
		return readMap[dt](read, config, callback)
	}

	return readMap[types.Generic](read, config, callback)
}

// WriteArray is a template function for stdio.Io
func WriteArray(writer Io, dt string) (ArrayWriter, error) {
	if writeArray[dt] != nil {
		return writeArray[dt](writer)
	}

	return nil, fmt.Errorf("murex data type `%s` has not implemented WriteArray() method", dt)
}

// WriteTo is a template function for stdio.Io
func WriteTo(std Io, w io.Writer) (int64, error) {
	var (
		total int64
		i, n  int
		p     = make([]byte, 1024*10)
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
