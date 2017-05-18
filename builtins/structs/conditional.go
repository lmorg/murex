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
}

func cmdIf(p *proc.Process) (err error) {
	var ifBlock, thenBlock, elseBlock []rune

	switch {
	case p.Parameters.Len() == 1 && p.IsMethod:
		// "if" taken from stdin, "then" from 1st parameter.
		thenBlock, err = p.Parameters.Block(0)
		if err != nil {
			return err
		}

	case p.Parameters.Len() == 2 && p.IsMethod:
		// "if" taken from stdin, "then" and "else" from 1st and 2nd parameter.
		thenBlock, err = p.Parameters.Block(0)
		if err != nil {
			return err
		}

		elseBlock, err = p.Parameters.Block(1)
		if err != nil {
			return err
		}

	case p.Parameters.Len() == 2 && !p.IsMethod:
		// "if" taken from 1st parameter, "then" from 2nd parameter.
		ifBlock, err = p.Parameters.Block(0)
		if err != nil {
			return err
		}

		thenBlock, err = p.Parameters.Block(1)
		if err != nil {
			return err
		}

	case p.Parameters.Len() == 3 && !p.IsMethod:
		// "if" taken from 1st parameter, "then" from 2nd, "else" from 3rd.
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
		if !p.IsNot {
			return errors.New(`Not a valid if statement. Usage:
  $conditional -> if: { $then }            # conditional result read from stdin or previous process exit number
  $conditional -> if: { $then } { $else }  # conditional result read from stdin or previous process exit number
  if: { $conditional } { $then }           # if / then
  if: { $conditional } { $then } { $else } # if / then / else
`)
		} else {
			return errors.New(`Not a valid if statement. Usage:
  $conditional -> !if: { $else }            # conditional result read from stdin or previous process exit number
  $conditional -> !if: { $else } { $then }  # conditional result read from stdin or previous process exit number
  !if: { $conditional } { $else }           # if / then
  !if: { $conditional } { $else } { $then } # if / then / else
`)
		}
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

	if (conditional && !p.IsNot) || (!conditional && p.IsNot) {
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
