# `path`  - Data-Type Reference

> Structured object for working with file and directory paths

## Description

The `path` type Turns file and directory paths into structured objects

## Supported Hooks

* `Marshal()`
    Supported
* `ReadArray()`
    Each element is a directory branch. Root, `/`, is treated as it's own element
* `ReadArrayWithType()`
    Same as `ReadArray()
* `ReadIndex()`
    Returns a directory branch or filename if last element is a file
* `ReadMap()`
    Not currently supported
* `ReadNotIndex()`
    Supported
* `Unmarshal()`
    Supported
* `WriteArray()`
    Each element is a directory branch

## See Also

* [PWD](../variables/PWD.md):
  
* [PWDHIST](../variables/PWDHIST.md):
  
* [`Marshal()` (type)](../apis/Marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [`ReadArray()` (type)](../apis/ReadArray.md):
  Read from a data type one array element at a time
* [`ReadIndex()` (type)](../apis/ReadIndex.md):
  Data type handler for the index, `[`, builtin
* [`ReadMap()` (type)](../apis/ReadMap.md):
  Treat data type as a key/value structure and read its contents
* [`ReadNotIndex()` (type)](../apis/ReadNotIndex.md):
  Data type handler for the bang-prefixed index, `![`, builtin
* [`Unmarshal()` (type)](../apis/Unmarshal.md):
  Converts a structured file format into structured memory
* [`WriteArray()` (type)](../apis/WriteArray.md):
  Write a data type, one array element at a time
* [paths](../types/paths.md):
  