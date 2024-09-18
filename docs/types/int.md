# `int`

> Whole number (primitive)

## Description

An integer is a whole number (eg 1, 2, 3, 4) rather than one with a decimal
point (such as 1.1).

Integers in Murex are sized based on the bit (or word) size of the target
CPU.

A 386, ARMv6 or other 32bit build of Murex would see the range of from
`-2147483648` (negative) through `2147483647` (positive).

AMD64 or other 64bit built of Murex would see the range from
`-9223372036854775808` (negative) through `9223372036854775807` (positive).

> Unless you specifically know you only want whole numbers, it is recommended
> that you use the default numeric data-type: `num`.

## Supported Hooks

* `Marshal()`
    Supported
* `Unmarshal()`
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
* [`num` (number)](../types/num.md):
  Floating point number (primitive)
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