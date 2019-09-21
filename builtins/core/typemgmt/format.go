package typemgmt

import (
	"errors"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["format"] = cmdFormat
}

func cmdFormat(p *lang.Process) (err error) {
	format, err := p.Parameters.String(0)
	if err != nil {
		return
	}

	dt := p.Stdin.GetDataType()

	if lang.Unmarshallers[dt] == nil {
		p.Stdout.SetDataType(types.Null)
		return errors.New("I don't know how to unmarshal `" + dt + "`.")
	}

	if lang.Marshallers[format] == nil {
		p.Stdout.SetDataType(types.Null)
		return errors.New("I don't know how to marshal `" + format + "`.")
	}

	v, err := lang.Unmarshallers[dt](p)
	if err != nil {
		p.Stdout.SetDataType(types.Null)
		return errors.New("[" + dt + " unmarshaller] " + err.Error())
	}

	b, err := lang.Marshallers[format](p, v)
	if err != nil {
		p.Stdout.SetDataType(types.Null)
		return errors.New("[" + format + " marshaller] " + err.Error())
	}

	p.Stdout.SetDataType(format)
	_, err = p.Stdout.Write(b)
	return
}
