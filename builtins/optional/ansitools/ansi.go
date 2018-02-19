// +build ignore

package ansitools

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils/ansi"
)

func init() {
	proc.GoFunctions["ansi"] = cmdAnsi
}

func cmdAnsi(p *proc.Process) error {
	p.Stdout.SetDataType(types.String)

	flags, _, err := p.Parameters.ParseFlags(
		&parameters.Arguments{
			Flags: map[string]string{
				"--foreground": "-f",
				"-f":           types.String,
				"--background": "-b",
				"-b":           types.String,
			},
			AllowAdditional: false,
		},
	)

	for f := range flags {
		switch f {
		case -f:
		default:
			return errors.New("Flag missing condition")
		}
	}

	_, err = p.Stdout.Write(s)
	return err
}
