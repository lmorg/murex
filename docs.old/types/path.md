# Murex Shell Docs

## Data-Type Reference: `path` (string) 

> path data type

## Description

This type is modelled closely on generic but is more tailored for textual
(non-tabulated) data.

## Supported Hooks

* `Marshal()`
    Supported
* `ReadArray()`
    Treats each new directory as a new array element
* `ReadArrayWithType()`
    Treats each directory as a new array element, each array element is `str` 
* `ReadIndex()`
    Indexes treated as a path separated list
* `ReadMap()`
    Treats each new directory as a numbered map element
* `Unmarshal()`
    Supported
* `WriteArray()`
    Writes a new path, each array element as a directory

## See Also

* [`str` (string) ](../types/str.md):
  string (primitive)