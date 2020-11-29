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
private autocomplete.test.run-unit {
    runtime: --tests -> [ unit ] -> foreach: test {
        out: $test[function]
    } -> prepend *
}

test unit private autocomplete.test.run-unit {
    "StdoutRegex": (^(([-_./a-zA-Z0-9]+|\*)\n)+),
	"StdoutType":  "jsonl",
    "StdoutBlock": ({
        -> len -> set len;
        if { = len>0 } then {
            out "Len greater than 0"
        } else {
            err "No elements returned"
        }
	}),
	"StdoutIsArray": true
}

autocomplete set test { [
    {
        "FlagsDesc": {
			"config":   "Enable or disable boolean test states (more options available in ` + "`" + `config` + "`" + `)",
			"define":   "Define an inlined test",
            "state":    "Define a state report",
            "run":      "Execute a code block, function or unit test with inline testing enabled",
			"unit": 	"Define a unit test",
			"report":   "Write the test report (happens automatically by default)"
		},
        "FlagValues": {
			"config": [{
				"FlagsDesc": {
					"enable":       "Enable inlined tests",
					"!enable":      "Disable inlined tests",
					"auto-report":  "Automatically output report (default)",
					"!auto-report": "Do not automatically output report",
					"verbose":      "Verbose report",
					"!verbose":     "Consice report (default)"
				},
				"AllowMultiple": true,
				"Optional": false
    		}],
            "run": [{
                "Dynamic": ({ autocomplete.test.run-unit })
            }]
        }
    }
] }
    `)
}

func errUsage(invalidParameter string, err error) error {
	usage := fmt.Sprintf(`Expected usage:
    test: config [ enable|!enable ] [ verbose|!verbose ] [ auto-report|!auto-report ]
    test: define test-name { json-properties }
    test: unit function|private|open|event test-name { json-properties }
    test: state name { code block }
    test: run { code-block }
    test: run package/module/test-name|*
    test: report
    !test`)

	switch {
	case invalidParameter != "":
		return fmt.Errorf("Invalid parameter: `%s`%s%s", invalidParameter, utils.NewLineString, usage)
	case err != nil:
		return fmt.Errorf("%s%s%s", err.Error(), utils.NewLineString, usage)
	default:
		return errors.New(usage)
	}
}

type testArgs struct {
	StdoutBlock string
	StdoutRegex string
	StderrBlock string
	StderrRegex string
	ExitNum     int
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

	case "unit":
		return testUnitDefine(p)

	case "state":
		return testState(p)

	case "run":
		return testRun(p)

	case "run-unit":
		return testUnitRun(p)

	case "config":
		for i := 1; i < p.Parameters.Len(); i++ {
			err := testConfig(p, i)
			if err != nil {
				return err
			}
		}
		return nil

	case "report":
		return lang.ShellProcess.Tests.WriteResults(p.Config, p.Stdout)

	default:
		return errUsage(option, nil)
	}
}

func testConfig(p *lang.Process, i int) error {
	option, _ := p.Parameters.String(i)

	switch option {
	case "enable":
		p.Stdout.Writeln([]byte("Enabling test mode...."))
		return p.Config.Set("test", "enabled", true)

	case "!enable":
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
		return errUsage(option, nil)
	}
}

func cmdTestDisable(p *lang.Process) error {
	if p.Parameters.Len() > 0 {
		return errUsage("", errors.New("Too many parameters! Usage: `!test` to disable testing"))
	}
	return p.Config.Set("test", "enabled", false)
}
