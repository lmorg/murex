# `*=` Multiply By Operator (expr)

> Multiplies a variable by the right hand value

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
`/=`. Not even when `strict-types` is disabled.

You can work around this by using the slightly longer syntax: **variable =
value op value**, for example:

```
» $i = "3"
» $i = $i + "2"
» $i
5
```

Please note that this behaviour might change in a later release of Murex.

## See Also

* [`*` Multiplication Operator (expr)](../parser/multiplication.md):
  Multiplies one numeric value with another
* [`+=` Add With Operator (expr)](../parser/add-with.md):
  Adds the right hand value to a variable
* [`-=` Subtract By Operator (expr)](../parser/subtract-by.md):
  Subtracts a variable by the right hand value
* [`/=` Divide By Operator (expr)](../parser/divide-by.md):
  Divides a variable by the right hand value
* [`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`config`](../commands/config.md):
  Query or define Murex runtime settings
* [`expr`](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [`float` (floating point number)](../types/float.md):
  Floating point number (primitive)
* [`int`](../types/int.md):
  Whole number (primitive)
* [`num` (number)](../types/num.md):
  Floating point number (primitive)

<hr/>

This document was generated from [gen/expr/multiply_by_op_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/expr/multiply_by_op_doc.yaml).