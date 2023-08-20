The interrupts listed above are a subset of what is supported on each operating
system. Please consult your operating systems docs for details on each signal
and what their function is.

### Windows Support

While Windows doesn't officially support signals, the following POSIX signals
are emulated:

```go
var interrupts = map[string]syscall.Signal{
	"SIGHUP":  syscall.SIGHUP,
	"SIGINT":  syscall.SIGINT,
	"SIGQUIT": syscall.SIGQUIT,
	"SIGILL":  syscall.SIGILL,
	"SIGTRAP": syscall.SIGTRAP,
	"SIGABRT": syscall.SIGABRT,
	"SIGBUS":  syscall.SIGBUS,
	"SIGFPE":  syscall.SIGFPE,
	"SIGKILL": syscall.SIGKILL,
	"SIGSEGV": syscall.SIGSEGV,
	"SIGPIPE": syscall.SIGPIPE,
	"SIGALRM": syscall.SIGALRM,
	"SIGTERM": syscall.SIGTERM,
}
```

### Plan 9 Support

Plan 9 is not supported.