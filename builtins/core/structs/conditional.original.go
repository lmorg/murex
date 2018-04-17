// +build ignore

// As much as I don't normally like to keep old code around as version control
// manages that shit better. I've completely re-written how `if` functions so
// I'm going to keep this file around until I merge into master just in case I
// have introduced some nasty regression bugs

package structs

import (
	"errors"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["if"] = cmdIf
	proc.GoFunctions["!if"] = cmdIf
}

func cmdIf(p *proc.Process) (err error) {
	var ifBlock, thenBlock, elseBlock []rune

	p.Stdout.SetDataType(types.Generic)

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
		} //else {
		return errors.New(`Not a valid if statement. Usage:
  $conditional -> !if: { $else }            # conditional result read from stdin or previous process exit number
  $conditional -> !if: { $else } { $then }  # conditional result read from stdin or previous process exit number
  !if: { $conditional } { $else }           # if / then
  !if: { $conditional } { $else } { $then } # if / then / else
`)
		//}
	}

	var conditional bool
	if len(ifBlock) != 0 {
		// --- IF ---
		stdout := streams.NewStdin()
		stderr := new(streams.Null)
		i, err := lang.RunBlockExistingNamespace(ifBlock, nil, stdout, stderr, p)
		if err != nil {
			return err
		}
		//stdout.Close()
		b, err := stdout.ReadAll()
		if err != nil {
			return err
		}
		conditional = types.IsTrue(b, i)

	} else {
		// --- IF ---
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		conditional = types.IsTrue(b, p.Previous.ExitNum)
	}

	if (conditional && !p.IsNot) || (!conditional && p.IsNot) {
		// --- THEN ---
		_, err = lang.RunBlockExistingNamespace(thenBlock, nil, p.Stdout, p.Stderr, p)
		if err != nil {
			return
		}

	} else {
		// --- ELSE ---
		if len(elseBlock) != 0 {
			_, err = lang.RunBlockExistingNamespace(elseBlock, nil, p.Stdout, p.Stderr, p)
			if err != nil {
				return
			}
		} else {
			p.ExitNum = 1
		}
	}

	return
}
