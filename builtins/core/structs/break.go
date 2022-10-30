package structs

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("break", cmdBreak, types.Null)
}

func cmdBreak(p *lang.Process) error {
	name, _ := p.Parameters.String(0)
	if name == "" {
		p.Done()
		p.Parent.Done()
		return fmt.Errorf(
			"missing parameter. Stopping execution of `%s` as a precaution",
			p.Parent.Name.String(),
		)
	}

	scope := p.Scope.Id
	proc := p.Parent
	for {
		proc.Done()

		if proc.Name.String() == name {
			return nil
		}
		if proc.Id == scope {
			return fmt.Errorf(
				"no block found named `%s` within the scope of `%s`",
				name, p.Scope.Name.String(),
			)
		}

		proc = proc.Parent
	}
}
