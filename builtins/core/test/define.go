package cmdtest

import (
	"regexp"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
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
	rx, err := regexp.Compile(args.OutRegexp)
	if err != nil {
		return err
	}
	stdout := &lang.TestChecks{
		Regexp:   rx,
		Block:    []rune(args.OutBlock),
		RunBlock: runBlock,
	}

	// stderr
	rx, err = regexp.Compile(args.ErrRegexp)
	if err != nil {
		return err
	}
	stderr := &lang.TestChecks{
		Regexp:   rx,
		Block:    []rune(args.ErrBlock),
		RunBlock: runBlock,
	}

	err = p.Tests.Define(name, stdout, stderr, args.ExitNum)
	return err
}

func runBlock(p *lang.Process, block []rune, expected []byte) ([]byte, []byte, error) {
	fork := p.Fork(lang.F_CREATE_STDIN | lang.F_CREATE_STDERR | lang.F_CREATE_STDOUT)

	fork.Stdin.SetDataType(types.Generic)
	_, err := fork.Stdin.Write(expected)
	if err != nil {
		return nil, nil, err
	}

	_, err = fork.Execute(block)
	if err != nil {
		return nil, nil, err
	}

	stdout, err := fork.Stdout.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	stderr, err := fork.Stderr.ReadAll()
	if err != nil {
		return utils.CrLfTrim(stdout), nil, err
	}

	return utils.CrLfTrim(stdout), utils.CrLfTrim(stderr), nil
}
