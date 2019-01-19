package encoders

import (
	"compress/bzip2"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"io"
)

func init() {
	lang.GoFunctions["!bz2"] = cmdUnbz2
}

func cmdUnbz2(p *lang.Process) error {
	p.Stdout.SetDataType(types.Generic)
	bz2 := bzip2.NewReader(p.Stdin)
	_, err := io.Copy(p.Stdout, bz2)
	if err != nil {
		return err
	}

	return nil
}
