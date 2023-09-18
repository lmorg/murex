# `str` (string)

> string (primitive)

## Description

This type is modelled closely on generic but is more tailored for textual
(non-tabulated) data.

## Supported Hooks

* `Marshal()`
    Supported
* `ReadArray()`
    Treats each new line as a new array element
* `ReadArrayWithType()`
    Treats each new line as a new array element, each array element is `str` 
* `ReadIndex()`
    Indexes treated as a new line separated list
* `ReadMap()`
    Treats each new line as a numbered map element
* `ReadNotIndex()`
    Indexes treated as a new line separated list
* `Unmarshal()`
    Supported
* `WriteArray()`
    Writes a new line per array element

## See Also

* [`*` (generic)](../types/generic.md):
  generic (primitive)
* [`int`](../types/int.md):
  Whole number (primitive)
* [`num` (number)](../types/num.md):
  Floating point number (primitive)
* [cast](../types/cast.md):
  
* [element](../types/element.md):
  
* [format](../types/format.md):
  
* [index](../types/index.md):
  
* [open](../types/open.md):
  
* [runtime](../types/runtime.md):
  

### Read more about type hooks

- [`ReadIndex()` (type)](../apis/ReadIndex.md): Data type handler for the index, `[`, builtin
- [`ReadNotIndex()` (type)](../apis/ReadNotIndex.md): Data type handler for the bang-prefixed index, `![`, builtin
- [`ReadArray()` (type)](../apis/ReadArray.md): Read from a data type one array element at a time
- [`WriteArray()` (type)](../apis/WriteArray.md): Write a data type, one array element at a time
- [`ReadMap()` (type)](../apis/ReadMap.md): Treat data type as a key/value structure and read its contents
- [`Marshal()` (type)](../apis/Marshal.md): Converts structured memory into a structured file format (eg for stdio)
- [`Unmarshal()` (type)](../apis/Unmarshal.md): Converts a structured file format into structured memory

<hr/>

This document was generated from [builtins/types/string/string_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/types/string/string_doc.yaml).