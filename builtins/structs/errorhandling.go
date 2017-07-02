package structs

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"io"
)

func init() {
	proc.GoFunctions["try"] = proc.GoFunction{Func: cmdTry, TypeIn: types.Null, TypeOut: types.Generic}
	proc.GoFunctions["catch"] = proc.GoFunction{Func: cmdCatch, TypeIn: types.Generic, TypeOut: types.Generic}
	proc.GoFunctions["!catch"] = proc.GoFunction{Func: cmdCatch, TypeIn: types.Generic, TypeOut: types.Generic}
}

func cmdTry(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Generic)

	block, err := p.Parameters.Block(0)
	if err != nil {
		return err
	}

	p.ExitNum, err = lang.ProcessNewBlock(block, nil, p.Stdout, p.Stderr, p.Name)
	if err != nil {
		return err
	}

	return
}

func cmdCatch(p *proc.Process) error {
	p.Stdout.SetDataType(types.Generic)

	block, err := p.Parameters.Block(0)
	if err != nil {
		return err
	}

	_, err = io.Copy(p.Stdout, p.Stdin)
	if err != nil {
		return err
	}

	p.ExitNum = p.Previous.ExitNum

	if p.Previous.ExitNum != 0 && !p.IsNot {
		_, err = lang.ProcessNewBlock(block, nil, p.Stdout, p.Stderr, types.Null)
		if err != nil {
			return err
		}

	} else if p.Previous.ExitNum == 0 && p.IsNot {
		_, err = lang.ProcessNewBlock(block, nil, p.Stdout, p.Stderr, types.Null)
		if err != nil {
			return err
		}
	}

	return nil
}
