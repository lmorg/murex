# _murex_ Shell Docs

## Data-Type Reference: `num` (number)

> Floating point number (primitive)

## Description

Any number. To be precise, a full set of all IEEE-754 64-bit floating-point
numbers.

> Unless you specifically know you only want whole numbers, it is recommended
> that you use this as your default numeric data-type as opposed to `int`.

## Supported Hooks

* `Marshal()`
    Supported
* `Unmashal()`
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
* [`int` ](../types/int.md):
  Whole number (primitive)
* [`open`](../commands/open.md):
  Open a file with a preferred handler
* [`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [`str` (string) ](../types/str.md):
  string (primitive)