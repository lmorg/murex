# `?:` Elvis Operator

> Returns the right operand if the left operand is falsy (expression)

## Description

The Elvis Operator is a little like a conditional where the result of the
operation is the first non-falsy value from left to right.

A falsy value is any of the following:

* an unset / undefined variable
* any value with a `null` data type
* a `str` or generic with the value `false`, `null`, `0`, `no`, `off`, `fail`,
  `failed`, or `disabled`
* a number (`num`, `float` or `int`) with the value `0`
* an empty object or zero length array 
* and, of course, a boolean with the value `false`



## Examples

### Assign with a default value

```
» $foo = $bar ?: "baz"
```

If `$bar` is falsy, then the value of `$foo` will be **"baz"**.

### Multiple elvis operators

```
» $unset_variable ?: null ?: false ?: "foobar"
foobar
```

## Detail

### Whats in a name?

[Wikipedia](https://en.wikipedia.org/wiki/Elvis_operator) explains this best
where it says:

> The name "Elvis operator" refers to the fact that when its common notation,
> `?:`, is viewed sideways, it resembles an emoticon of Elvis Presley with his
> signature hairstyle.

## See Also

* [Error String (`err`)](../commands/err.md):
  Print a line to the stderr
* [Expressions (`expr`)](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [Operators And Tokens](../user-guide/operators-and-tokens.md):
  A table of all supported operators and tokens
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
* [`??` Null Coalescing Operator](../parser/null-coalescing.md):
  Returns the right operand if the left operand is empty / undefined (expression)
* [`?` stderr Pipe](../parser/pipe-err.md):
  Pipes stderr from the left hand command to stdin of the right hand command (DEPRECATED)
* [`||` Or Logical Operator](../parser/logical-or.md):
  Continues next operation only if previous operation fails
* [null](../commands/devnull.md):
  null function. Similar to /dev/null

<hr/>

This document was generated from [gen/expr/elvis-op_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/expr/elvis-op_doc.yaml).