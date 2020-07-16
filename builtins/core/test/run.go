package cmdtest

import (
	"github.com/lmorg/murex/lang"
)

func testRun(p *lang.Process) error {
	r, err := p.Parameters.Block(1)
	if err == nil {
		return testRunBlock(p, r)
	}

	return testUnitRun(p)
}

func testRunBlock(p *lang.Process, block []rune) error {
	fork := p.Fork(lang.F_FUNCTION)
	fork.Name = "(test run)"

	err := fork.Config.Set("test", "enabled", true)
	if err != nil {
		return err
	}

	err = fork.Config.Set("test", "auto-report", false)
	if err != nil {
		return err
	}

	_, err = fork.Execute(block)
	if err != nil {
		return err
	}

	return p.Tests.WriteResults(p.Config, p.Stdout)
}
