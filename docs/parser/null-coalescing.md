# `??` Null Coalescing Operator

> Returns the right operand if the left operand is empty / undefined (expression)

## Description

The Null Coalescing operator is a little like a conditional where the result of the
operation is the first non-empty value from left to right.

An empty value is any of the following:

* an unset / undefined variable
* any value with a `null` data type

Other "falsy" values such as numerical values of `0`, boolean `false`, zero
length strings and strings containing `"null"` are not considered empty by the
null coalescing operator.



## Examples

### Assign with a default value

```
» $foo = $bar ?? "baz"
```

If `$bar` is unset then the value of `$foo` will be **"baz"**.

### Multiple operators

```
» $unset_variable ?? null ?? "foobar"
foobar
```

## Detail

The following extract was taken from [Wikipedia](https://en.wikipedia.org/wiki/Null_coalescing_operator):

> The null coalescing operator (called the Logical Defined-Or operator in Perl)
> is a binary operator that is part of the syntax for a basic conditional
> expression in several programming languages. While its behavior differs
> between implementations, the null coalescing operator generally returns the
> result of its left-most operand if it exists and is not null, and otherwise
> returns the right-most operand. This behavior allows a default value to be
> defined for cases where a more specific value is not available.
>
> In contrast to the ternary conditional if operator used as `x ? x : y`, but
> like the binary Elvis operator used as `x ?: y`, the null coalescing operator
> is a binary operator and thus evaluates its operands at most once, which is
> significant if the evaluation of `x` has side-effects. 

## See Also

* [Error String (`err`)](../commands/err.md):
  Print a line to the stderr
* [Expressions (`expr`)](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [Is Value Null (`is-null`)](../commands/is-null.md):
  Checks if a variable is null or undefined
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
* [`?:` Elvis Operator](../parser/elvis.md):
  Returns the right operand if the left operand is falsy (expression)
* [`?` stderr Pipe](../parser/pipe-err.md):
  Pipes stderr from the left hand command to stdin of the right hand command (DEPRECATED)
* [`||` Or Logical Operator](../parser/logical-or.md):
  Continues next operation only if previous operation fails
* [null](../commands/devnull.md):
  null function. Similar to /dev/null

<hr/>

This document was generated from [gen/expr/null-coalescing-op_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/expr/null-coalescing-op_doc.yaml).