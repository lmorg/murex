package structs

import (
	"errors"
	"fmt"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("while", cmdWhile, types.Null)
	lang.DefineFunction("!while", cmdWhile, types.Null)
}

const (
	whileConditional int = iota + 1
	whileCheckStdout
)

func cmdWhile(p *lang.Process) error {
	p.Stdout.SetDataType(types.Generic)

	var state, iteration int

	switch p.Parameters.Len() {
	case 2:
		state = whileConditional

	case 1:
		state = whileCheckStdout

	default:
		return fmt.Errorf("invalid usage. Please check docs at https://murex.rocks or `murex-docs %s`", p.Name.String())
	}

	switch state {
	case whileCheckStdout:
		// Condition is taken from the while loop.
		block, err := p.Parameters.Block(0)
		if err != nil {
			return err
		}

		for {
			if p.HasCancelled() {
				return errors.New(errCancelled)
			}

			iteration++
			if !setMetaValues(p, iteration) {
				return fmt.Errorf("cancelled")
			}

			fork := p.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN)
			var stdout stdio.Io
			fork.Stdout, stdout = streams.NewTee(p.Stdout)

			i, err := fork.Execute(block)
			if err != nil {
				return err
			}
			b, err := stdout.ReadAll()
			if err != nil {
				return err
			}

			conditional := types.IsTrue(b, i)

			if (!p.IsNot && !conditional) ||
				(p.IsNot && conditional) {
				return nil
			}

		}

	case whileConditional:
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

			iteration++
			if !setMetaValues(p, iteration) {
				return fmt.Errorf("cancelled")
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
			_, err = fork.Execute(whileBlock)
			if err != nil {
				return err
			}
		}

	default:
		return errors.New("this condition should never be reached. Please file a bug at https://github.com/lmorg/murex/issues")

	}
}
