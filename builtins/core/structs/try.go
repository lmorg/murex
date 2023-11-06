package structs

import (
	"errors"
	"fmt"
	"io"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/runmode"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("runmode", cmdRunmode, types.Null)

	lang.DefineFunction("unsafe", cmdUnsafe, types.Any)
	lang.DefineFunction("try", cmdTry, types.Any)
	lang.DefineFunction("trypipe", cmdTryPipe, types.Any)
	lang.DefineFunction("tryerr", cmdTryErr, types.Any)
	lang.DefineFunction("trypipeerr", cmdTryPipeErr, types.Any)

	lang.DefineFunction("catch", cmdCatch, types.Any)
	lang.DefineFunction("!catch", cmdCatch, types.Any)
}

func cmdRunmode(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)
	return errors.New("`runmode` should only be used as the first statement in a block")
}

func cmdTry(p *lang.Process) error        { return tryModes(p, runmode.BlockTry) }
func cmdTryPipe(p *lang.Process) error    { return tryModes(p, runmode.BlockTryPipe) }
func cmdTryErr(p *lang.Process) error     { return tryModes(p, runmode.BlockTryErr) }
func cmdTryPipeErr(p *lang.Process) error { return tryModes(p, runmode.BlockTryPipeErr) }

func tryModes(p *lang.Process, runMode runmode.RunMode) (err error) {
	p.Stdout.SetDataType(types.Null)

	block, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	r := []rune(block)
	if types.IsBlockRune(r) {
		p.RunMode = runMode
		p.ExitNum, err = p.Fork(lang.F_PARENT_VARTABLE).Execute(r)
		return
	}

	return fmt.Errorf("unexpected parameter '%s', expecting a code block inside curly braces", block)
}

func cmdUnsafe(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	block, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	r := []rune(block)
	if types.IsBlockRune(r) {
		p.RunMode = runmode.BlockUnsafe
		p.ExitNum, err = p.Fork(lang.F_PARENT_VARTABLE).Execute(r)
		if err != nil {
			p.Stderr.Writeln([]byte(err.Error()))
		}
		return nil
	}

	return fmt.Errorf("unexpected parameter '%s', expecting a code block inside curly braces", block)
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
		_, err = p.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN).Execute(block)
		if err != nil {
			return err
		}

	} else if p.Previous.ExitNum == 0 && p.IsNot {
		_, err = p.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN).Execute(block)
		if err != nil {
			return err
		}
	}

	return nil
}
