package encoders

import (
	"encoding/base64"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"io"
)

func init() {
	proc.GoFunctions["base64"] = proc.GoFunction{Func: cmdBase64, TypeIn: types.Generic, TypeOut: types.String}
	proc.GoFunctions["!base64"] = proc.GoFunction{Func: cmdUnbase64, TypeIn: types.Generic, TypeOut: types.Generic}
}

func cmdBase64(p *proc.Process) (err error) {
	encoder := base64.NewEncoder(base64.StdEncoding, p.Stdout)
	_, err = io.Copy(encoder, p.Stdin)
	//p.Stdin.WriteTo(encoder)

	encoder.Close()
	p.Stdout.Writeln([]byte{})
	return
}

func cmdUnbase64(p *proc.Process) (err error) {
	decoder := base64.NewDecoder(base64.StdEncoding, p.Stdin)
	_, err = io.Copy(p.Stdout, decoder)
	//p.Stdout.ReadFrom(decoder)
	return
}
