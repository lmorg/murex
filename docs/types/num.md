# _murex_ Language Guide

## Data-Type Reference: `num` (number)

> Floating point number (primitive)

### Description

Any number. To be precise, a full set of all IEEE-754 64-bit floating-point
numbers.

> Unless you specifically know you only want whole numbers, it is recommended
> that you use this as your default numeric data-type as opposed to `int`.



### Default Associations




### Supported Hooks

* `Marshaller()`
    Supported
* `Unmashaller()`
    Supported

### See Also

* [`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`int` (integer)](../types/int.md):
  Whole number (primitive)
* [`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [element](../commands/element.md):
  
* [format](../commands/format.md):
  
* [open](../commands/open.md):
  
* [str](../types/str.md):
  