package processes

import (
	"fmt"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.GoFunctions["fid-list"] = cmdFidList

	defaults.AppendProfile(`
	autocomplete: set fid-list { [{
		"DynamicDesc": ({ fid-list --help })
	}] }

	alias: jobs=fid-list --jobs
	config: eval shell safe-commands { -> append jobs }
`)
}

func yn(state bool) (s string) {
	if state {
		return "yes"
	}
	return "no"
}

func getParams(p *lang.Process) string {
	params := p.Parameters.StringAll()
	if len(params) == 0 && len(p.Parameters.Tokens) > 1 {
		newParams := parameters.Parameters{
			Tokens: p.Parameters.Tokens,
		}
		lang.ParseParameters(p, &newParams)
		params = "(subject to change) " + newParams.StringAll()
	}
	return params
}

func cmdFidList(p *lang.Process) error {
	flag, _ := p.Parameters.String(0)
	switch flag {
	case "":
		if p.Stdout.IsTTY() {
			return cmdFidListTTY(p)
		}
		return cmdFidListPipe(p)

	case "--csv":
		return cmdFidListCSV(p)

	case "--jsonl":
		return cmdFidListPipe(p)

	case "--tty":
		return cmdFidListTTY(p)

	case "--jobs":
		return cmdJobs(p)

	case "--stopped":
		return cmdJobsStopped(p)

	case "--background":
		return cmdJobsBackground(p)

	case "--help":
		fallthrough

	default:
		return cmdFidListHelp(p)
	}
}

func cmdFidListHelp(p *lang.Process) error {
	flags := map[string]string{
		"--csv":        "Outputs as CSV table",
		"--jsonl":      "Outputs as a jsonlines (a greppable array of JSON objects). This is the default mode when `fid-list` is piped",
		"--tty":        "Outputs as a human readable table. This is the default mode when outputting to a TTY",
		"--stopped":    "JSON map of all stopped processes running under murex",
		"--background": "JSON map of all background processes running under murex",
		"--jobs":       "List stopped or background processes (similar to POSIX jobs)",
		"--help":       "Displays a list of parameters",
	}
	p.Stdout.SetDataType(types.Json)
	b, err := json.Marshal(flags, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func cmdFidListTTY(p *lang.Process) error {
	p.Stdout.SetDataType(types.Generic)
	p.Stdout.Writeln([]byte(fmt.Sprintf("%7s  %7s  %7s  %-12s  %-8s  %-3s  %-10s  %-10s  %-10s  %s",
		"FID", "Parent", "Scope", "State", "Run Mode", "BG", "Out Pipe", "Err Pipe", "Command", "Parameters")))

	procs := lang.GlobalFIDs.ListAll()
	for i := range procs {
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
			getParams(&procs[i]),
		)
		_, err := p.Stdout.Writeln([]byte(s))
		if err != nil {
			return err
		}
	}
	return nil
}

func cmdFidListCSV(p *lang.Process) error {
	p.Stdout.SetDataType("csv")
	p.Stdout.Writeln([]byte(fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s",
		"FID", "Parent", "Scope", "State", "Run Mode", "BG", "Out Pipe", "Err Pipe", "Command", "Parameters")))

	procs := lang.GlobalFIDs.ListAll()
	for i := range procs {
		s := fmt.Sprintf("%d,%d,%d,%s,%s,%s,%s,%s,%s,%s",
			procs[i].Id,
			procs[i].Parent.Id,
			procs[i].Scope.Id,
			procs[i].State,
			procs[i].RunMode,
			yn(procs[i].IsBackground),
			procs[i].NamedPipeOut,
			procs[i].NamedPipeErr,
			procs[i].Name,
			getParams(&procs[i]),
		)
		_, err := p.Stdout.Writeln([]byte(s))
		if err != nil {
			return err
		}
	}
	return nil
}

type fidList struct {
	FID        uint32
	Parent     uint32
	Scope      uint32
	State      string
	RunMode    string `json:"Run Mode"`
	BG         bool   `json:"Background"`
	OutPipe    string `json:"Out Pipe"`
	ErrPipe    string `json:"Err Pipe"`
	Command    string
	Parameters string
}

func cmdFidListPipe(p *lang.Process) error {
	var fids []interface{}
	p.Stdout.SetDataType(types.JsonLines)

	procs := lang.GlobalFIDs.ListAll()
	for i := range procs {
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
			Parameters: getParams(&procs[i]),
		})
	}

	b, err := lang.MarshalData(p, types.JsonLines, fids)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func cmdJobsStopped(p *lang.Process) error {
	procs := lang.GlobalFIDs.ListAll()
	m := make(map[uint32]string)

	for i := range procs {
		if procs[i].State != state.Stopped {
			continue
		}
		m[procs[i].Id] = procs[i].Name + " " + getParams(&procs[i])
	}

	b, err := lang.MarshalData(p, types.Json, m)
	if err != nil {
		return err
	}

	p.Stdout.SetDataType(types.Json)
	_, err = p.Stdout.Write(b)
	return err
}

func cmdJobsBackground(p *lang.Process) error {
	procs := lang.GlobalFIDs.ListAll()
	m := make(map[uint32]string)

	for i := range procs {
		if !procs[i].IsBackground {
			continue
		}
		m[procs[i].Id] = procs[i].Name + " " + getParams(&procs[i])
	}

	b, err := lang.MarshalData(p, types.Json, m)
	if err != nil {
		return err
	}

	p.Stdout.SetDataType(types.Json)
	_, err = p.Stdout.Write(b)
	return err
}
