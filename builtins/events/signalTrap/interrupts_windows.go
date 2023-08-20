//go:build windows
// +build windows

package signaltrap

import "syscall"

var interrupts = []syscall.Signal{
	syscall.SIGHUP,
	syscall.SIGINT,
	syscall.SIGQUIT,
	syscall.SIGILL,
	syscall.SIGTRAP,
	syscall.SIGABRT,
	syscall.SIGBUS,
	syscall.SIGFPE,
	syscall.SIGKILL,
	syscall.SIGSEGV,
	syscall.SIGPIPE,
	syscall.SIGALRM,
	syscall.SIGTERM,
}
