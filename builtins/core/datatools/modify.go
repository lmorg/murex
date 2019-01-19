package datatools

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types/define"
	"github.com/lmorg/murex/utils/alter"
)

func init() {
	lang.GoFunctions["append"] = cmdAppend
	lang.GoFunctions["prepend"] = cmdPrepend
	lang.GoFunctions["alter"] = cmdAlter
}

func cmdPrepend(p *lang.Process) error {
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

func cmdAppend(p *lang.Process) error {
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

func cmdAlter(p *lang.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	v, err := define.UnmarshalData(p, dt)
	if err != nil {
		return err
	}

	s, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	new, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	path, err := alter.SplitPath(s)
	if err != nil {
		return err
	}

	v, err = alter.Alter(p.Context, v, path, new)
	if err != nil {
		return err
	}

	b, err := define.MarshalData(p, dt, v)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
