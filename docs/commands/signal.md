# `signal`

> Sends a signal RPC

## Description

`signal` sends an operating system RPC (known as "signal") to a specified
process, identified via it's process ID ("pid").

The following quote from [Wikipedia](https://en.wikipedia.org/wiki/Signal_(IPC))
explains what signals are:

> Signals are standardized messages sent to a running program to trigger
> specific behavior, such as quitting or error handling. They are a limited
> form of inter-process communication (IPC), typically used in Unix, Unix-like,
> and other POSIX-compliant operating systems.
>
> A signal is an asynchronous notification sent to a process or to a specific
> thread within the same process to notify it of an event. Common uses of
> signals are to interrupt, suspend, terminate or kill a process.

### Listing supported signals

Signals will differ from one operating system to another. You can retrieve a
JSON map with supported signals by running `signal` without any parameters.

## Usage

**Send a signal:**

1. The first parameter is the process ID (int)
2. The second parameter is the signal name (str). This will be all in
   UPPERCASE and prefixed "SIG"

```
signal pid SIGNAL
```

**List supported signals:**

```
signal -> <stdout>
```

## Examples

### Send a signal

```
function signal.SIGUSR1.trap {
    bg {
        exec <pid:MOD.SIGNAL_TRAP_PID> $MUREX_EXE -c %(
            event onSignalReceived example=SIGUSR1 {
                out "SIGUSR1 received..."
            }

            out "waiting for signal..."
            sleep 5
        )
    }
    sleep 2 # just in case `exec` hasn't started yet
    signal $MOD.SIGNAL_TRAP_PID SIGUSR1
}

test unit function signal.SIGUSR1.trap %{
    StdoutMatch: "waiting for signal...\nSIGUSR1 received...\n"
    DataType:    str
    ExitNum:     0
}
```

### List supported signals

```
» signal
{
    "SIGABRT": "aborted",
    "SIGALRM": "alarm clock",
    "SIGBUS": "bus error",
    "SIGCHLD": "child exited",
    "SIGCONT": "continued",
    "SIGFPE": "floating point exception",
    "SIGHUP": "hangup",
    "SIGILL": "illegal instruction",
    "SIGINT": "interrupt",
    "SIGIO": "I/O possible",
    "SIGKILL": "killed",
    "SIGPIPE": "broken pipe",
    "SIGPROF": "profiling timer expired",
    "SIGPWR": "power failure",
    "SIGQUIT": "quit",
    "SIGSEGV": "segmentation fault",
    "SIGSTKFLT": "stack fault",
    "SIGSTOP": "stopped (signal)",
    "SIGSYS": "bad system call",
    "SIGTRAP": "trace/breakpoint trap",
    "SIGTSTP": "stopped",
    "SIGTTIN": "stopped (tty input)",
    "SIGTTOU": "stopped (tty output)",
    "SIGURG": "urgent I/O condition",
    "SIGUSR1": "user defined signal 1",
    "SIGUSR2": "user defined signal 2",
    "SIGVTALRM": "virtual timer expired",
    "SIGWINCH": "window changed",
    "SIGXCPU": "CPU time limit exceeded",
    "SIGXFSZ": "file size limit exceeded"
}
```

## Flags

* `SIGINT`
    **"Signal interrupt"** -- equivalent to pressing `ctrl`+`c`
* `SIGQUIT`
    **"Signal quit"** -- requests the process quits and performs a core dump
* `SIGTERM`
    **"Signal terminate"** -- request for a processes termination. Similar to `SIGINT`
* `SIGUSR1`
    **"Signal user 1"** -- user defined
* `SIGUSR2`
    **"Signal user 2"** -- user defined

## Detail

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

### Catching incoming signals

Signals can be caught (often referred to as "trapped") in Murex with an event:
`signalTrap`. Read below for details.

## See Also

* [Background Process (`bg`)](../commands/bg.md):
  Run processes in the background
* [Interactive Shell](../user-guide/interactive-shell.md):
  What's different about Murex's interactive shell?
* [MUREX_EXE](../variables/murex_exe.md):
  Absolute path to running shell
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Terminal Hotkeys](../user-guide/terminal-keys.md):
  A list of all the terminal hotkeys and their uses
* [`event`](../commands/event.md):
  Event driven programming for shell scripts
* [onSignalReceived](../events/onsignalreceived.md):
  Trap OS signals

<hr/>

This document was generated from [builtins/events/onSignalReceived/signal_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/events/onSignalReceived/signal_doc.yaml).