//go:build js
// +build js

package lang

import (
	"syscall"
)

func osExecFork(p *Process, argv []string) error {
	return execForkFallback(p, argv)
}

func osSysProcAttr(_ int) *syscall.SysProcAttr {
	return nil
}
