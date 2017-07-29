package management

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["fid-list"] = proc.GoFunction{Func: cmdFidList, TypeIn: types.Null, TypeOut: types.String}
	proc.GoFunctions["fid-kill"] = proc.GoFunction{Func: cmdFidKill, TypeIn: types.Null, TypeOut: types.String}
}

func cmdFidList(p *proc.Process) error {
	yn := func(state bool) (s string) {
		if state {
			return "yes"
		}
		return "no"
	}

	p.Stdout.SetDataType(types.Generic)
	p.Stdout.Writeln([]byte(fmt.Sprintf("%7s  %7s  %-12s  %-3s  %-10s  %s", "FID", "Parent", "State", "BG", "Command", "Parameters")))

	procs := proc.GlobalFIDs.ListAll()
	for i := range procs {
		params := procs[i].Parameters.StringAll()
		if len(params) == 0 && len(procs[i].Parameters.Tokens) > 1 {
			b, _ := json.Marshal(procs[i].Parameters.Tokens)
			params = "Unparsed: " + string(b)
		}
		s := fmt.Sprintf("%7d  %7d  %-12s  %-3s  %-10s  %s",
			procs[i].Id,
			procs[i].Parent.Id,
			procs[i].State,
			yn(procs[i].IsBackground),
			procs[i].Name,
			params,
		)
		p.Stdout.Writeln([]byte(s))
	}
	return nil
}

func cmdFidKill(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Null)

	fid, err := p.Parameters.Int(0)
	if err != nil {
		return err
	}

	process, err := proc.GlobalFIDs.Proc(fid)
	if err != nil {
		return err
	}

	if process.Kill != nil {
		process.Kill()
	} else {
		err = errors.New(fmt.Sprintf("fid `%d` cannot be killed. `Kill` method == `nil`.", fid))
	}

	return err
}
