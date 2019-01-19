package management

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.GoFunctions["fid-list"] = cmdFidList
	lang.GoFunctions["fid-kill"] = cmdFidKill
	lang.GoFunctions["fid-killall"] = cmdKillAll
}

func cmdFidList(p *lang.Process) error {
	if p.Stdout.IsTTY() {
		return cmdFidListTTY(p)
	}
	return cmdFidListPipe(p)
}

func cmdFidListTTY(p *lang.Process) error {
	yn := func(state bool) (s string) {
		if state {
			return "yes"
		}
		return "no"
	}

	p.Stdout.SetDataType(types.Generic)
	p.Stdout.Writeln([]byte(fmt.Sprintf("%7s  %7s  %7s  %-12s  %-8s  %-3s  %-10s  %-10s  %-10s  %s",
		"FID", "Parent", "Scope", "State", "Run Mode", "BG", "Out Pipe", "Err Pipe", "Command", "Parameters")))

	procs := lang.GlobalFIDs.ListAll()
	for i := range procs {
		params := procs[i].Parameters.StringAll()
		if len(params) == 0 && len(procs[i].Parameters.Tokens) > 1 {
			b, _ := json.Marshal(procs[i].Parameters.Tokens, false)
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
		_, err := p.Stdout.Writeln([]byte(s))
		if err != nil {
			return err
		}
	}
	return nil
}

type fidList struct {
	FID        int
	Parent     int
	Scope      int
	State      string
	RunMode    string `json:"Run Mode"`
	BG         bool
	OutPipe    string `json:"Out Pipe"`
	ErrPipe    string `json:"Err Pipe"`
	Command    string
	Parameters string
}

func cmdFidListPipe(p *lang.Process) error {
	var fids []fidList

	p.Stdout.SetDataType(types.Json)

	procs := lang.GlobalFIDs.ListAll()
	for i := range procs {
		params := procs[i].Parameters.StringAll()
		if len(params) == 0 && len(procs[i].Parameters.Tokens) > 1 {
			b, _ := json.Marshal(procs[i].Parameters.Tokens, false)
			params = "Unparsed: " + string(b)
		}
		fids = append(fids, fidList{
			FID:        procs[i].Id,
			Parent:     procs[i].Parent.Id,
			Scope:      procs[i].Scope.Id,
			State:      procs[i].State.String(),
			RunMode:    procs[i].RunMode.String(),
			BG:         procs[i].IsBackground,
			OutPipe:    procs[i].NamedPipeOut,
			ErrPipe:    procs[i].NamedPipeErr,
			Command:    procs[i].Name,
			Parameters: params,
		})
	}

	b, err := json.Marshal(fids, false)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func cmdFidKill(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Null)

	for i := 0; i < p.Parameters.Len(); i++ {
		fid, err := p.Parameters.Int(i)
		if err != nil {
			return err
		}

		process, err := lang.GlobalFIDs.Proc(fid)
		if err != nil {
			return err
		}

		if process.Kill != nil {
			process.Kill()
		} else {
			err = fmt.Errorf("fid `%d` cannot be killed. `Kill` method == `nil`", fid)
		}
	}

	return err
}

func cmdKillAll(*lang.Process) error {
	fids := lang.GlobalFIDs.ListAll()
	for _, p := range fids {
		if p.Kill != nil /*&& !p.HasTerminated()*/ {
			procName := p.Name
			procParam, _ := p.Parameters.String(0)
			if p.Name == "exec" {
				procName = procParam
				procParam, _ = p.Parameters.String(1)
			}
			if len(procParam) > 10 {
				procParam = procParam[:10]
			}
			lang.ShellProcess.Stderr.Write([]byte(fmt.Sprintf("!!! Sending kill signal to fid %d: %s %s !!!\n", p.Id, procName, procParam)))
			p.Kill()
		}
	}

	return nil
}
