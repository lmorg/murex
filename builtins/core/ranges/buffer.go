package ranges

import (
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
)

func buffer(p *lang.Process, dt string) (stdio.Io, int, error) {
	stdin := streams.NewStdinWithContext(p.Context, p.Done)
	stdin.SetDataType(dt)

	array, err := stdin.WriteArray(dt)
	if err != nil {
		return nil, 0, err
	}

	var (
		nestedErr error
		length    int
	)

	err = p.Stdin.ReadArray(p.Context, func(b []byte) {
		nestedErr = array.Write(b)
		if nestedErr != nil {
			p.Done()
			return
		}
		length++
	})

	if nestedErr != nil {
		return nil, 0, nestedErr
	}

	if err != nil {
		return nil, 0, err
	}

	return stdin, length, array.Close()
}
