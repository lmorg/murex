The interrupts listed above are a subset of what is supported on each operating
system. Please consult your operating systems docs for details on each signal
and what their function is.

### Windows Support

While Windows doesn't officially support signals, the following POSIX signals
are emulated:

```go
{{ include "builtins/events/signalTrap/interrupts_windows.go" }}
```

### Plan 9 Support

Plan 9 is not currently supported. Please raise a feature request on Github if
this is a feature you would like added.