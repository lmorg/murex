//go:build windows
// +build windows

package lang

import (
	"syscall"
)

func osExecGetArgv(p *Process) []string {
	argv := []string{"cmd", "/c"}
	argv = append(argv, p.Parameters.StringArray()...)
	return argv
}

func osExecFork(p *Process, argv []string) error {
	return execForkFallback(p, argv)
}

func osSysProcAttr(_ int) *syscall.SysProcAttr {
	return nil
}
