# _murex_ Shell Docs

## Command Reference: `fg`

> Sends a background process into the foreground

## Description

`fg` resumes a stopped process and sends it into the foreground.

## Usage

POSIX only:

    fg fid

## Detail

This builtin is only supported on POSIX systems. There is no support planned
for Windows (due to the kernel not supporting the right signals) nor Plan 9.

## See Also

* [commands/`bg`](../commands/bg.md):
  Run processes in the background
* [commands/`exec`](../commands/exec.md):
  Runs an executable
* [commands/`fid-kill`](../commands/fid-kill.md):
  Terminate a running _murex_ function
* [commands/`fid-killall`](../commands/fid-killall.md):
  Terminate _all_ running _murex_ functions
* [commands/`fid-list`](../commands/fid-list.md):
  Lists all running functions within the current _murex_ session
* [commands/jobs](../commands/jobs.md):
  