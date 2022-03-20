package encoders

import (
	"encoding/base64"
	"io"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("base64", cmdBase64, types.Any, types.String)
	lang.DefineMethod("!base64", cmdUnbase64, types.String, types.Generic)
}

func cmdBase64(p *lang.Process) (err error) {
	if err = p.ErrIfNotAMethod(); err != nil {
		p.Stdout.SetDataType(types.Null)
		return err
	}

	p.Stdout.SetDataType(types.String)
	encoder := base64.NewEncoder(base64.StdEncoding, p.Stdout)
	_, err = io.Copy(encoder, p.Stdin)

	encoder.Close()
	p.Stdout.Writeln([]byte{})
	return
}

func cmdUnbase64(p *lang.Process) (err error) {
	if err = p.ErrIfNotAMethod(); err != nil {
		p.Stdout.SetDataType(types.Null)
		return err
	}

	p.Stdout.SetDataType(types.Generic)
	decoder := base64.NewDecoder(base64.StdEncoding, p.Stdin)
	_, err = io.Copy(p.Stdout, decoder)
	return
}
