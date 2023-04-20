# `fg` - Command Reference

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

* [`bg`](../commands/bg.md):
  Run processes in the background
* [`exec`](../commands/exec.md):
  Runs an executable
* [`fid-kill`](../commands/fid-kill.md):
  Terminate a running _murex_ function
* [`fid-killall`](../commands/fid-killall.md):
  Terminate _all_ running _murex_ functions
* [`fid-list`](../commands/fid-list.md):
  Lists all running functions within the current _murex_ session
* [`jobs`](../commands/fid-list.md):
  Lists all running functions within the current _murex_ session