package processes

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/state"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.DefineFunction("fid-list-new", cmdFidListNew, types.JsonLines)

	/*defaults.AppendProfile(`
	autocomplete set fid-list { [{
		"DynamicDesc": ({ fid-list --help })
	}] }

	alias jobs=fid-list --jobs
	method define jobs {
		"in": "null",
		"out": "*"
	}
	config eval shell safe-commands { -> append jobs }`)*/
}

const (
	fCsv           = "--csv"
	fJsonL         = "--jsonl"
	fTty           = "--tty"
	fJobs          = "--jobs"
	fStopped       = "--stopped"
	fBackground    = "--background"
	fExcChildrenOf = "--exc-children-of"
	fIncChildrenOf = "--inc-children-of"
	fHelp          = "--help"
)

var argsFidList = &parameters.Arguments{
	AllowAdditional:     false,
	IgnoreInvalidFlags:  false,
	StrictFlagPlacement: false,
	Flags: map[string]string{
		fCsv:   types.Boolean,
		fJsonL: types.Boolean,
		fTty:   types.Boolean,

		fJobs:       types.Boolean,
		fStopped:    types.Boolean,
		fBackground: types.Boolean,

		fIncChildrenOf: types.Integer,
		fExcChildrenOf: types.Integer,

		fHelp: types.Boolean,
	},
}

func cmdFidListNew(p *lang.Process) error {
	flags, _, err := p.Parameters.ParseFlags(argsFidList)
	if err != nil {
		return err
	}

	switch {
	case flags.GetValue(fCsv).Boolean():
		return cmdFidListCSV(p)

	case flags.GetValue(fJsonL).Boolean():
		return cmdFidListPipe(p)

	case flags.GetValue(fTty).Boolean():
		return cmdFidListTTY(p)

	case flags.GetValue(fJobs).Boolean():
		return cmdJobs(p)

	case flags.GetValue(fStopped).Boolean():
		return cmdJobsStopped(p)

	case flags.GetValue(fBackground).Boolean():
		return cmdJobsBackground(p)

	case flags.GetValue(fHelp).Boolean():
		return cmdFidListHelp(p)

	default:
		if p.Stdout.IsTTY() {
			return cmdFidListTTY(p)
		}
		return cmdFidListPipe(p)
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

func cmdFidListNewTTY(p *lang.Process, procs []*lang.Process) error {
	p.Stdout.SetDataType(types.Generic)
	p.Stdout.Writeln([]byte(fmt.Sprintf("%7s  %7s  %7s  %-12s  %-8s  %-3s  %-10s  %-10s  %-10s  %s",
		"FID", "Parent", "Scope", "State", "Run Mode", "BG", "Out Pipe", "Err Pipe", "Command", "Parameters")))

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

func cmdFidListNewCSV(p *lang.Process, procs []*lang.Process) error {
	p.Stdout.SetDataType("csv")
	p.Stdout.Writeln([]byte(fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s",
		"FID", "Parent", "Scope", "State", "Run Mode", "BG", "Out Pipe", "Err Pipe", "Command", "Parameters")))

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

func cmdFidListNewPipe(p *lang.Process, procs []*lang.Process) error {
	p.Stdout.SetDataType(types.JsonLines)

	b, err := lang.MarshalData(p, types.Json, []any{
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

	for _, process := range procs {
		b, err = lang.MarshalData(p, types.Json, []any{
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

func cmdJobsNewStopped(p *lang.Process, procs []*lang.Process) error {
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

func cmdJobsNewBackground(p *lang.Process, procs []*lang.Process) error {
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
