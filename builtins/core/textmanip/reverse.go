package textmanip

import (
	"fmt"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types/define"
)

func init() {
	proc.GoFunctions["mtac"] = cmdMtac
}

func cmdMtac(p *proc.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	v, err := define.UnmarshalData(p, dt)
	if err != nil {
		return err
	}

	switch t := v.(type) {
	case []string:
		var rev []string
		for i := len(t) - 1; i > -1; i-- {
			rev = append(rev, t[i])
		}

		b, err := define.MarshalData(p, dt, &rev)
		if err != nil {
			return err
		}
		_, err = p.Stdout.Write(b)
		return err

	case []float64:
		var rev []float64
		for i := len(t) - 1; i > -1; i-- {
			rev = append(rev, t[i])
		}

		b, err := define.MarshalData(p, dt, &rev)
		if err != nil {
			return err
		}
		_, err = p.Stdout.Write(b)
		return err

	case []int:
		var rev []int
		for i := len(t) - 1; i > -1; i-- {
			rev = append(rev, t[i])
		}

		b, err := define.MarshalData(p, dt, &rev)
		if err != nil {
			return err
		}
		_, err = p.Stdout.Write(b)
		return err

	case []bool:
		var rev []bool
		for i := len(t) - 1; i > -1; i-- {
			rev = append(rev, t[i])
		}

		b, err := define.MarshalData(p, dt, &rev)
		if err != nil {
			return err
		}
		_, err = p.Stdout.Write(b)
		return err

	case []interface{}:
		var rev []interface{}
		for i := len(t) - 1; i > -1; i-- {
			rev = append(rev, t[i])
		}

		b, err := define.MarshalData(p, dt, &rev)
		if err != nil {
			return err
		}
		_, err = p.Stdout.Write(b)
		return err

	default:
		return fmt.Errorf("Cannot reverse Go-type %T", v)
	}
}
