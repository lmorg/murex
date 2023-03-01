package paths

import (
	"bytes"
	"context"

	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/utils/consts"
)

func readArrayPath(ctx context.Context, read stdio.Io, callback func([]byte)) error {
	return readArray(ctx, read, callback, []byte(consts.PathSlash))
}

func readArrayPaths(ctx context.Context, read stdio.Io, callback func([]byte)) error {
	return readArray(ctx, read, callback, []byte{':'})
}

func readArray(ctx context.Context, read stdio.Io, callback func([]byte), separator []byte) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	split := bytes.Split(b, separator)
	for i := range split {

		select {
		case <-ctx.Done():
			return nil

		default:
			callback(split[i])
		}

	}

	return nil
}
