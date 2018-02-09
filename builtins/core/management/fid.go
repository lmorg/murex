package management

import (
	"encoding/json"
	"fmt"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["fid-list"] = cmdFidList
	proc.GoFunctions["fid-kill"] = cmdFidKill
	proc.GoFunctions["fid-killall"] = cmdKillAll
}

func cmdFidList(p *proc.Process) error {
	yn := func(state bool) (s string) {
		if state {
			return "yes"
		}
		return "no"
	}

	p.Stdout.SetDataType(types.Generic)
	p.Stdout.Writeln([]byte(fmt.Sprintf("%7s  %7s  %7s  %-12s  %-8s  %-3s  %-10s  %-10s  %-10s  %s",
		"FID", "Parent", "Scope", "State", "Run Mode", "BG", "Out Pipe", "Err Pipe", "Command", "Parameters")))

	procs := proc.GlobalFIDs.ListAll()
	for i := range procs {
		params := procs[i].Parameters.StringAll()
		if len(params) == 0 && len(procs[i].Parameters.Tokens) > 1 {
			b, _ := json.Marshal(procs[i].Parameters.Tokens)
			params = "Unparsed: " + string(b)
		}
		s := fmt.Sprintf("%7d  %7d  %7d  %-12s  %-8s  %-3s  %-10s  %-10s  %-10s  %s",
			procs[i].Id,
			procs[i].Parent.Id,
			procs[i].Scope.Id,
			procs[i].State,
			procs[i].RunMode,
			yn(procs[i].IsBackground),
			procs[i].NamedPipeOut,
			procs[i].NamedPipeErr,
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
		err = fmt.Errorf("fid `%d` cannot be killed. `Kill` method == `nil`.", fid)
	}

	return err
}

func cmdKillAll(*proc.Process) (err error) {
	fids := proc.GlobalFIDs.ListAll()
	for _, f := range fids {
		if f.Kill != nil {
			go f.Kill()
		}
	}

	return err
}
