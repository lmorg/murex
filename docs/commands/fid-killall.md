# _murex_ Shell Docs

## Command Reference: `fid-killall`

> Terminate _all_ running _murex_ functions

## Description

`fid-killall` will terminate _all_ running _murex_ functions.

## Usage

    fid-killall

## Detail

`fid-killall` works by the same mechanisms as `fid-kill`, described below:

`fid-kill` doesn't send a kernel signal to the process but since _murex_ is
a multi-threaded shell with a single signal, `fid-kill` will send a
cancellation context to any builtins executing (which covers builtins,
aliases, public and private functions and any external executables running
which were launched within the current _murex_ shell).

The FID (function ID) sent is not the same as a POSIX (eg Linux, macOS, BSD)
PID (process ID). You can obtain a FID from `fid-list`.

## See Also

* [commands/`bg`](../commands/bg.md):
  Run processes in the background
* [commands/`exec`](../commands/exec.md):
  Runs an executable
* [commands/`fg`](../commands/fg.md):
  Sends a background process into the foreground
* [commands/`fid-kill`](../commands/fid-kill.md):
  Terminate a running _murex_ function
* [commands/`fid-list`](../commands/fid-list.md):
  Lists all running functions within the current _murex_ session
* [commands/`murex-update-exe-list`](../commands/murex-update-exe-list.md):
  Forces _murex_ to rescan $PATH looking for exectables
* [commands/bexists](../commands/bexists.md):
  
* [commands/builtins](../commands/builtins.md):
  
* [commands/jobs](../commands/jobs.md):
  