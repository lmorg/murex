//go:build windows
// +build windows

package lang

import (
	"syscall"

	"github.com/lmorg/murex/utils/which"
)

func osExecGetArgv(p *Process) []string {
	argv := []string{"cmd", "/c"}
	argv = append(argv, p.Parameters.StringArray()...)
	argv[2] = which.WhichIgnoreFail(argv[2])
	return argv
}

func osExecFork(p *Process, argv []string) error {
	return execForkFallback(p, argv)
}

func osSysProcAttr(_ int) *syscall.SysProcAttr {
	return nil
}
