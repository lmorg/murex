package typemgmt

import (
	"errors"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("is-null", cmdIsNull, types.Boolean)
}

var (
	defined   = []byte("defined and not null")
	undefined = []byte("undefined or null")
)

func cmdIsNull(p *lang.Process) error {
	p.Stdout.SetDataType(types.Boolean)

	if p.IsMethod {
		return errors.New("this shouldn't be run as a method. Run `murex-docs is-null` for usage")
	}

	_, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	for _, name := range p.Parameters.StringArray() {

		p.Stdout.Write([]byte(name + ": "))

		dt := p.Variables.GetDataType(name)
		v, err := p.Variables.GetValue(name)

		switch {
		case dt == "", dt == types.Null:
			p.Stdout.Writeln(undefined)

		case err != nil:
			p.Stdout.Writeln([]byte(err.Error()))

		case v == nil:
			p.Stdout.Writeln(undefined)

		default:
			p.Stdout.Writeln(defined)
			p.ExitNum++
		}
	}

	return nil
}
