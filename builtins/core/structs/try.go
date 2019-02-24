package structs

import (
	"io"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/runmode"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["try"] = cmdTry
	lang.GoFunctions["trypipe"] = cmdTryPipe
	lang.GoFunctions["catch"] = cmdCatch
	lang.GoFunctions["!catch"] = cmdCatch
}

func cmdTry(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Generic)

	block, err := p.Parameters.Block(0)
	if err != nil {
		return err
	}

	p.RunMode = runmode.Try

	//p.ExitNum, err = lang.RunBlockExistingConfigSpace(block, p.Stdin, p.Stdout, p.Stderr, p)
	p.ExitNum, err = p.Fork(lang.F_PARENT_VARTABLE).Execute(block)
	return
}

func cmdTryPipe(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Generic)

	block, err := p.Parameters.Block(0)
	if err != nil {
		return err
	}

	p.RunMode = runmode.TryPipe

	//p.ExitNum, err = lang.RunBlockExistingConfigSpace(block, p.Stdin, p.Stdout, p.Stderr, p)
	p.ExitNum, err = p.Fork(lang.F_PARENT_VARTABLE | lang.F_DEFAULTS).Execute(block)
	return
}

func cmdCatch(p *lang.Process) error {
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
		//_, err = lang.RunBlockExistingConfigSpace(block, nil, p.Stdout, p.Stderr, p)
		_, err = p.Fork(lang.F_NO_STDIN).Execute(block)
		if err != nil {
			return err
		}

	} else if p.Previous.ExitNum == 0 && p.IsNot {
		//_, err = lang.RunBlockExistingConfigSpace(block, nil, p.Stdout, p.Stderr, p)
		_, err = p.Fork(lang.F_NO_STDIN).Execute(block)
		if err != nil {
			return err
		}
	}

	return nil
}
