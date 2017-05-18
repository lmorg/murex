package encoders

import (
	"compress/bzip2"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"io"
)

func init() {
	lang.GoFunctions["!bz2"] = lang.GoFunction{Func: cmdUnbz2, TypeIn: types.Generic, TypeOut: types.Generic}
}

func cmdUnbz2(p *lang.Process) error {
	bz2 := bzip2.NewReader(p.Stdin)
	_, err := io.Copy(p.Stdout, bz2)
	if err != nil {
		return err
	}

	return nil
}
