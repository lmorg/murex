# proc.fg

> Sends a background process into the foreground

## Description

`fg` resumes a stopped process and sends it into the foreground.

## Usage

POSIX only:

```
fg fid
```

## Detail

This builtin is only supported on POSIX systems. There is no support planned
for Windows (due to the kernel not supporting the right signals) nor Plan 9.

## Synonyms

* `fg`
* `proc.fg`


## See Also

* [exec.file: `exec`](../commands/exec.md):
  Runs an executable
* [jobs](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [proc.bg](../commands/bg.md):
  Run processes in the background
* [proc.kill](../commands/fid-kill.md):
  Terminate a running Murex function
* [proc.kill.all](../commands/fid-killall.md):
  Terminate _all_ running Murex functions
* [proc.list](../commands/fid-list.md):
  Lists all running functions within the current Murex session

<hr/>

This document was generated from [builtins/core/processes/bgfg_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/processes/bgfg_doc.yaml).