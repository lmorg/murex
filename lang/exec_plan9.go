//go:build plan9
// +build plan9

package lang

import (
	"syscall"
)

func unixProcAttrFauxTTY() *syscall.SysProcAttr {
	return nil
}

func UnixPidToFg(_ *Process) {}
