package lists

import (
	"bytes"
	"errors"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("match", cmdMatch, types.ReadArray, types.WriteArray)
	lang.DefineMethod("!match", cmdMatch, types.ReadArray, types.WriteArray)
}

func cmdMatch(p *lang.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	if p.Parameters.StringAll() == "" {
		return errors.New("no parameters supplied")
	}

	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	p.Stdin.ReadArray(p.Context, func(b []byte) {
		matched := bytes.Contains(b, p.Parameters.ByteAll())
		if (matched && !p.IsNot) || (!matched && p.IsNot) {
			err = aw.Write(b)
			if err != nil {
				p.Stdin.ForceClose()
				p.Done()
			}
		}
	})

	if p.HasCancelled() {
		return err
	}

	return aw.Close()
}
