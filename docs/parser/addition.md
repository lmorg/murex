# `+` Addition Operator

> Adds two numeric values together (expression)

## Description

The Addition Operator adds two numeric values together in an expression. Those
values are placed either side of the addition operator.



## Examples

### Expression

```
» 3+2
5
```

### Statement

```
out (3+2)
» 5
```

## Detail

### String Concatenation

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

### Type Safety

For occasions when type safety is more important than the convenience of silent
data casting, you can disable the above behaviour via `config`:

```
» config set proc strict-types true
» $str + $int
Error in `expr` (0,1): cannot Add with string types
                    > Expression: $str + $int
                    >           : ^
                    > Character : 1
                    > Symbol    : Scalar
                    > Value     : '$str'
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
* [Strict Types In Expressions](../user-guide/strict-types.md):
  Expressions can auto-convert types or strictly honour data types
* [`*` Multiplication Operator](../parser/multiplication.md):
  Multiplies one numeric value with another (expression)
* [`+=` Add With Operator](../parser/add-with.md):
  Adds the right hand value to a variable (expression)
* [`-` Subtraction Operator](../parser/subtraction.md):
  Subtracts one numeric value from another (expression)
* [`/` Division Operator](../parser/division.md):
  Divides one numeric value from another (expression)
* [`float` (floating point number)](../types/float.md):
  Floating point number (primitive)
* [`int`](../types/int.md):
  Whole number (primitive)
* [`num` (number)](../types/num.md):
  Floating point number (primitive)

<hr/>

This document was generated from [gen/expr/addition-op_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/expr/addition-op_doc.yaml).