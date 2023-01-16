package string

import (
	"bytes"
	"context"

	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func readArray(_ context.Context, read stdio.Io, callback func([]byte)) error {
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

func readArrayWithType(_ context.Context, read stdio.Io, callback func(interface{}, string)) error {
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
		callback(split[i], types.String)
	}

	return nil
}
