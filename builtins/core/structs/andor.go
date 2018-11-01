package structs

import (
	"errors"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["and"] = cmdAnd
	proc.GoFunctions["!and"] = cmdAnd
	proc.GoFunctions["or"] = cmdOr
	proc.GoFunctions["!or"] = cmdOr
}

const errCancelled = "User has cancelled processing mid-way through the execution of this control flow structure."

func cmdAnd(p *proc.Process) error { return cmdAndOr(p, true) }
func cmdOr(p *proc.Process) error  { return cmdAndOr(p, false) }

func cmdAndOr(p *proc.Process, isAnd bool) error {
	p.Stdout.SetDataType(types.Boolean)

	for i := 0; i < p.Parameters.Len(); i++ {
		if p.HasCancelled() {
			return errors.New(errCancelled)
		}

		block, err := p.Parameters.Block(i)
		if err != nil {
			return err
		}

		stdout := streams.NewStdin()
		i, err := lang.RunBlockExistingConfigSpace(block, nil, stdout, nil, p)
		if err != nil {
			return err
		}

		if p.HasCancelled() {
			return errors.New(errCancelled)
		}

		b, err := stdout.ReadAll()
		if err != nil {
			return err
		}
		conditional := types.IsTrue(b, i)

		if isAnd {
			// --- and ---
			if (!conditional && !p.IsNot) || (conditional && p.IsNot) {
				p.ExitNum = 1
				return nil
			}
		} else {
			// --- or ---
			if (conditional && !p.IsNot) || (!conditional && p.IsNot) {
				p.ExitNum = -1
				return nil
			}
		}
	}

	if isAnd {
		// --- and ---
		p.ExitNum = -1
		return nil
	}

	// --- or ---
	p.ExitNum = 1
	return nil
}
