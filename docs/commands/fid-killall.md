# `fid-killall` - Command Reference

> Terminate _all_ running _murex_ functions

## Description

`fid-killall` will terminate _all_ running _murex_ functions.

## Usage

    fid-killall

## Detail

`fid-killall` works by the same mechanisms as `fid-kill`, described below:

`fid-kill` doesn't send a kernel signal to the process since _murex_ is
a multi-threaded shell with a single signal, `fid-kill` will send a
cancellation context to any builtins executing (which covers builtins,
aliases, public and private functions and any external executables running
which were launched within the current _murex_ shell).

The FID (function ID) sent is not the same as a POSIX (eg Linux, macOS, BSD)
PID (process ID). You can obtain a FID from `fid-list`.

## See Also

* [`bexists`](../commands/bexists.md):
  Check which builtins exist
* [`bg`](../commands/bg.md):
  Run processes in the background
* [`builtins`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [`exec`](../commands/exec.md):
  Runs an executable
* [`fexec` ](../commands/fexec.md):
  Execute a command or function, bypassing the usual order of precedence.
* [`fg`](../commands/fg.md):
  Sends a background process into the foreground
* [`fid-kill`](../commands/fid-kill.md):
  Terminate a running _murex_ function
* [`fid-list`](../commands/fid-list.md):
  Lists all running functions within the current _murex_ session
* [`jobs`](../commands/fid-list.md):
  Lists all running functions within the current _murex_ session
* [`murex-update-exe-list`](../commands/murex-update-exe-list.md):
  Forces _murex_ to rescan $PATH looking for exectables