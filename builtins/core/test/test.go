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

	"github.com/lmorg/murex/builtins/pipes/streams"
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
				"run",
				"define"
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

func cmdTest(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	if p.Parameters.Len() == 0 {
		//return errors.New("Missing parameters.")
		s, err := p.Config.Get("test", "enabled", types.String)
		if err != nil {
			return err
		}
		_, err = p.Stdout.Write([]byte(s.(string)))
		return err
	}

	if p.Parameters.Len() == 1 {
		return testConfig(p)
	}

	option, _ := p.Parameters.String(0)
	switch option {
	case "define":
		return testDefine(p)

	case "run":
		return testRun(p)

	default:
		return errors.New("Invalid parameter: " + option)
	}
}

func testConfig(p *lang.Process) error {
	option, _ := p.Parameters.String(0)

	switch option {
	case "enable":
		return p.Config.Set("test", "enabled", true)

	case "!enable", "disable":
		return p.Config.Set("test", "enabled", false)

	case "auto-report":
		return p.Config.Set("test", "auto-report", true)

	case "!auto-report":
		return p.Config.Set("test", "auto-report", false)

	default:
		v := types.IsTrue([]byte(option), 0)
		p.Stderr.Writeln([]byte(fmt.Sprintf(
			"Invalid parameter. Assuming `test %t`%sExpected usage: test [ enable | !enable ]%s                test [ auto-report | !auto-report ]%s                test define { json-properties }%s                test run { code-block }",
			v,
			utils.NewLineString,
			utils.NewLineString,
			utils.NewLineString,
			utils.NewLineString,
		)))
		return p.Config.Set("test", "enabled", v)
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

func runBlock(p *lang.Process, block []rune) ([]byte, error) {
	stdout := streams.NewStdin()
	_, err := lang.RunBlockExistingConfigSpace(block, nil, stdout, lang.ShellProcess.Stderr, p)
	if err != nil {
		return nil, err
	}

	b, err := stdout.ReadAll()
	if err != nil {
		return nil, err
	}
	return utils.CrLfTrim(b), nil
}

func testRun(p *lang.Process) error {
	block, err := p.Parameters.Block(1)
	if err != nil {
		return err
	}

	branch := p.BranchFID()
	defer branch.Close()

	err = branch.Config.Set("test", "enabled", true)
	if err != nil {
		return err
	}

	err = branch.Config.Set("test", "auto-report", true)
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

	err = branch.Config.Set("test", "report-pipe", pipeName)
	if err != nil {
		return err
	}

	_, err = lang.RunBlockExistingConfigSpace(block, p.Stdin, p.Stdout, p.Stderr, branch.Process)
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
