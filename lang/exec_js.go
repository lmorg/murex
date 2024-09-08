//go:build js
// +build js

package lang

import (
	"syscall"
)

func osExecFork(p *Process, argv []string) error {
	return execForkFallback(p, argv)
}

func unixProcAttrFauxTTY() *syscall.SysProcAttr {
	return nil
}

func UnixPidToFg(_ *Process) {}
