# _murex_ Shell Guide

## Data-Type Reference: `json` (JSON)

> JavaScript Object Notation (JSON) (primitive)

### Description

JSON is a primitive data-type within _murex_.



### Examples

Example JSON document taken from [Wikipedia](https://en.wikipedia.org/wiki/JSON)

    {
      "firstName": "John",
      "lastName": "Smith",
      "isAlive": true,
      "age": 27,
      "address": {
        "streetAddress": "21 2nd Street",
        "city": "New York",
        "state": "NY",
        "postalCode": "10021-3100"
      },
      "phoneNumbers": [
        {
          "type": "home",
          "number": "212 555-1234"
        },
        {
          "type": "office",
          "number": "646 555-4567"
        },
        {
          "type": "mobile",
          "number": "123 456-7890"
        }
      ],
      "children": [],
      "spouse": null
    }

### Default Associations

* **Extension**: `json`
* **MIME**: `application/json`


### Supported Hooks

* `Marshal()`
    Writes minified JSON when no TTY detected and human readable JSON when stdout is a TTY
* `ReadArray()`
    Works with JSON arrays. Maps are converted into arrays
* `ReadIndex()`
    Works against all properties in JSON
* `ReadMap()`
    Works with JSON maps
* `ReadNotIndex()`
    Works against all properties in JSON
* `Unmashal()`
    Supported
* `WriteArray()`
    Works with JSON arrays

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
* [types/`hcl` (HCL)](../types/hcl.md):
  HashiCorp Configuration Language (HCL)
* [commands/`pretty`](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [commands/`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [types/`toml` (TOML)](../types/toml.md):
  Tom's Obvious, Minimal Language (TOML)
* [types/`yaml` (YAML)](../types/yaml.md):
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
  