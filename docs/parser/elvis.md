# `?:` Elvis Operator

> Returns the right operand if the left operand is empty

## Description

The elvis operator is a little like a conditional where the result of the
operation is the first non-empty value from left to right.

An empty value is any of the following:

* An unset / undefined variable
* Any value with a `null` data type

Other "falsy" values such as numerical values of `0`, boolean `false`, zero
length strings and strings containing `"null"` are not considered empty by the
elvis operator.



## Examples

**Assign a variable with a default value:**

```
» $foo = $bar ?: "baz"
```

If `$bar` is unset then the value of `$foo` will be **"baz"**.

**Multiple elvis operators:**

```
» $unset_variable ?: null ?: "foobar"
foobar
```

## Detail

### Whats in a name?

[Wikipedia](https://en.wikipedia.org/wiki/Elvis_operator) explains this best
where it says:

> The name "Elvis operator" refers to the fact that when its common notation,
> ?:, is viewed sideways, it resembles an emoticon of Elvis Presley with his
> signature hairstyle.

## See Also

* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Schedulers](../user-guide/schedulers.md):
  Overview of the different schedulers (or 'run modes') in Murex
* [`&&` And Logical Operator](../parser/logical-and.md):
  Continues next operation if previous operation passes
* [`?` STDERR Pipe](../parser/pipe-err.md):
  Pipes STDERR from the left hand command to STDIN of the right hand command
* [`err`](../commands/err.md):
  Print a line to the STDERR
* [`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [`try`](../commands/try.md):
  Handles errors inside a block of code
* [`trypipe`](../commands/trypipe.md):
  Checks state of each function in a pipeline and exits block on error
* [`||` Or Logical Operator](../parser/logical-or.md):
  Continues next operation only if previous operation fails
* [null](../commands/devnull.md):
  null function. Similar to /dev/null

<hr/>

This document was generated from [gen/parser/elvis_op_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/elvis_op_doc.yaml).