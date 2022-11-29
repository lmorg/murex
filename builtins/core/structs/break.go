package structs

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("break", cmdBreak, types.Null)
	lang.DefineFunction("continue", cmdContinue, types.Null)
}

func cmdBreak(p *lang.Process) error {
	name, _ := p.Parameters.String(0)
	if name == "" {
		p.Done()
		p.Parent.Done()
		killProcChain(p.Parent)
		return fmt.Errorf(
			"missing parameter. Stopping execution of `%s` as a precaution",
			p.Parent.Name.String(),
		)
	}

	scope := p.Scope.Id
	proc := p.Parent
	for {
		killProcChain(proc)
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

func killProcChain(p *lang.Process) {
	forks := p.Forks.GetForks()
	for _, procs := range forks {
		for i := range *procs {
			(*procs)[i].Done()
		}
	}
}

func cmdContinue(p *lang.Process) error {
	var name string

	name, _ = p.Parameters.String(0)
	if name == "" {
		name = p.Parent.Name.String()
		p.Stderr.Writeln([]byte(fmt.Sprintf(
			"missing parameter. Jumping to `%s` as a precaution", name)))
	}

	scope := p.Scope.Id
	proc := p.Parent
	for {
		if proc.Name.String() == name {
			//proc.IpcContinue <- true
			//proc.Done()
			return nil
		}
		if proc.Id == scope {
			return fmt.Errorf(
				"no block found named `%s` within the scope of `%s`",
				name, p.Scope.Name.String(),
			)
		}
		//go sendIpc(proc, true)
		proc.Done()
		proc = proc.Next
	}
}
