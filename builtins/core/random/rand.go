package random

import (
	"errors"
	"math/rand"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

const funcName = "rand"

func init() {
	lang.DefineFunction(funcName, cmdRand, types.Any)
}

func cmdRand(p *lang.Process) error {
	dt, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	p.Stdout.SetDataType(dt)
	var v any

	switch dt {
	case types.Integer, types.Number:
		max, _ := p.Parameters.Int(1)
		if max > 0 {
			v = rand.Intn(max + 1)
		} else {
			v = rand.Int()
		}

	case types.Float:
		v = rand.Float64()

	case types.String, types.Generic:
		max, _ := p.Parameters.Int(1)
		if max < 1 {
			max = 20
		}

		a := make([]rune, max)
		for i := 0; i < max; i++ {
			a[i] = rune(rand.Intn(126-31) + 32)
		}
		v = string(a)

	default:
		return errors.New("I don't know how to generate random data for the data type `" + dt + "`")
	}

	s, err := types.ConvertGoType(v, types.String)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write([]byte(s.(string)))
	return err
}
