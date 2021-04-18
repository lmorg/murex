# _murex_ Shell Docs

## Data-Type Reference: `commonlog` 

> Apache httpd "common" log format

## Description

Apache httpd supports a few different log formats. This _murex_ type is for
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
    Supported. Each line is considered an index (like with `str` data-type)
* `ReadIndex()`
    Entire log file is read and then the indexes are derived from there
* `ReadMap()`
    Not supported, currently a work in progress
* `ReadNotIndex()`
    Entire log file is read and then the indexes are derived from there

## See Also

* [types/`*` (generic) ](../types/generic.md):
  generic (primitive)
* [apis/`Marshal()` (type)](../apis/Marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [apis/`ReadArray()` (type)](../apis/ReadArray.md):
  Read from a data type one array element at a time
* [apis/`ReadIndex()` (type)](../apis/ReadIndex.md):
  Data type handler for the index, `[`, builtin
* [apis/`ReadMap()` (type)](../apis/ReadMap.md):
  Treat data type as a key/value structure and read its contents
* [apis/`ReadNotIndex()` (type)](../apis/ReadNotIndex.md):
  Data type handler for the bang-prefixed index, `![`, builtin
* [apis/`Unmarshal()` (type)](../apis/Unmarshal.md):
  Converts a structured file format into structured memory
* [commands/`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [commands/`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [commands/`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [commands/`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [types/`str` (string) ](../types/str.md):
  string (primitive)