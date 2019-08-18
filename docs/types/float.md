# _murex_ Language Guide

## Data-Type Reference: `float` (floating point number)

> Floating point number (primitive)

### Description

Any number. To be precise, a full set of all IEEE-754 64-bit floating-point
numbers.

> This data-type is going to be deprecated in favour of `num` (since it is
> literally the same underlying data-type anyway). Please do not use `float



### Supported Hooks

* `Marshal()`
    Supported
* `Unmashal()`
    Supported

### See Also

* [`Marshal()](../apis/marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [`int` (integer)](../types/int.md):
  Whole number (primitive)
* [`num` (number)](../types/num.md):
  Floating point number (primitive)
* [unmarshal](../apis/unmarshal.md):
  