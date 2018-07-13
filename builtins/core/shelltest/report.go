package shelltest

import (
	"github.com/lmorg/murex/lang/proc"
)

func init() {
	proc.GoFunctions["test-report"] = cmdTestReport
}

func cmdTestReport(p *proc.Process) error {
	return p.Tests.WriteResults(p.Config, p.Stdout)
}
