package encoders

import (
	"encoding/base64"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"io"
)

func init() {
	lang.GoFunctions["base64"] = lang.GoFunction{Func: cmdBase64, TypeIn: types.Generic, TypeOut: types.String}
	lang.GoFunctions["!base64"] = lang.GoFunction{Func: cmdUnbase64, TypeIn: types.Generic, TypeOut: types.Generic}
}

func cmdBase64(p *lang.Process) (err error) {
	encoder := base64.NewEncoder(base64.StdEncoding, p.Stdout)
	_, err = io.Copy(encoder, p.Stdin)
	//p.Stdin.WriteTo(encoder)

	encoder.Close()
	p.Stdout.Writeln([]byte{})
	return
}

func cmdUnbase64(p *lang.Process) (err error) {
	decoder := base64.NewDecoder(base64.StdEncoding, p.Stdin)
	_, err = io.Copy(p.Stdout, decoder)
	//p.Stdout.ReadFrom(decoder)
	return
}
