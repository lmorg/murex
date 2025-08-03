package lists

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("prepend", cmdPrepend, types.ReadArray, types.WriteArray)
	lang.DefineMethod("append", cmdAppend, types.ReadArray, types.WriteArray)
}

func cmdPrepend(p *lang.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	var (
		array    []any
		cachedDt string
	)

	err := p.Stdin.ReadArrayWithType(p.Context, func(v any, dt string) {
		array = append(array, v)
		cachedDt = dt
	})

	if err != nil {
		return err
	}

	var new []any
	params := p.Parameters.StringArray()
	for i := range params {
		v, err := types.ConvertGoType(params[i], cachedDt)
		if err != nil {
			return err
		}
		new = append(new, v)
	}
	array = append(new, array...)

	b, err := lang.MarshalData(p, dt, array)
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

	var (
		array    []any
		cachedDt string
	)

	err := p.Stdin.ReadArrayWithType(p.Context, func(v any, dt string) {
		array = append(array, v)
		cachedDt = dt
	})

	if err != nil {
		return err
	}

	params := p.Parameters.StringArray()
	for i := range params {
		v, err := types.ConvertGoType(params[i], cachedDt)
		if err != nil {
			return err
		}
		array = append(array, v)
	}

	b, err := lang.MarshalData(p, dt, array)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
