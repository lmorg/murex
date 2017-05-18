package encoders

import (
	"compress/gzip"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"io"
)

func init() {
	lang.GoFunctions["gz"] = lang.GoFunction{Func: cmdGz, TypeIn: types.Generic, TypeOut: types.Binary}
	lang.GoFunctions["!gz"] = lang.GoFunction{Func: cmdUngz, TypeIn: types.Binary, TypeOut: types.Generic}
}

func cmdGz(p *lang.Process) error {
	gz := gzip.NewWriter(p.Stdout)
	_, err := io.Copy(gz, p.Stdin)
	if err != nil {
		return err
	}

	gz.Close()

	return nil
}

func cmdUngz(p *lang.Process) error {
	gz, err := gzip.NewReader(p.Stdin)
	if err != nil {
		return err
	}

	_, err = io.Copy(p.Stdout, gz)
	if err != nil {
		return err
	}

	gz.Close()
	return nil
}
