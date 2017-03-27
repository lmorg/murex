package structs

import (
	"errors"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["if"] = proc.GoFunction{Func: cmdIf, TypeIn: types.Generic, TypeOut: types.Generic}
	proc.GoFunctions["!if"] = proc.GoFunction{Func: cmdIf, TypeIn: types.Generic, TypeOut: types.Generic}
	proc.GoFunctions["else"] = proc.GoFunction{Func: cmdElse, TypeIn: types.Generic, TypeOut: types.Generic}
}

// `else` is essentially an alias for `if!`
func cmdElse(p *proc.Process) error {
	p.Not = true
	return cmdIf(p)
}

func cmdIf(p *proc.Process) (err error) {
	var ifBlock, thenBlock, elseBlock []rune

	switch p.Parameters.Len() {
	case 1:
		// "if" taken from stdin, "then" from second parameter.
		thenBlock, err = p.Parameters.Block(0)
		if err != nil {
			return err
		}

	case 2:
		// "if" taken from first parameter, "then" from second parameter.
		ifBlock, err = p.Parameters.Block(0)
		if err != nil {
			return err
		}

		thenBlock, err = p.Parameters.Block(1)
		if err != nil {
			return err
		}

	case 3:
		// "if" taken from first parameter, "then" from second, "else" from third.
		ifBlock, err = p.Parameters.Block(0)
		if err != nil {
			return err
		}

		thenBlock, err = p.Parameters.Block(1)
		if err != nil {
			return err
		}

		elseBlock, err = p.Parameters.Block(2)
		if err != nil {
			return err
		}

	default:
		// Error
		return errors.New(`Not a valid if statement. Usage:
  $conditional -> if: { $then }            # conditional result read from stdin or previous process exit number
  if: { $conditional } { $then }           # if / then
  if: { $conditional } { $then } { $else } # if / then / else
`)
	}

	var conditional bool
	if len(ifBlock) != 0 {
		// --- IF ---
		stdout := streams.NewStdin()
		i, err := lang.ProcessNewBlock(ifBlock, nil, stdout, nil, types.Null)
		if err != nil {
			return err
		}
		stdout.Close()
		b := stdout.ReadAll()
		conditional = types.IsTrue(b, i)

	} else {
		// --- IF ---
		b := p.Stdin.ReadAll()
		conditional = types.IsTrue(b, p.Previous.ExitNum)
	}

	if (conditional && !p.Not) || (!conditional && p.Not) {
		// --- THEN ---
		_, err = lang.ProcessNewBlock(thenBlock, nil, p.Stdout, p.Stderr, types.Null)
		if err != nil {
			return
		}

	} else {
		// --- ELSE ---
		if len(elseBlock) != 0 {
			_, err = lang.ProcessNewBlock(elseBlock, nil, p.Stdout, p.Stderr, types.Null)
			if err != nil {
				return
			}
		} else {
			p.ExitNum = 1
		}
	}

	return
}
