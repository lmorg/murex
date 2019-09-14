package cmdtest

import (
	"errors"
	"fmt"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
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
				"run",
				"define-unit",
				"run-unit"
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

	case "define-unit":
		return testUnitDefine(p)

	case "state":
		return testState(p)

	case "run":
		return testRun(p)

	case "run-unit":
		return testUnitRun(p)

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
