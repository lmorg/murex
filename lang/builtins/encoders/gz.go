package encoders

import (
	"compress/gzip"
	"io"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["gz"] = proc.GoFunction{Func: cmdGz, TypeIn: types.Generic, TypeOut: types.Binary}
	proc.GoFunctions["gz!"] = proc.GoFunction{Func: cmdUngz, TypeIn: types.Binary, TypeOut: types.Generic}
}

func cmdGz(p *proc.Process) error {
	gz := gzip.NewWriter(p.Stdout)
	_, err := io.Copy(gz, p.Stdin)
	if err != nil {
		return err
	}

	gz.Close()

	return nil
}

func cmdUngz(p *proc.Process) error {
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
