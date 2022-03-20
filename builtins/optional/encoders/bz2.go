package encoders

import (
	"compress/bzip2"
	"io"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("!bz2", cmdUnbz2, types.Generic, types.Generic)
}

func cmdUnbz2(p *lang.Process) (err error) {
	if err = p.ErrIfNotAMethod(); err != nil {
		p.Stdout.SetDataType(types.Null)
		return err
	}

	p.Stdout.SetDataType(types.Generic)
	bz2 := bzip2.NewReader(p.Stdin)
	_, err = io.Copy(p.Stdout, bz2)
	if err != nil {
		return err
	}

	return nil
}
