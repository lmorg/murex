# Foreground Process (`fg`)

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


## See Also

* [Background Process (`bg`)](../commands/bg.md):
  Run processes in the background
* [Display Running Functions (`fid-list`)](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [Display Running Functions (`jobs`)](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [Execute External Command (`exec`)](../commands/exec.md):
  Runs an executable
* [Kill All In Session (`fid-killall`)](../commands/fid-killall.md):
  Terminate all running Murex functions in current session
* [Kill Function (`fid-kill`)](../commands/fid-kill.md):
  Terminate a running Murex function

<hr/>

This document was generated from [builtins/core/processes/bgfg_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/processes/bgfg_doc.yaml).