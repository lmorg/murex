package encoders

import (
	"encoding/base64"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"io"
)

func init() {
	lang.GoFunctions["base64"] = cmdBase64
	lang.GoFunctions["!base64"] = cmdUnbase64
}

func cmdBase64(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.String)
	encoder := base64.NewEncoder(base64.StdEncoding, p.Stdout)
	_, err = io.Copy(encoder, p.Stdin)

	encoder.Close()
	p.Stdout.Writeln([]byte{})
	return
}

func cmdUnbase64(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Generic)
	decoder := base64.NewDecoder(base64.StdEncoding, p.Stdin)
	_, err = io.Copy(p.Stdout, decoder)
	return
}
