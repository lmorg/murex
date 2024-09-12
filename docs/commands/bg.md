# Background Process (`bg`)

> Run processes in the background

## Description

`bg` supports two modes: it can either be run as a function block which will
execute in the background, or it can take stopped processes and daemonize
them.

## Usage

POSIX only:

 ```
 bg { code block }

 bg fid
 ```

## Examples


```
bg { sleep 5; out "Morning" }
```

## Detail

The examples above will work on any system (Windows included). However the
`ctrl+z` usage of backgrounding a stopped process (like Bash) is only
supported on POSIX systems due to the limitation of required signals on
non-platforms. This means the usage described in the examples is cross
cross platform while `bg int` currently does not work on Windows nor Plan 9.

## Synonyms

* `bg`


## See Also

* [Display Running Functions (`fid-list`)](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [Display Running Functions (`jobs`)](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [Execute External Command (`exec`)](../commands/exec.md):
  Runs an executable
* [Foreground Process (`fg`)](../commands/fg.md):
  Sends a background process into the foreground
* [Kill All In Session (`fid-killall`)](../commands/fid-killall.md):
  Terminate all running Murex functions in current session
* [Kill Function (`fid-kill`)](../commands/fid-kill.md):
  Terminate a running Murex function

<hr/>

This document was generated from [builtins/core/processes/bgfg_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/processes/bgfg_doc.yaml).