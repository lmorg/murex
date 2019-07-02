package cmdtest

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"time"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.GoFunctions["test"] = cmdTest
	lang.GoFunctions["!test"] = cmdTestDisable

	defaults.AppendProfile(`
		autocomplete set test { [{
			"Flags": [
				"enable",
				"!enable",
				"auto-report",
				"!auto-report",
				"verbose",
				"!verbose",
				"define",
				"state",
				"run"
			],
			"AllowMultiple": true
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

func cmdTest(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	if p.Parameters.Len() == 0 {
		s, err := p.Config.Get("test", "enabled", types.String)
		if err != nil {
			return err
		}
		_, err = p.Stdout.Write([]byte(s.(string)))
		return err
	}

	option, _ := p.Parameters.String(0)
	switch option {
	case "define":
		return testDefine(p)

	case "state":
		return testState(p)

	case "run":
		return testRun(p)

	default:
		for i := range p.Parameters.StringArray() {
			err := testConfig(p, i)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func testConfig(p *lang.Process, i int) error {
	option, _ := p.Parameters.String(i)

	switch option {
	case "enable", "on":
		p.Stdout.Writeln([]byte("Enabling test mode...."))
		return p.Config.Set("test", "enabled", true)

	case "!enable", "disable", "off":
		p.Stdout.Writeln([]byte("Disabling test mode...."))
		return p.Config.Set("test", "enabled", false)

	case "auto-report":
		p.Stdout.Writeln([]byte("Enabling auto-report...."))
		return p.Config.Set("test", "auto-report", true)

	case "!auto-report":
		p.Stdout.Writeln([]byte("Disabling auto-report...."))
		return p.Config.Set("test", "auto-report", false)

	case "verbose":
		p.Stdout.Writeln([]byte("Enabling verbose reporting...."))
		return p.Config.Set("test", "verbose", true)

	case "!verbose":
		p.Stdout.Writeln([]byte("Disabling verbose reporting...."))
		return p.Config.Set("test", "verbose", false)

	default:
		return fmt.Errorf(
			"Invalid parameter: `%s`%sExpected usage: test [ enable | !enable ] [ verbose | !verbose ] [ auto-report | !auto-report ]%s                test define { json-properties }%s                test state { code block }%s                test run { code-block }",
			option,
			utils.NewLineString,
			utils.NewLineString,
			utils.NewLineString,
			utils.NewLineString,
		)
	}
}

func cmdTestDisable(p *lang.Process) error {
	if p.Parameters.Len() > 0 {
		return errors.New("Too many parameters! Usage: `!test` to disable testing")
	}
	return p.Config.Set("test", "enabled", false)
}

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

func testState(p *lang.Process) error {
	name, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	block, err := p.Parameters.Block(2)
	if err != nil {
		return err
	}

	return p.Tests.State(name, block)
}

func testRun(p *lang.Process) error {
	block, err := p.Parameters.Block(1)
	if err != nil {
		return err
	}

	fork := p.Fork(lang.F_FUNCTION)
	fork.Name = "(test run)"

	err = fork.Config.Set("test", "enabled", true)
	if err != nil {
		return err
	}

	err = fork.Config.Set("test", "auto-report", true)
	if err != nil {
		return err
	}

	h := md5.New()
	_, err = h.Write([]byte(time.Now().String() + ":" + strconv.Itoa(p.Id)))
	if err != nil {
		return err
	}

	pipeName := "system_test_" + hex.EncodeToString(h.Sum(nil))

	err = lang.GlobalPipes.CreatePipe(pipeName, "std", "")
	if err != nil {
		return err
	}

	pipe, err := lang.GlobalPipes.Get(pipeName)
	if err != nil {
		return err
	}

	err = fork.Config.Set("test", "report-pipe", pipeName)
	if err != nil {
		return err
	}

	_, err = fork.Execute(block)
	if err != nil {
		return err
	}

	err = lang.GlobalPipes.Close(pipeName)
	if err != nil {
		return err
	}

	reportType, err := p.Config.Get("test", "report-format", types.String)
	if err != nil {
		return err
	}
	if reportType.(string) == "table" {
		p.Stderr.Writeln([]byte(consts.TestTableHeadings))
	}

	_, err = io.Copy(p.Stderr, pipe)
	return err
}
