# _murex_ Shell Guide

## Data-Type Reference: `int` 

> Whole number (primitive)

### Description

An integer is a whole number (eg 1, 2, 3, 4) rather than one with a decimal
point (such as 1.1).

Integers in _murex_ are sized based on the bit (or word) size of the target
CPU.

A 386, ARMv6 or other 32bit build of _murex_ would see the range of from
`-2147483648` (negative) through `2147483647` (positive).

AMD64 or other 64bit built of _murex_ would see the range from
`-9223372036854775808` (negative) through `9223372036854775807` (positive).

> Unless you specifically know you only want whole numbers, it is recommended
> that you use the default numeric data-type: `num`.

### Supported Hooks

* `Marshal()`
    Supported
* `Unmarshal()`
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
* [types/`num` (number)](../types/num.md):
  Floating point number (primitive)
* [commands/`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [commands/open](../commands/open.md):
  
* [types/str](../types/str.md):
  