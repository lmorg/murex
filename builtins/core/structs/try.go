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
	lang.DefineFunction("try", cmdTry, types.Any)
	lang.DefineFunction("trypipe", cmdTryPipe, types.Any)
	lang.DefineFunction("catch", cmdCatch, types.Any)
	lang.DefineFunction("!catch", cmdCatch, types.Any)
}

func cmdRunmode(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)
	return errors.New("`runmode` should only be used as the first statement in a block")
}

func cmdTry(p *lang.Process) error     { return tryModes(p, 0) }
func cmdTryPipe(p *lang.Process) error { return tryModes(p, 1) }

func tryModes(p *lang.Process, adjust runmode.RunMode) (err error) {
	p.Stdout.SetDataType(types.Null)

	scope, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	switch scope {
	/*case "function":
		p.Scope.RunMode = runmode.FunctionTry + adjust
		return nil

	case "module":
		p.Scope.RunMode = runmode.ModuleTry + adjust
		lang.ModuleRunModes[p.FileRef.Source.Module] = runmode.ModuleTry + adjust
		return nil*/

	default:
		r := []rune(scope)
		if types.IsBlockRune(r) {
			p.RunMode = runmode.BlockTry + adjust
			p.ExitNum, err = p.Fork(lang.F_PARENT_VARTABLE).Execute(r)
			return
		}

		//return fmt.Errorf("unexpected parameter '%s'\nExpecting either 'function', 'module' or a code block inside curly braces", scope)
		return fmt.Errorf("unexpected parameter '%s'.\nExpecting either a code block inside curly braces", scope)
	}
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
