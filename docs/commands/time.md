# `time`

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

```
» time sleep 5
5.000151513

» time { out "Going to sleep"; sleep 5; out "Waking up" }
Going to sleep
Waking up
5.000240977
```

## Detail

`time`'s output is written to STDERR. However any output and errors written
by the commands executed by time will also be written to `time`'s STDOUT
and STDERR as usual.

## See Also

* [`exec`](../commands/exec.md):
  Runs an executable
* [`source`](../commands/source.md):
  Import Murex code from another file of code block
* [sleep](../commands/sleep.md):
  

<hr/>

This document was generated from [builtins/core/time/time_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/time/time_doc.yaml).