# `num` (number)

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

* [Define Type (`cast`)](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [Get Nested Element (`[[ Element ]]`)](../parser/element.md):
  Outputs an element from a nested structure
* [Open File (`open`)](../commands/open.md):
  Open a file with a preferred handler
* [Reformat Data type (`format`)](../commands/format.md):
  Reformat one data-type into another data-type
* [Shell Runtime (`runtime`)](../commands/runtime.md):
  Returns runtime information on the internal state of Murex
* [`int`](../types/int.md):
  Whole number (primitive)
* [`str` (string)](../types/str.md):
  string (primitive)
* [index](../parser/item-index.md):
  Outputs an element from an array, map or table

### Read more about type hooks

- [`ReadIndex()` (type)](../apis/ReadIndex.md): Data type handler for the index, `[`, builtin
- [`ReadNotIndex()` (type)](../apis/ReadNotIndex.md): Data type handler for the bang-prefixed index, `![`, builtin
- [`ReadArray()` (type)](../apis/ReadArray.md): Read from a data type one array element at a time
- [`WriteArray()` (type)](../apis/WriteArray.md): Write a data type, one array element at a time
- [`ReadMap()` (type)](../apis/ReadMap.md): Treat data type as a key/value structure and read its contents
- [`Marshal()` (type)](../apis/Marshal.md): Converts structured memory into a structured file format (eg for stdio)
- [`Unmarshal()` (type)](../apis/Unmarshal.md): Converts a structured file format into structured memory

<hr/>

This document was generated from [builtins/types/numeric/numeric_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/types/numeric/numeric_doc.yaml).