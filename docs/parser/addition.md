# `+` Addition Operator (expr)

> Adds two numeric values together

## Description

The Addition Operator adds two numeric values together in an expression. Those
values are placed either side of the addition operator.



## Examples

#### Expression

```
» 3+2
5
```

#### Statement

```
out (3+2)
» 5
```

## Detail

Unlike in some other programming languages, the `+` operator cannot be used to
concatenate strings. This is because shells are historically untyped so you
cannot always guarantee that numeric-looking value isn't a string. To solve
this problem, by default Murex assumes anything that looks like a number is a
number when performing addition. Thus overloading the `+` operator to
concatenate strings would lead to a large class of bugs.

```
» str = "3"
» int = 2
» $str + $int
5
```

For occasions when type safety is more important than the convenience of silent
data casting, you can disable the above behaviour via `config`:

```
» config set proc strict-types false
» $str + $int
Error in `expr` (0,1): cannot Add with string types
                    > Expression: $str + $int
                    >           : ^
                    > Character : 1
                    > Symbol    : Scalar
                    > Value     : '$str'
```

## See Also

* [`*` Multiplication Operator (expr)](../parser/multiplication.md):
  Multiplies one numeric value with another
* [`+=` Add With Operator (expr)](../parser/add-with.md):
  Adds the right hand value to a variable
* [`-` Subtraction Operator (expr)](../parser/subtraction.md):
  Subtracts one numeric value from another
* [`/` Division Operator (expr)](../parser/division.md):
  Divides one numeric value from another
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

This document was generated from [gen/expr/addition_op_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/expr/addition_op_doc.yaml).