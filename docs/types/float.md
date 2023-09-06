# `float` (floating point number)

> Floating point number (primitive)

## Description

Any number. To be precise, a full set of all IEEE-754 64-bit floating-point
numbers.

> This data-type is going to be deprecated in favour of `num` (since it is
> literally the same underlying data-type anyway). Please do not use `float`

## Supported Hooks

* `Marshal()`
    Supported
* `Unmarshal()`
    Supported

## See Also

* [`int`](../types/int.md):
  Whole number (primitive)
* [`num` (number)](../types/num.md):
  Floating point number (primitive)

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