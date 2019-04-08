package cmdtest

import (
	"github.com/lmorg/murex/lang"
)

func init() {
	lang.GoFunctions["test-report"] = cmdTestReport
}

func cmdTestReport(p *lang.Process) error {
	return lang.ShellProcess.Tests.WriteResults(p.Config, p.Stdout)
}
