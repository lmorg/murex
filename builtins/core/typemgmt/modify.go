package typemgmt

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types/define"
)

func init() {
	proc.GoFunctions["append"] = cmdAppend
	proc.GoFunctions["prepend"] = cmdPrepend
}

func cmdPrepend(p *proc.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	var array []string

	err := p.Stdin.ReadArray(func(b []byte) {
		array = append(array, string(b))
	})

	if err != nil {
		return err
	}

	array = append(p.Parameters.StringArray(), array...)

	b, err := define.MarshalData(p, dt, array)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func cmdAppend(p *proc.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	var array []string

	err := p.Stdin.ReadArray(func(b []byte) {
		array = append(array, string(b))
	})

	if err != nil {
		return err
	}

	array = append(array, p.Parameters.StringArray()...)

	b, err := define.MarshalData(p, dt, array)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

/*func cmdUpdate(p *proc.Process) error {
	return nil
}*/
