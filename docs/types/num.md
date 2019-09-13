# _murex_ Shell Guide

## Data-Type Reference: `num` (number)

> Floating point number (primitive)

### Description

Any number. To be precise, a full set of all IEEE-754 64-bit floating-point
numbers.

> Unless you specifically know you only want whole numbers, it is recommended
> that you use this as your default numeric data-type as opposed to `int`.

### Supported Hooks

* `Marshal()`
    Supported
* `Unmashal()`
    Supported

### See Also

* [apis/`Marshal()` ](../apis/marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [apis/`Unmarshal()` ](../apis/unmarshal.md):
  Converts a structured file format into structured memory
* [commands/`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [commands/`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [commands/`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [commands/`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [types/`int` ](../types/int.md):
  Whole number (primitive)
* [commands/`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [commands/open](../commands/open.md):
  
* [types/str](../types/str.md):
  