package typemgmt

import (
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
)

func init() {
	proc.GoFunctions["format"] = cmdFormat
}

func cmdFormat(p *proc.Process) (err error) {
	format, err := p.Parameters.String(0)
	if err != nil {
		return
	}

	dt := p.Stdin.GetDataType()

	if define.Unmarshallers[dt] == nil {
		p.Stdout.SetDataType(types.Null)
		return errors.New("I don't know how to unmarshal `" + dt + "`.")
	}

	if define.Marshallers[format] == nil {
		p.Stdout.SetDataType(types.Null)
		return errors.New("I don't know how to marshal `" + format + "`.")
	}

	v, err := define.Unmarshallers[dt](p)
	if err != nil {
		p.Stdout.SetDataType(types.Null)
		return errors.New("[" + dt + " unmarshaller] " + err.Error())
	}

	b, err := define.Marshallers[format](p, v)
	if err != nil {
		p.Stdout.SetDataType(types.Null)
		return errors.New("[" + format + " marshaller] " + err.Error())
	}

	p.Stdout.SetDataType(format)
	_, err = p.Stdout.Write(b)
	return
}
