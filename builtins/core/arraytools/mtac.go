package arraytools

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types/define"
)

func init() {
	lang.GoFunctions["mtac"] = cmdMtac
}

func cmdMtac(p *lang.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	v, err := define.UnmarshalData(p, dt)
	if err != nil {
		return err
	}

	var b []byte

	switch v.(type) {
	case []string:
		for i := len(v.([]string))/2 - 1; i >= 0; i-- {
			opp := len(v.([]string)) - 1 - i
			v.([]string)[i], v.([]string)[opp] = v.([]string)[opp], v.([]string)[i]
		}

	case [][]string:
		for i := len(v.([][]string))/2 - 1; i >= 0; i-- {
			opp := len(v.([][]string)) - 1 - i
			v.([][]string)[i], v.([][]string)[opp] = v.([][]string)[opp], v.([][]string)[i]
		}

	case []int:
		for i := len(v.([]int))/2 - 1; i >= 0; i-- {
			opp := len(v.([]int)) - 1 - i
			v.([]int)[i], v.([]int)[opp] = v.([]int)[opp], v.([]int)[i]
		}

	case []float64:
		for i := len(v.([]float64))/2 - 1; i >= 0; i-- {
			opp := len(v.([]float64)) - 1 - i
			v.([]float64)[i], v.([]float64)[opp] = v.([]float64)[opp], v.([]float64)[i]
		}

	case []bool:
		for i := len(v.([]bool))/2 - 1; i >= 0; i-- {
			opp := len(v.([]bool)) - 1 - i
			v.([]bool)[i], v.([]bool)[opp] = v.([]bool)[opp], v.([]bool)[i]
		}

	case []interface{}:
		for i := len(v.([]interface{}))/2 - 1; i >= 0; i-- {
			opp := len(v.([]interface{})) - 1 - i
			v.([]interface{})[i], v.([]interface{})[opp] = v.([]interface{})[opp], v.([]interface{})[i]
		}

	default:
		return fmt.Errorf("I don't know how to read %T as an array", v)
	}

	b, err = define.MarshalData(p, dt, v)
	if err != nil {
		return err
	}
	_, err = p.Stdout.Write(b)
	return err
}
