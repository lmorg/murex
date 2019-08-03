# _murex_ Language Guide

## Data-Type Reference: `yaml` (YAML)

> YAML Ain't Markup Language

### Description

YAML support within _murex_ is pretty mature however it is not considered a
primitive. Which means, while it is a recommended builtin which you should
expect in most deployments of _murex_, it's still an optional package and
thus may not be present in some edge cases. This is because it relies on
external source packages for the shell to compile.



### Default File Extensions And MIME Types

* MIME: `application/yaml`
* MIME: `application/x-yaml`
* MIME: `text/yaml`
* MIME: `text/x-yaml`
* Extension: `yaml`
* Extension: `yml`


### Supported Hooks

* `Marshaller()`
    Supported
* `ReadArray()`
    Works with YAML arrays. Maps are converted into arrays
* `ReadIndex()`
    Works against all properties in YAML
* `ReadMap()`
    Works with YAML maps
* `ReadNotIndex()`
    Works against all properties in YAML
* `Unmashaller()`
    Supported
* `WriteArray()`
    Works with YAML arrays

### See Also

* [`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [element](../commands/element.md):
  
* [format](../commands/format.md):
  
* [json](../types/json.md):
  
* [jsonl](../types/jsonl.md):
  
* [open](../commands/open.md):
  