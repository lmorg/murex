//go:build darwin
// +build darwin

package signaltrap

import "syscall"

var interrupts = map[string]syscall.Signal{
	syscall.SIGHUP,
	syscall.SIGINT,
	syscall.SIGQUIT,
	syscall.SIGILL,
	syscall.SIGTRAP,
	syscall.SIGABRT,
	syscall.SIGEMT,
	syscall.SIGFPE,
	syscall.SIGKILL,
	syscall.SIGBUS,
	syscall.SIGSEGV,
	syscall.SIGSYS,
	syscall.SIGPIPE,
	syscall.SIGALRM,
	syscall.SIGTERM,
	syscall.SIGURG,
	syscall.SIGSTOP,
	syscall.SIGTSTP,
	syscall.SIGCONT,
	syscall.SIGCHLD,
	syscall.SIGTTIN,
	syscall.SIGTTOU,
	syscall.SIGIO,
	syscall.SIGXCPU,
	syscall.SIGXFSZ,
	syscall.SIGVTALRM,
	syscall.SIGPROF,
	syscall.SIGWINCH,
	syscall.SIGINFO,
	syscall.SIGUSR1,
	syscall.SIGUSR2,
}
