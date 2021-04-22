# _murex_ Shell Docs

## Data-Type Reference: `str` (string) 

> string (primitive)

## Description

This type is modelled closely on generic but is more tailored for textual
(non-tabulated) data.

## Supported Hooks

* `Marshal()`
    Supported
* `ReadArray()`
    Treats each new line as a new array element
* `ReadArrayByType()`
    Treats each new line as a new array element, each array element is `str
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

* [types/`*` (generic) ](../types/generic.md):
  generic (primitive)
* [apis/`Marshal()` (type)](../apis/Marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
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
* [types/`int` ](../types/int.md):
  Whole number (primitive)
* [types/`num` (number)](../types/num.md):
  Floating point number (primitive)
* [commands/`open`](../commands/open.md):
  Open a file with a preferred handler
* [commands/`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_