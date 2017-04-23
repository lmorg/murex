package structs

import (
	"bytes"
	"errors"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["foreach"] = proc.GoFunction{Func: cmdForEach, TypeIn: types.Generic, TypeOut: types.Generic}
	proc.GoFunctions["while"] = proc.GoFunction{Func: cmdWhile, TypeIn: types.Null, TypeOut: types.Generic}
}

func cmdForEach(p *proc.Process) (err error) {
	block, err := p.Parameters.Block(1)
	if err != nil {
		return err
	}

	p.Stdin.ReadLineFunc(func(b []byte) {
		b = bytes.TrimSpace(b)
		if len(b) == 0 {
			return
		}

		proc.GlobalVars.Set(p.Parameters[0], string(b), p.Previous.ReturnType)

		stdin := streams.NewStdin()
		stdin.Writeln(b)
		stdin.Close()

		lang.ProcessNewBlock(block, stdin, p.Stdout, p.Stderr, p.Previous.Name)
	})

	return nil
}

func cmdWhile(p *proc.Process) error {
	switch p.Parameters.Len() {
	case 1:
		// Condition is taken from the while loop.
		block, err := p.Parameters.Block(0)
		if err != nil {
			return err
		}

		for {
			i, err := lang.ProcessNewBlock(block, nil, p.Stdout, p.Stderr, types.Null)
			if err != nil || !types.IsTrue([]byte{}, i) {
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
			stdout := streams.NewStdin()
			i, err := lang.ProcessNewBlock(ifBlock, nil, stdout, nil, types.Null)
			if err != nil {
				return err
			}
			stdout.Close()
			b := stdout.ReadAll()
			conditional := types.IsTrue(b, i)

			if !conditional {
				return nil
			}

			lang.ProcessNewBlock(whileBlock, nil, p.Stdout, p.Stderr, types.Null)
		}

	default:
		// Error
		return errors.New("Invalid number of parameters. Please read usage notes.")
	}

	return errors.New("cmdWhile(p *proc.Process) unexpected escaped a switch with default case.")
}
