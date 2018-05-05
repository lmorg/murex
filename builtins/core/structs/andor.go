package structs

import (
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

func cmdAnd(p *proc.Process) error { return cmdAndOr(p, true) }
func cmdOr(p *proc.Process) error  { return cmdAndOr(p, false) }

func cmdAndOr(p *proc.Process, isAnd bool) error {
	p.Stdout.SetDataType(types.Boolean)

	for i := 0; i < p.Parameters.Len(); i++ {
		block, err := p.Parameters.Block(i)
		if err != nil {
			return err
		}

		stdout := streams.NewStdin()
		i, err := lang.RunBlockExistingConfigSpace(block, nil, stdout, nil, p)
		if err != nil {
			return err
		}

		b, err := stdout.ReadAll()
		if err != nil {
			return err
		}
		conditional := types.IsTrue(b, i)

		if isAnd {
			// --- and ---
			if (!conditional && !p.IsNot) || (conditional && p.IsNot) {
				_, err = p.Stdout.Write(types.FalseByte)
				return err
			}
		} else {
			// --- or ---
			if (conditional && !p.IsNot) || (!conditional && p.IsNot) {
				_, err = p.Stdout.Write(types.TrueByte)
				return err
			}
		}
	}

	if isAnd {
		// --- and ---
		_, err := p.Stdout.Write(types.TrueByte)
		return err
	}

	// --- or ---
	_, err := p.Stdout.Write(types.FalseByte)
	return err
}
