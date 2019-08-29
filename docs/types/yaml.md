# _murex_ Shell Guide

## Data-Type Reference: `yaml` (YAML)

> YAML Ain't Markup Language (YAML)

### Description

YAML support within _murex_ is pretty mature however it is not considered a
primitive. Which means, while it is a recommended builtin which you should
expect in most deployments of _murex_, it's still an optional package and
thus may not be present in some edge cases. This is because it relies on
external source packages for the shell to compile.



### Default Associations

* **Extension**: `yaml`
* **Extension**: `yml`
* **MIME**: `application/x-yaml`
* **MIME**: `application/yaml`
* **MIME**: `text/x-yaml`
* **MIME**: `text/yaml`


### Supported Hooks

* `Marshal()`
    Supported
* `ReadArray()`
    Works with YAML arrays. Maps are converted into arrays
* `ReadIndex()`
    Works against all properties in YAML
* `ReadMap()`
    Works with YAML maps
* `ReadNotIndex()`
    Works against all properties in YAML
* `Unmashal()`
    Supported
* `WriteArray()`
    Works with YAML arrays

### See Also

* [`Marshal()` ](../apis/marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [`Unmarshal()` ](../apis/unmarshal.md):
  Converts a structured file format into structured memory
* [`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [`json` (JSON)](../types/json.md):
  JavaScript Object Notation (JSON) (primitive)
* [`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [jsonl](../types/jsonl.md):
  
* [open](../commands/open.md):
  
* [readarray](../apis/readarray.md):
  
* [readindex](../apis/readindex.md):
  
* [readmap](../apis/readmap.md):
  
* [readnotindex](../apis/readnotindex.md):
  
* [writearray](../apis/writearray.md):
  