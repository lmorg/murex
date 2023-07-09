package structs

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("return", cmdReturn, types.Null)
	lang.DefineFunction("break", cmdBreak, types.Null)
	lang.DefineFunction("continue", cmdContinue, types.Null)
}

func cmdReturn(p *lang.Process) error {
	p.ExitNum, _ = p.Parameters.Int(0)
	return breakUpwards(p, p.Scope.Name.String(), p.ExitNum)
}

func cmdBreak(p *lang.Process) error {
	name, _ := p.Parameters.String(0)
	if name == "" {
		p.Done()
		p.Parent.Done()
		killProcChain(p.Parent, 0)
		return fmt.Errorf(
			"missing parameter. Stopping execution of `%s` as a precaution",
			p.Parent.Name.String(),
		)
	}

	return breakUpwards(p, name, 0)
}

func breakUpwards(p *lang.Process, name string, exitNum int) error {
	scope := p.Scope.Id
	proc := p.Parent
	for {
		proc.ExitNum = exitNum
		killProcChain(proc, exitNum)
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

func killProcChain(p *lang.Process, exitNum int) {
	forks := p.Forks.GetForks()
	for _, procs := range forks {
		for i := range *procs {
			(*procs)[i].ExitNum = exitNum
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
