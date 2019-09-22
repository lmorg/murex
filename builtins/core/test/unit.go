package cmdtest

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/json"
)

func testUnitDefine(p *lang.Process) error {
	mod, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	function, err := p.Parameters.String(2)
	if err != nil {
		return err
	}

	b, err := p.Parameters.Byte(3)
	if err != nil {
		return err
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
		function = lang.UnitTestAutocomplete + "/" + function

	case "event":
		function = lang.UnitTestAutocomplete + "/" + function

	default:
		return fmt.Errorf("Unsupported block type (eg `function`, `private`, `autocomplete`): `%s`", mod)
	}

	lang.GlobalUnitTests.Add(function, plan, p.FileRef)

	return nil
}

func testUnitRun(p *lang.Process) error {
	function, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	if !lang.GlobalUnitTests.Run(p, function) {
		p.ExitNum = 1
	}

	return nil
}
