//go:build plan9
// +build plan9

package lang

import (
	"syscall"
)

func osExecFork(p *Process, argv []string) error {
	return execForkFallback(p, argv)
}

func unixProcAttrFauxTTY(_ int) *syscall.SysProcAttr {
	return nil
}

func UnixPidToFg(_ *Process) {}
