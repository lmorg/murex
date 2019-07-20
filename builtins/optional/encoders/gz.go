package encoders

import (
	"compress/gzip"
	"io"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["gz"] = cmdGz
	lang.GoFunctions["!gz"] = cmdUngz
}

func cmdGz(p *lang.Process) (err error) {
	if err = p.ErrIfNotAMethod(); err != nil {
		p.Stdout.SetDataType(types.Null)
		return err
	}

	p.Stdout.SetDataType(types.Binary)
	gz := gzip.NewWriter(p.Stdout)
	_, err = io.Copy(gz, p.Stdin)
	if err != nil {
		return err
	}

	gz.Close()

	return nil
}

func cmdUngz(p *lang.Process) error {
	if err := p.ErrIfNotAMethod(); err != nil {
		p.Stdout.SetDataType(types.Null)
		return err
	}

	p.Stdout.SetDataType(types.Generic)
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
