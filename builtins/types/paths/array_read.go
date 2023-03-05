package paths

import (
	"bytes"
	"context"

	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/utils/path"
)

func readArrayPath(ctx context.Context, read stdio.Io, callback func([]byte)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	split, err := path.SplitPath(b)
	if split == nil {
		return err
	}

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

func readArrayPaths(ctx context.Context, read stdio.Io, callback func([]byte)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	split := bytes.Split(b, pathsSeparator)
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
