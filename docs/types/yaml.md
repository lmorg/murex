# `yaml`

> YAML Ain't Markup Language (YAML)

## Description

YAML support within Murex is pretty mature however it is not considered a
primitive. Which means, while it is a recommended builtin which you should
expect in most deployments of Murex, it's still an optional package and
thus may not be present in some edge cases. This is because it relies on
external source packages for the shell to compile.

## Default Associations

* **Extension**: `yaml`
* **Extension**: `yml`
* **MIME**: `application/x-yaml`
* **MIME**: `application/yaml`
* **MIME**: `text/x-yaml`
* **MIME**: `text/yaml`


## Supported Hooks

* `Marshal()`
    Supported
* `ReadArray()`
    Works with YAML arrays. Maps are converted into arrays
* `ReadArrayWithType()`
    Works with YAML arrays. Maps are converted into arrays. Element type returned in Murex should match element type in YAML
* `ReadIndex()`
    Works against all properties in YAML
* `ReadMap()`
    Works with YAML maps
* `ReadNotIndex()`
    Works against all properties in YAML
* `Unmarshal()`
    Supported
* `WriteArray()`
    Works with YAML arrays

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
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)
* [`jsonl`](../types/jsonl.md):
  JSON Lines
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

This document was generated from [builtins/types/yaml/yaml_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/types/yaml/yaml_doc.yaml).