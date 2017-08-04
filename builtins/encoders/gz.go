package encoders

import (
	"compress/gzip"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"io"
)

func init() {
	proc.GoFunctions["gz"] = cmdGz
	proc.GoFunctions["!gz"] = cmdUngz
}

func cmdGz(p *proc.Process) error {
	p.Stdout.SetDataType(types.Binary)
	gz := gzip.NewWriter(p.Stdout)
	_, err := io.Copy(gz, p.Stdin)
	if err != nil {
		return err
	}

	gz.Close()

	return nil
}

func cmdUngz(p *proc.Process) error {
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
