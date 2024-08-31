# `*=` Multiply By Operator

> Multiplies a variable by the right hand value (expression)

## Description

The Multiply By operator takes the value of the variable specified on the left
side of the operator and multiplies it with the value on the right hand side
Then it assigns the result back to the variable specified on the left side.

It is ostensibly just shorthand for `$i = $i * value`.

This operator is only available in expressions.



## Examples

```
» $i = 3
» $i *= 2
» $i
5
```

## Detail

### Strict Types

Unlike with the standard arithmetic operators (`+`, `-`, `*`, `/`), silent data
casting isn't supported with arithmetic assignments like `+=`, `-=`, `*=` and
`/=`. Not even when `strict-types` is disabled ([read more](/docs/user-guide/strict-types.md))

You can work around this by using the slightly longer syntax: **variable =
value op value**, for example:

```
» $i = "3"
» $i = $i + "2"
» $i
5
```

## See Also

* [Define Type (`cast`)](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [Expressions (`expr`)](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [Operators And Tokens](../user-guide/operators-and-tokens.md):
  A table of all supported operators and tokens
* [Shell Configuration And Settings (`config`)](../commands/config.md):
  Query or define Murex runtime settings
* [`*` Multiplication Operator](../parser/multiplication.md):
  Multiplies one numeric value with another (expression)
* [`+=` Add With Operator](../parser/add-with.md):
  Adds the right hand value to a variable (expression)
* [`-=` Subtract By Operator](../parser/subtract-by.md):
  Subtracts a variable by the right hand value (expression)
* [`/=` Divide By Operator](../parser/divide-by.md):
  Divides a variable by the right hand value (expression)
* [`float` (floating point number)](../types/float.md):
  Floating point number (primitive)
* [`int`](../types/int.md):
  Whole number (primitive)
* [`num` (number)](../types/num.md):
  Floating point number (primitive)

<hr/>

This document was generated from [gen/expr/multiply-by-op_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/expr/multiply-by-op_doc.yaml).