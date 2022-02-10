package arraytools

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("mtac", cmdMtac, types.Unmarshal, types.Marshal)
}

func cmdMtac(p *lang.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	v, err := lang.UnmarshalData(p, dt)
	if err != nil {
		return err
	}

	var b []byte

	switch v := v.(type) {
	case []string:
		for i := len(v)/2 - 1; i >= 0; i-- {
			opp := len(v) - 1 - i
			v[i], v[opp] = v[opp], v[i]
		}

	case [][]string:
		for i := len(v)/2 - 1; i >= 0; i-- {
			opp := len(v) - 1 - i
			v[i], v[opp] = v[opp], v[i]
		}

	case []int:
		for i := len(v)/2 - 1; i >= 0; i-- {
			opp := len(v) - 1 - i
			v[i], v[opp] = v[opp], v[i]
		}

	case []float64:
		for i := len(v)/2 - 1; i >= 0; i-- {
			opp := len(v) - 1 - i
			v[i], v[opp] = v[opp], v[i]
		}

	case []bool:
		for i := len(v)/2 - 1; i >= 0; i-- {
			opp := len(v) - 1 - i
			v[i], v[opp] = v[opp], v[i]
		}

	case []interface{}:
		for i := len(v)/2 - 1; i >= 0; i-- {
			opp := len(v) - 1 - i
			v[i], v[opp] = v[opp], v[i]
		}

	default:
		return fmt.Errorf("murex doesn't know how to read `%T` as an array. Please report this to https://github.com/lmorg/murex/issues", v)
	}

	b, err = lang.MarshalData(p, dt, v)
	if err != nil {
		return err
	}
	_, err = p.Stdout.Write(b)
	return err
}
