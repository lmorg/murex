//go:build js
// +build js

package lang

import (
	"syscall"
)

func unixProcAttrFauxTTY() *syscall.SysProcAttr {
	return nil
}

func UnixPidToFg(_ *Process) {}
