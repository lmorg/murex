package typemgmt

import (
	"errors"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("format", cmdFormat, types.Unmarshal, types.Marshal)
}

func cmdFormat(p *lang.Process) error {
	format, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	dt := p.Stdin.GetDataType()

	v, err := lang.UnmarshalData(p, dt)
	if err != nil {
		p.Stdout.SetDataType(types.Null)
		return errors.New("[" + dt + " unmarshaller] " + err.Error())
	}

	b, err := lang.MarshalData(p, format, v)
	if err != nil {
		p.Stdout.SetDataType(types.Null)
		return errors.New("[" + format + " marshaller] " + err.Error())
	}

	p.Stdout.SetDataType(format)
	_, err = p.Stdout.Write(b)
	return err
}
