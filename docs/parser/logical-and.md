# `&&` And Logical Operator

> Continues next operation if previous operation passes

## Description

When in the **normal** run mode (see "schedulers" link below) this will only
run the command on the right hand side if the command on the left hand side
does not error. Neither stdout nor stderr are piped.

This has no effect in `try` nor `trypipe` run modes because they automatically
apply stricter error handling.



## Examples

### When true

Second command runs because the first command doesn't error:

```
» out one && out two
one
two
```

When false

Second command does not run because the first command produces an error:

```
» err one && out two
one
```

## Detail

This is equivalent to a `try` block:

```
try {
    err one
    out two
}
```

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
* [`?:` Elvis Operator](../parser/elvis.md):
  Returns the right operand if the left operand is falsy (expression)
* [`?` stderr Pipe](../parser/pipe-err.md):
  Pipes stderr from the left hand command to stdin of the right hand command (DEPRECATED)
* [`||` Or Logical Operator](../parser/logical-or.md):
  Continues next operation only if previous operation fails

<hr/>

This document was generated from [gen/parser/logical_ops_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/logical_ops_doc.yaml).