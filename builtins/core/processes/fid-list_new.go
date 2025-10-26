package processes

import (
	"fmt"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/state"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.DefineFunction("fid-list", cmdFidListNew, types.JsonLines)

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

	options := []lang.OptFidList{lang.FidWithIsFork(false)}
	//options := []lang.OptFidList{}

	if v, ok := flags.GetNullable(fIncChildrenOf); ok {
		options = append(options, lang.FidWithIsChildOf(uint32(v.Integer()), true))
	}

	if v, ok := flags.GetNullable(fExcChildrenOf); ok {
		options = append(options, lang.FidWithIsChildOf(uint32(v.Integer()), false))
	}

	procs := lang.GlobalFIDs.List(options...)

	switch {
	case flags.GetValue(fCsv).Boolean():
		return cmdFidListNewCSV(p, procs)

	case flags.GetValue(fJsonL).Boolean():
		return cmdFidListNewPipe(p, procs)

	case flags.GetValue(fTty).Boolean():
		return cmdFidListNewTTY(p, procs)

	case flags.GetValue(fJobs).Boolean():
		return cmdJobs(p)

	case flags.GetValue(fStopped).Boolean():
		return cmdJobsNewStopped(p, procs)

	case flags.GetValue(fBackground).Boolean():
		return cmdJobsNewBackground(p, procs)

	case flags.GetValue(fHelp).Boolean():
		return cmdFidListNewHelp(p)

	default:
		if p.Stdout.IsTTY() {
			return cmdFidListNewTTY(p, procs)
		}
		return cmdFidListNewPipe(p, procs)
	}
}

func cmdFidListNewHelp(p *lang.Process) error {
	flags := map[string]string{
		fCsv:           "Outputs as CSV table",
		fJsonL:         "Outputs as a jsonlines (a greppable array of JSON objects). This is the default mode when `fid-list` is piped",
		fTty:           "Outputs as a human readable table. This is the default mode when outputting to a TTY",
		fStopped:       "JSON map of all stopped processes running under murex",
		fBackground:    "JSON map of all background processes running under murex",
		fJobs:          "List stopped or background processes (similar to POSIX jobs)",
		fIncChildrenOf: "<int> Include only children of FID",
		fExcChildrenOf: "<int> Exclude all children of FID",
		fHelp:          "Displays a list of parameters",
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
