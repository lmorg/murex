package shelltest

import (
	"regexp"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	proc.GoFunctions["test"] = cmdTest

	defaults.AppendProfile(`
		autocomplete set test { [{
			"Flags": [
				"on",
				"off",
				"auto-report",
				"!auto-report"
			]
		}] }
    `)
}

type testArgs struct {
	OutBlock  string
	OutRegexp string
	ErrBlock  string
	ErrRegexp string
	ExitNum   int
}

func cmdTest(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	if p.Parameters.Len() == 1 {
		return testConfig(p)
	}

	enabled, err := p.Config.Get("test", "enabled", types.Boolean)
	if err != nil || !enabled.(bool) {
		return err
	}

	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	b, err := p.Parameters.Byte(1)
	if err != nil {
		return err
	}

	var args testArgs
	err = json.UnmarshalMurex(b, &args)
	if err != nil {
		return err
	}

	// stdout
	rx, err := regexp.Compile(args.OutRegexp)
	if err != nil {
		return err
	}
	stdout := &proc.TestChecks{
		Regexp:   rx,
		Block:    []rune(args.OutBlock),
		RunBlock: runTest,
	}

	// stderr
	rx, err = regexp.Compile(args.ErrRegexp)
	if err != nil {
		return err
	}
	stderr := &proc.TestChecks{
		Regexp:   rx,
		Block:    []rune(args.ErrBlock),
		RunBlock: runTest,
	}

	err = p.Tests.Define(name, stdout, stderr, args.ExitNum)
	return err
}

func runTest(p *proc.Process, block []rune) ([]byte, error) {
	stdout := streams.NewStdin()
	_, err := lang.RunBlockExistingConfigSpace(block, nil, stdout, proc.ShellProcess.Stderr, p)
	if err != nil {
		return nil, err
	}

	b, err := stdout.ReadAll()
	if err != nil {
		return nil, err
	}
	return utils.CrLfTrim(b), nil
}

func testConfig(p *proc.Process) (err error) {
	option, _ := p.Parameters.String(0)

	switch option {
	case "enable":
		err = p.Config.Set("test", "enabled", true)

	case "!enable", "disable":
		err = p.Config.Set("test", "enabled", false)

	case "auto-report":
		err = p.Config.Set("test", "auto-report", true)

	case "!auto-report":
		err = p.Config.Set("test", "auto-report", false)

	default:
		err = p.Config.Set("test", "enabled", types.IsTrue([]byte(option), 0))
	}

	return
}
