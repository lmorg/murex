# `commonlog`

> Apache httpd "common" log format

## Description

Apache httpd supports a few different log formats. This Murex type is for
parsing the "common" log format.

## Detail

The code here is very rudimentary. If you have large log files or need more complex
data querying then this data-type is probably not the right tool. Maybe try one of
the following:

* [Firesword](https://github.com/lmorg/firesword) - for command line analysis
* [Plasmasword](https://github.com/lmorg/plasmasword) - exports fields to an sqlite3 or mysql database

## Supported Hooks

* `Marshal()`
    Supported though no unmarshalling is currently supported
* `ReadArray()`
    Supported. Each line is considered an index (like with `str` data type)
* `ReadArrayWithType()`
    Supported. Each line is considered an index with `commonlog` data type
* `ReadIndex()`
    Entire log file is read and then the indexes are derived from there
* `ReadMap()`
    Not supported, currently a work in progress
* `ReadNotIndex()`
    Entire log file is read and then the indexes are derived from there

## See Also

* [Define Type (`cast`)](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [Get Nested Element (`[[ Element ]]`)](../parser/element.md):
  Outputs an element from a nested structure
* [Reformat Data type (`format`)](../commands/format.md):
  Reformat one data-type into another data-type
* [`*` (generic)](../types/generic.md):
  generic (primitive)
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

This document was generated from [builtins/types/apachelogs/apachelogs_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/types/apachelogs/apachelogs_doc.yaml).