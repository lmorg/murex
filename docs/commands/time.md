# Processes Execution Time (`time`)

> Returns the execution run time of a command or block

## Description

`time` is an optional builtin which runs a command or block of code and
returns it's running time.

## Usage

```
time command parameters -> <stderr>

time { code-block } -> <stderr>
```

## Examples

### Time a command

```
» time sleep 5
5.000151513
```

### Time a block of code

```
» time { out "Going to sleep"; sleep 5; out "Waking up" }
Going to sleep
Waking up
5.000240977
```

## Detail

`time`'s output is written to stderr. However any output and errors written
by the commands executed by time will also be written to `time`'s stdout
and stderr as usual.

## See Also

* [Execute External Command (`exec`)](../commands/exec.md):
  Runs an executable
* [Include / Evaluate Murex Code (`source`)](../commands/source.md):
  Import Murex code from another file or code block
* [`sleep`](../optional/sleep.md):
  Suspends the shell for a number of seconds

<hr/>

This document was generated from [builtins/core/time/time_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/time/time_doc.yaml).