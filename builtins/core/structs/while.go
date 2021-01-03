package structs

import (
	"errors"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["while"] = cmdWhile
	lang.GoFunctions["!while"] = cmdWhile
}

func cmdWhile(p *lang.Process) error {
	p.Stdout.SetDataType(types.Generic)

	switch p.Parameters.Len() {
	case 1:
		// Condition is taken from the while loop.
		block, err := p.Parameters.Block(0)
		if err != nil {
			return err
		}

		for {
			if p.HasCancelled() {
				return errors.New(errCancelled)
			}

			fork := p.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN | lang.F_CREATE_STDOUT)
			i, err := fork.Execute(block)
			if err != nil {
				return err
			}
			b, err := fork.Stdout.ReadAll()
			if err != nil {
				return err
			}

			_, err = p.Stdout.Write(b)
			if err != nil {
				return err
			}

			conditional := types.IsTrue(b, i)

			if (!p.IsNot && !conditional) ||
				(p.IsNot && conditional) {
				return nil
			}

		}

	case 2:
		// Condition is first parameter, while loop is second.
		ifBlock, err := p.Parameters.Block(0)
		if err != nil {
			return err
		}

		whileBlock, err := p.Parameters.Block(1)
		if err != nil {
			return err
		}

		for {
			if p.HasTerminated() {
				return nil
			}

			fork := p.Fork(lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_NO_STDERR)
			i, err := fork.Execute(ifBlock)
			if err != nil {
				return err
			}
			b, err := fork.Stdout.ReadAll()
			if err != nil {
				return err
			}
			conditional := types.IsTrue(b, i)

			if (!p.IsNot && !conditional) ||
				(p.IsNot && conditional) {
				return nil
			}

			fork = p.Fork(lang.F_NO_STDIN)
			err = fork.ExecuteAsRunMode(whileBlock)
			if err != nil {
				return err
			}
		}

	default:
		// Error
		return errors.New("Invalid number of parameters. Please read usage notes")
	}
}
