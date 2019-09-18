package cmdtest

import (
	"regexp"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func testDefine(p *lang.Process) error {
	enabled, err := p.Config.Get("test", "enabled", types.Boolean)
	if err != nil || !enabled.(bool) {
		return err
	}

	name, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	b, err := p.Parameters.Byte(2)
	if err != nil {
		return err
	}

	var args testArgs
	err = json.UnmarshalMurex(b, &args)
	if err != nil {
		return err
	}

	// stdout
	rx, err := regexp.Compile(args.StdoutRegex)
	if err != nil {
		return err
	}
	stdout := &lang.TestChecks{
		Regexp: rx,
		Block:  []rune(args.StdoutBlock),
	}

	// stderr
	rx, err = regexp.Compile(args.StderrRegex)
	if err != nil {
		return err
	}
	stderr := &lang.TestChecks{
		Regexp: rx,
		Block:  []rune(args.StderrBlock),
	}

	err = p.Tests.Define(name, stdout, stderr, args.ExitNum)
	return err
}
