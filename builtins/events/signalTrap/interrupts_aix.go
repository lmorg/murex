//go:build aix
// +build aix

package signaltrap

import "syscall"

var interrupts = map[string]syscall.Signal{
	"SIGHUP":     syscall.SIGHUP,
	"SIGINT":     syscall.SIGINT,
	"SIGQUIT":    syscall.SIGQUIT,
	"SIGILL":     syscall.SIGILL,
	"SIGTRAP":    syscall.SIGTRAP,
	"SIGABRT":    syscall.SIGABRT,
	"SIGBUS":     syscall.SIGBUS,
	"SIGFPE":     syscall.SIGFPE,
	"SIGKILL":    syscall.SIGKILL,
	"SIGUSR1":    syscall.SIGUSR1,
	"SIGSEGV":    syscall.SIGSEGV,
	"SIGUSR2":    syscall.SIGUSR2,
	"SIGPIPE":    syscall.SIGPIPE,
	"SIGALRM":    syscall.SIGALRM,
	"SIGCHLD":    syscall.SIGCHLD,
	"SIGCONT":    syscall.SIGCONT,
	"SIGSTOP":    syscall.SIGSTOP,
	"SIGTSTP":    syscall.SIGTSTP,
	"SIGTTIN":    syscall.SIGTTIN,
	"SIGTTOU":    syscall.SIGTTOU,
	"SIGURG":     syscall.SIGURG,
	"SIGXCPU":    syscall.SIGXCPU,
	"SIGXFSZ":    syscall.SIGXFSZ,
	"SIGVTALRM":  syscall.SIGVTALRM,
	"SIGPROF":    syscall.SIGPROF,
	"SIGWINCH":   syscall.SIGWINCH,
	"SIGIO":      syscall.SIGIO,
	"SIGPWR":     syscall.SIGPWR,
	"SIGSYS":     syscall.SIGSYS,
	"SIGTERM":    syscall.SIGTERM,
	"SIGEMT":     syscall.SIGEMT,
	"SIGWAITING": syscall.SIGWAITING,
}
