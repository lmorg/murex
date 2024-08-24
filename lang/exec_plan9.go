//go:build plan9
// +build plan9

package lang

import (
	"syscall"

	"github.com/lmorg/murex/utils/which"
)

func osExecGetArgv(p *Process) []string {
	argv := p.Parameters.StringArray()
	argv[0] = which.WhichIgnoreFail(argv[0])
	return argv
}

func osExecFork(p *Process, argv []string) error {
	return execForkFallback(p, argv)
}

func osSysProcAttr(_ int) *syscall.SysProcAttr {
	return nil
}
