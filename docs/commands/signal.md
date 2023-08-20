# `signal`

> Sends a signal RPC

## Description

`signal` sends an operating system RPC (known as "signal") to a specified
process, identified via it's process ID ("pid").

The following quote from [Wikipedia explains what signals](https://en.wikipedia.org/wiki/Signal_(IPC))
are:

> Signals are standardized messages sent to a running program to trigger
> specific behavior, such as quitting or error handling. They are a limited
> form of inter-process communication (IPC), typically used in Unix, Unix-like,
> and other POSIX-compliant operating systems.
>
> A signal is an asynchronous notification sent to a process or to a specific
> thread within the same process to notify it of an event. Common uses of
> signals are to interrupt, suspend, terminate or kill a process.

## Usage

1. The first parameter is the process ID (int)
2. The second parameter is the signal name (str). This will be all in
   UPPERCASE and prefixed "SIG"

```
signal pid SIGNAL
```

## Examples

```
bg {
    exec <pid:GLOBAL.SIGNAL_TRAP_PID> $MUREX_EXE -c %(
        event signalTrap example=SIGINT {
            out "SIGINT received, not quitting"
        }
        sleep 4
    )
}
sleep 2 # just in case `exec` hasn't started yet
signal $GLOBAL.SIGNAL_TRAP_PID SIGINT
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
builtins/events/signalTrap/interrupts_windows.go
```

### Plan 9 Support

Plan 9 is not currently supported. Please raise a feature request on Github if
this is a feature you would like added.

### Catching incoming signals

Signals can be caught (often referred to as "trapped") in Murex with an event:
`signalTrap`. Read below for details.

## See Also

* [`MUREX_EXE` (path)](../variables/MUREX_EXE.md):
  Absolute path to running shell
* [`bg`](../commands/bg.md):
  Run processes in the background
* [`event`](../commands/event.md):
  Event driven programming for shell scripts
* [`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [`signalTrap`](../events/signaltrap.md):
  Trap OS signals