package arraytools

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("addheading", cmdAddHeading, types.Unmarshal, types.Marshal)
}

func cmdAddHeading(p *lang.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	v, err := lang.UnmarshalData(p, dt)
	if err != nil {
		return err
	}

	switch t := v.(type) {
	case [][]string:
		v = append([][]string{p.Parameters.StringArray()}, t...)

	default:
		return fmt.Errorf("this doesn't appear to be a table or I don't know how to tabulate `%T`", v)
	}

	b, err := lang.MarshalData(p, dt, v)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
