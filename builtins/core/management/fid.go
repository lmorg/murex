package management

import (
	"fmt"
	"strconv"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.GoFunctions["fid-list"] = cmdFidList
	lang.GoFunctions["fid-kill"] = cmdFidKill
	lang.GoFunctions["fid-killall"] = cmdKillAll

	defaults.AppendProfile(`
	autocomplete set fid-list { [{
		"DynamicDesc": ({ fid-list --help })
	}] }
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
	for _, process := range procs {
		s := fmt.Sprintf("%7d  %7d  %7d  %-12s  %-8s  %-3s  %-10s  %-10s  %-10s  %s",
			process.Id,
			process.Parent.Id,
			process.Scope.Id,
			process.State,
			process.RunMode,
			yn(process.IsBackground),
			process.NamedPipeOut,
			process.NamedPipeErr,
			process.Name,
			getParams(process),
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
	for _, process := range procs {
		s := fmt.Sprintf("%d,%d,%d,%s,%s,%s,%s,%s,%s,%s",
			process.Id,
			process.Parent.Id,
			process.Scope.Id,
			process.State,
			process.RunMode,
			yn(process.IsBackground),
			process.NamedPipeOut,
			process.NamedPipeErr,
			process.Name,
			getParams(process),
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
	var fids []fidList

	p.Stdout.SetDataType(types.JsonLines)

	procs := lang.GlobalFIDs.ListAll()
	for _, process := range procs {
		fids = append(fids, fidList{
			FID:        process.Id,
			Parent:     process.Parent.Id,
			Scope:      process.Scope.Id,
			State:      process.State.String(),
			RunMode:    process.RunMode.String(),
			BG:         process.IsBackground,
			OutPipe:    process.NamedPipeOut,
			ErrPipe:    process.NamedPipeErr,
			Command:    process.Name,
			Parameters: getParams(process),
		})
	}

	b, err := json.Marshal(fids, false)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func cmdFidKill(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	for i := 0; i < p.Parameters.Len(); i++ {
		fid, err := p.Parameters.Uint32(i)
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
			return fmt.Errorf("fid `%d` cannot be killed. `Kill` method == `nil`", fid)
		}
	}

	return nil
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

func cmdJobsStopped(p *lang.Process) error {
	procs := lang.GlobalFIDs.ListAll()
	m := make(map[uint32]string)

	for _, process := range procs {
		if process.State != state.Stopped {
			continue
		}
		m[process.Id] = process.Name + " " + getParams(process)
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

	for _, process := range procs {
		if !process.IsBackground {
			continue
		}
		m[process.Id] = process.Name + " " + getParams(process)
	}

	b, err := lang.MarshalData(p, types.Json, m)
	if err != nil {
		return err
	}

	p.Stdout.SetDataType(types.Json)
	_, err = p.Stdout.Write(b)
	return err
}

func cmdJobs(p *lang.Process) error {
	var dt, dtLine string
	if p.Stdout.IsTTY() {
		dt = types.Generic
		dtLine = types.Generic
	} else {
		dt = types.JsonLines
		dtLine = types.Json
	}
	p.Stdout.SetDataType(dt)

	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	procs := lang.GlobalFIDs.ListAll()
	for _, process := range procs {
		if process.IsBackground || process.State == state.Stopped {
			b, err := lang.MarshalData(p, dtLine, []string{
				strconv.Itoa(int(process.Id)),
				process.State.String(),
				yn(process.IsBackground),
				process.Name,
				getParams(process),
			})
			if err != nil {
				return err
			}
			err = aw.Write(b)
			if err != nil {
				return err
			}
		}
	}

	return aw.Close()
}
