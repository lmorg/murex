# `fg`

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

- [`bg`](./bg.md):
  Run processes in the background
- [`exec`](./exec.md):
  Runs an executable
- [`fid-kill`](./fid-kill.md):
  Terminate a running Murex function
- [`fid-killall`](./fid-killall.md):
  Terminate _all_ running Murex functions
- [`fid-list`](./fid-list.md):
  Lists all running functions within the current Murex session
- [`jobs`](./fid-list.md):
  Lists all running functions within the current Murex session
