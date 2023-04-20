# `int`  - Data-Type Reference

> Whole number (primitive)

## Description

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

## Supported Hooks

* `Marshal()`
    Supported
* `Unmarshal()`
    Supported

## See Also

* [`Marshal()` (type)](../apis/Marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [`Unmarshal()` (type)](../apis/Unmarshal.md):
  Converts a structured file format into structured memory
* [`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [`num` (number)](../types/num.md):
  Floating point number (primitive)
* [`open`](../commands/open.md):
  Open a file with a preferred handler
* [`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [`str` (string) ](../types/str.md):
  string (primitive)