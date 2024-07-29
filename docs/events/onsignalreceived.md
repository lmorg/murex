# `onSignalReceived`

> Trap OS signals

## Description

`onSignalReceived` events are triggered by OS signals.

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

This event is designed to be used in shell scripts. While this event can be
used with the shell in interactive mode (ie from the REPL prompt), this might
result in unexpected behaviour. Thus it is only recommended to use
`onSignalReceived` for shell scripts.

## Usage

```
event onSignalReceived name=SIGNAL { code block }

!event onSignalReceived [SIGNAL]name
```

## Valid Interrupts

* `SIGHUP`
    **"Signal hangup"** -- triggered when a controlling terminal is closed (eg the terminal emulator closed)
* `SIGINT`
    **"Signal interrupt"** -- triggered when a user interrupts a process, typically via `ctrl`+`c`
* `SIGQUIT`
    **"Signal quit"** -- when the user requests that the process quits and performs a core dump
* `SIGTERM`
    **"Signal terminate"** -- triggered by a request for a processes termination. Similar to `SIGINT`
* `SIGUSR1`
    **"Signal user 1"** -- user defined
* `SIGUSR2`
    **"Signal user 2"** -- user defined
* `SIGWINCH`
    **"Signal window change"** -- triggered when the TTY (eg terminal emulator) is resized

## Payload

The following payload is passed to the function via stdin:

```
{
    "Name": "",
    "Interrupt": {
        "Name": "",
        "Signal": ""
    }
}
```

### Name

This is the **namespaced** name -- ie the name and operation.

### Interrupt/Name

This is the name you specified when defining the event.

### Interrupt/Signal

This is the signal you specified when defining the event.

Valid interrupt operation values are specified below. All interrupts / signals
are UPPERCASE strings.

## Event Return

This event doesn't have any `$EVENT_RETURN` parameters.

## Examples

Interrupt 'SIGINT'

```
event onSignalReceived example=SIGINT {
    out "SIGINT received, not quitting"
}
```

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

### Stdout

Stdout and stderr are both written to the terminal.

### Order of execution

Interrupts are run in alphabetical order. So an event named "alfa" would run
before an event named "zulu". If you are writing multiple events and the order
of execution matters, then you can prefix the names with a number, eg `10_jump`

### Namespacing

The `onSignalReceived` event differs a little from other events when it comes
to the namespacing of interrupts. Typically you cannot have multiple interrupts
with the same name for an event. However with `onPrompt` their names are
further namespaced by the interrupt name. In layman's terms this means
`example=SIGINT` wouldn't overwrite `example=SIGQUIT`.

The reason for this namespacing is because, unlike other events, you might
legitimately want the same name for different interrupts.

## See Also

* [Interactive Shell](../user-guide/interactive-shell.md):
  What's different about Murex's interactive shell?
* [Terminal Hotkeys](../user-guide/terminal-keys.md):
  A list of all the terminal hotkeys and their uses
* [`event`](../commands/event.md):
  Event driven programming for shell scripts
* [`onCommandCompletion`](../events/oncommandcompletion.md):
  Trigger an event upon a command's completion
* [`onPrompt`](../events/onprompt.md):
  Events triggered by changes in state of the interactive shell
* [`signal`](../commands/signal.md):
  Sends a signal RPC

<hr/>

This document was generated from [builtins/events/onSignalReceived/onSignalReceived_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/events/onSignalReceived/onSignalReceived_doc.yaml).