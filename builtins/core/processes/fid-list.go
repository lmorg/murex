package processes

import (
	"fmt"
	"strings"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/state"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.DefineFunction("fid-list", cmdFidList, types.JsonLines)

	defaults.AppendProfile(`
		autocomplete set fid-list { [{
			"DynamicDesc": ({ fid-list --help })
		}] }

		alias jobs=fid-list --jobs
		method define jobs {
			"in": "null",
			"out": "*"
		}
		config eval shell safe-commands { -> append jobs }`)
}

func getParams(p *lang.Process) string {
	s := string(p.Parameters.GetRaw())
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.TrimSpace(s)
	return s
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
			process.State.String(),
			process.RunMode.String(),
			process.Background.String(),
			process.NamedPipeOut,
			process.NamedPipeErr,
			process.Name.String(),
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
			process.State.String(),
			process.RunMode.String(),
			process.Background.String(),
			process.NamedPipeOut,
			process.NamedPipeErr,
			process.Name.String(),
			getParams(process),
		)
		_, err := p.Stdout.Writeln([]byte(s))
		if err != nil {
			return err
		}
	}
	return nil
}

func cmdFidListPipe(p *lang.Process) error {
	p.Stdout.SetDataType(types.JsonLines)

	b, err := lang.MarshalData(p, types.Json, []interface{}{
		"FID",
		"Parent",
		"Scope",
		"State",
		"RunMode",
		"BG",
		"OutPipe",
		"ErrPipe",
		"Command",
		"Parameters",
	})
	if err != nil {
		return err
	}
	_, err = p.Stdout.Writeln(b)
	if err != nil {
		return err
	}

	procs := lang.GlobalFIDs.ListAll()
	for _, process := range procs {
		b, err = lang.MarshalData(p, types.Json, []interface{}{
			process.Id,
			process.Parent.Id,
			process.Scope.Id,
			process.State.String(),
			process.RunMode.String(),
			process.Background.Get(),
			process.NamedPipeOut,
			process.NamedPipeErr,
			process.Name.String(),
			getParams(process),
		})
		if err != nil {
			return err
		}
		_, err = p.Stdout.Writeln(b)
		if err != nil {
			return err
		}
	}

	return nil
}

func cmdJobsStopped(p *lang.Process) error {
	procs := lang.GlobalFIDs.ListAll()
	m := make(map[uint32]string)

	for _, process := range procs {
		if process.State.Get() != state.Stopped {
			continue
		}
		m[process.Id] = process.Name.String() + " " + getParams(process)
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
		if !process.Background.Get() {
			continue
		}
		m[process.Id] = process.Name.String() + " " + getParams(process)
	}

	b, err := lang.MarshalData(p, types.Json, m)
	if err != nil {
		return err
	}

	p.Stdout.SetDataType(types.Json)
	_, err = p.Stdout.Write(b)
	return err
}
