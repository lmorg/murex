# `||` Or Logical Operator

> Continues next operation only if previous operation fails

## Description

When in the **normal** run mode (see "schedulers" link below) this will only
run the command on the right hand side if the command on the left hand side
does not error. Neither stdout nor stderr are piped.

This has no effect in `try` nor `trypipe` run modes because they automatically
apply stricter error handling. See detail below.



## Examples

### When true

Second command does not run because the first command doesn't error:

```
» out one || out two
one
```

### When false

Second command does run because the first command produces an error:

```
» err one || out two
one
two
```

## Detail

This has no effect in `try` nor `trypipe` run modes because they automatically
apply stricter error handling. You can achieve a similar behavior in `try` with
the following code:

```
try {
    err one -> !if { out two }
}
```

There is no workaround for `trypipe`.

## See Also

* [Error String (`err`)](../commands/err.md):
  Print a line to the stderr
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Pipe Fail (`trypipe`)](../commands/trypipe.md):
  Checks for non-zero exits of each function in a pipeline
* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Schedulers](../user-guide/schedulers.md):
  Overview of the different schedulers (or 'run modes') in Murex
* [Try Block (`try`)](../commands/try.md):
  Handles non-zero exits inside a block of code
* [`&&` And Logical Operator](../parser/logical-and.md):
  Continues next operation if previous operation passes
* [`?:` Elvis Operator](../parser/elvis.md):
  Returns the right operand if the left operand is falsy (expression)
* [`?` stderr Pipe](../parser/pipe-err.md):
  Pipes stderr from the left hand command to stdin of the right hand command (DEPRECATED)

<hr/>

This document was generated from [gen/parser/logical_ops_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/logical_ops_doc.yaml).