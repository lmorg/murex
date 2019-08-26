# _murex_ Shell Guide

## Command Reference: `fg`

> Sends a background process into the foreground

### Description

`fg` resumes a stopped process and sends it into the foreground.

### Usage

POSIX only:

    fg fid

### Detail

This builtin is only supported on POSIX systems. There is no support planned
for Windows (due to the kernel not supporting the right signals) nor Plan 9.

### See Also

* [`bg`](../commands/bg.md):
  Run processes in the background
* [fid-kill](../commands/fid-kill.md):
  
* [fid-killall](../commands/fid-killall.md):
  
* [fid-list](../commands/fid-list.md):
  