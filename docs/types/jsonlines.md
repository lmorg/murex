# _murex_ Shell Guide

## Data-Type Reference: `jsonl` 

> JSON Lines (primitive)

### Description

The following description is taken from [jsonlines.org](http://jsonlines.org/):

> JSON Lines is a convenient format for storing structured data that may be
> processed one record at a time. It works well with unix-style text
> processing tools and shell pipelines. It's a great format for log files.
> It's also a flexible format for passing messages between cooperating
> processes.



### Examples

Example JSON lines document taken from [jsonlines.org](http://jsonlines.org/examples/)

    {"name": "Gilbert", "wins": [["straight", "7♣"], ["one pair", "10♥"]]}
    {"name": "Alexa", "wins": [["two pair", "4♠"], ["two pair", "9♠"]]}
    {"name": "May", "wins": []}
    {"name": "Deloise", "wins": [["three of a kind", "5♣"]]}

### Default Associations

* **Extension**: `jsonl`
* **Extension**: `jsonlines`
* **Extension**: `murex_history`
* **MIME**: `application/jsonl`
* **MIME**: `application/jsonlines`
* **MIME**: `application/x-jsonl`
* **MIME**: `application/x-jsonlines`
* **MIME**: `text/jsonl`
* **MIME**: `text/jsonlines`
* **MIME**: `text/x-jsonl`
* **MIME**: `text/x-jsonlines`


### Supported Hooks

* `Marshal()`
    Supported
* `ReadArray()`
    Works with JSON arrays. Maps are converted into arrays
* `ReadIndex()`
    Works against all properties in JSON
* `ReadMap()`
    Not currently supported.
* `ReadNotIndex()`
    Works against all properties in JSON
* `Unmarshal()`
    Supported
* `WriteArray()`
    Supported

### See Also

* [apis/`Marshal()` ](../apis/marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [apis/`Unmarshal()` ](../apis/unmarshal.md):
  Converts a structured file format into structured memory
* [commands/`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [commands/`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [commands/`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [commands/`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [types/`hcl` ](../types/hcl.md):
  HashiCorp Configuration Language (HCL)
* [commands/`pretty`](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [commands/`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [types/`toml` ](../types/toml.md):
  Tom's Obvious, Minimal Language (TOML)
* [types/`yaml` ](../types/yaml.md):
  YAML Ain't Markup Language (YAML)
* [types/jsonl](../types/jsonl.md):
  
* [types/mxjson](../types/mxjson.md):
  Murex-flavoured JSON (primitive)
* [commands/open](../commands/open.md):
  
* [apis/readarray](../apis/readarray.md):
  
* [apis/readindex](../apis/readindex.md):
  
* [apis/readmap](../apis/readmap.md):
  
* [apis/readnotindex](../apis/readnotindex.md):
  
* [apis/writearray](../apis/writearray.md):
  