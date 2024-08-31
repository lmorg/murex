# `-` Subtraction Operator

> Subtracts one numeric value from another (expression)

## Description

The Subtraction Operator takes the right hand number from the left hand number
in an expression.



## Examples

### Expression

```
» 3-2
1
```

### Statement

```
out (3-2)
» 1
```

## Detail

### Type Safety

Because shells are historically untyped, you cannot always guarantee that a
numeric-looking value isn't a string. To solve this problem, by default Murex
assumes anything that looks like a number is a number when performing addition.

```
» str = "2"
» int = 3
» $str + $int
1
```

For occasions when type safety is more important than the convenience of silent
data casting, you can disable the above behaviour via `config` ([read more](/docs/user-guide/strict-types.md)):

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
* [`+` Addition Operator](../parser/addition.md):
  Adds two numeric values together (expression)
* [`-=` Subtract By Operator](../parser/subtract-by.md):
  Subtracts a variable by the right hand value (expression)
* [`/` Division Operator](../parser/division.md):
  Divides one numeric value from another (expression)
* [`float` (floating point number)](../types/float.md):
  Floating point number (primitive)
* [`int`](../types/int.md):
  Whole number (primitive)
* [`num` (number)](../types/num.md):
  Floating point number (primitive)

<hr/>

This document was generated from [gen/expr/subtraction-op_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/expr/subtraction-op_doc.yaml).