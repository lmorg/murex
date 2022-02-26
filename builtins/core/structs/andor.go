package structs

import (
	"errors"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("and", cmdAnd, types.Boolean)
	lang.DefineFunction("!and", cmdAnd, types.Boolean)
	lang.DefineFunction("or", cmdOr, types.Boolean)
	lang.DefineFunction("!or", cmdOr, types.Boolean)
}

const errCancelled = "user has cancelled processing mid-way through the execution of this control flow structure"

func cmdAnd(p *lang.Process) error { return cmdAndOr(p, true) }
func cmdOr(p *lang.Process) error  { return cmdAndOr(p, false) }

func cmdAndOr(p *lang.Process, isAnd bool) error {
	p.Stdout.SetDataType(types.Boolean)

	for i := 0; i < p.Parameters.Len(); i++ {
		if p.HasCancelled() {
			return errors.New(errCancelled)
		}

		block, err := p.Parameters.Block(i)
		if err != nil {
			return err
		}

		//stdout := streams.NewStdin()
		//i, err := lang.RunBlockExistingConfigSpace(block, nil, stdout, nil, p)
		fork := p.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_NO_STDERR)
		i, err := fork.Execute(block)
		if err != nil {
			return err
		}

		if p.HasCancelled() {
			return errors.New(errCancelled)
		}

		b, err := fork.Stdout.ReadAll()
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
