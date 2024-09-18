# Kill Function (`fid-kill`)

> Terminate a running Murex function

## Description

`fid-kill` will terminate a running Murex function in a similar way
that the POSIX `kill` (superficially speaking).

## Usage

```
fid-kill fid
```

## Detail

`fid-kill` doesn't send a kernel signal to the process since Murex is
a multi-threaded shell with a single signal, `fid-kill` will send a
cancellation context to any builtins executing (which covers builtins,
aliases, public and private functions and any external executables running
which were launched within the current Murex shell).

The FID (function ID) sent is not the same as a POSIX (eg Linux, macOS, BSD)
PID (process ID). You can obtain a FID from `fid-list`.

## Synonyms

* `fid-kill`


## See Also

* [Background Process (`bg`)](../commands/bg.md):
  Run processes in the background
* [Check Builtin Exists (`bexists`)](../commands/bexists.md):
  Check which builtins exist
* [Display Running Functions (`fid-list`)](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [Display Running Functions (`jobs`)](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [Execute External Command (`exec`)](../commands/exec.md):
  Runs an executable
* [Execute Shell Function or Builtin (`fexec`)](../commands/fexec.md):
  Execute a command or function, bypassing the usual order of precedence.
* [Foreground Process (`fg`)](../commands/fg.md):
  Sends a background process into the foreground
* [Kill All In Session (`fid-killall`)](../commands/fid-killall.md):
  Terminate all running Murex functions in current session
* [Re-Scan $PATH For Executables](../commands/murex-update-exe-list.md):
  Forces Murex to rescan $PATH looking for executables
* [Shell Runtime (`builtins`)](../commands/runtime.md):
  Returns runtime information on the internal state of Murex

<hr/>

This document was generated from [builtins/core/processes/kill_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/processes/kill_doc.yaml).