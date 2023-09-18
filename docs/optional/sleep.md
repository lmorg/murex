# `sleep`

> Suspends the shell for a number of seconds

## Description

`sleep` is an optional builtin which suspends the shell for a defined number
of seconds.

## Usage

```
sleep integer
```

## Examples

```
» sleep 5
# murex sleeps for 5 seconds
```

## Detail

`sleep` is very simplistic - particularly when compared to its GNU coreutil
(for example) counterpart. If you want to use the `sleep` binary on Linux
or similar platforms then you will need to launch with the `exec` builtin:

```
» exec sleep 5
```

## See Also

* [exec](../optional/exec.md):
  
* [source](../optional/source.md):
  
* [time](../optional/time.md):
  

<hr/>

This document was generated from [builtins/optional/time/time_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/optional/time/time_doc.yaml).