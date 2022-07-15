package cmdtest

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/json"
)

func testUnitDefine(p *lang.Process) error {
	mod, err := p.Parameters.String(1)
	if err != nil {
		return errUsage("", err)
	}

	function, err := p.Parameters.String(2)
	if err != nil {
		return errUsage("", err)
	}

	b, err := p.Parameters.Byte(3)
	if err != nil {
		return errUsage("", err)
	}

	plan := new(lang.UnitTestPlan)
	err = json.UnmarshalMurex(b, plan)
	if err != nil {
		return err
	}

	switch mod {
	case "function":
		// do nothing

	case "private":
		function = p.FileRef.Source.Module + "/" + function

	case "autocomplete":
		function = lang.UnitTestAutocomplete + "/" + function

	case "open":
		function = lang.UnitTestOpen + "/" + function

	case "event":
		function = lang.UnitTestEvent + "/" + function

	default:
		return errUsage("", fmt.Errorf("Unsupported block type (eg `function`, `private`, `event`): `%s`", mod))
	}

	lang.GlobalUnitTests.Add(function, plan, p.FileRef)

	return nil
}

func testUnitRun(p *lang.Process) error {
	function, err := p.Parameters.String(1)
	if err != nil {
		return errUsage("", err)
	}

	err = p.Config.Set("test", "enabled", true, p.FileRef)
	if err != nil {
		return err
	}

	err = p.Config.Set("test", "auto-report", true, p.FileRef)
	if err != nil {
		return err
	}

	if !lang.GlobalUnitTests.Run(p, function) {
		p.ExitNum = 1
	}

	return p.Tests.WriteResults(p.Config, p.Stdout)
}
