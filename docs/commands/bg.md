# `bg`

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

## See Also

* [`exec`](../commands/exec.md):
  Runs an executable
* [`fg`](../commands/fg.md):
  Sends a background process into the foreground
* [`fid-kill`](../commands/fid-kill.md):
  Terminate a running Murex function
* [`fid-killall`](../commands/fid-killall.md):
  Terminate _all_ running Murex functions
* [`fid-list`](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [`jobs`](../commands/fid-list.md):
  Lists all running functions within the current Murex session

<hr/>

This document was generated from [builtins/core/processes/bgfg_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/processes/bgfg_doc.yaml).