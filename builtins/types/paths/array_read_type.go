package paths

import (
	"bytes"
	"context"

	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
)

func readArrayWithTypePath(ctx context.Context, read stdio.Io, callback func(interface{}, string)) error {
	return readArrayWithType(ctx, read, callback, []byte(consts.PathSlash), types.String)
}

func readArrayWithTypePaths(ctx context.Context, read stdio.Io, callback func(interface{}, string)) error {
	return readArrayWithType(ctx, read, callback, []byte{':'}, typePath)
}

func readArrayWithType(ctx context.Context, read stdio.Io, callback func(interface{}, string), separator []byte, retType string) error {
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
			callback(string(split[i]), retType)
		}

	}

	return nil
}
