# _murex_ Language Guide

## Data-Type Reference: `toml` (TOML)

> Tom's Obvious, Minimal Language (TOML)

### Description

TOML support within _murex_ is pretty mature however it is not considered a
primitive. Which means, while it is a recommended builtin which you should
expect in most deployments of _murex_, it's still an optional package and
thus may not be present in some edge cases. This is because it relies on
external source packages for the shell to compile.



### Examples

Example TOML document taken from [Wikipedia](https://en.wikipedia.org/wiki/TOML)

    # This is a TOML document.
    
    title = "TOML Example"
    
    [owner]
    name = "Tom Preston-Werner"
    dob = 1979-05-27T07:32:00-08:00 # First class dates
    
    [database]
    server = "192.168.1.1"
    ports = [ 8001, 8001, 8002 ]
    connection_max = 5000
    enabled = true
    
    [servers]
    
      # Indentation (tabs and/or spaces) is allowed but not required
      [servers.alpha]
      ip = "10.0.0.1"
      dc = "eqdc10"
    
      [servers.beta]
      ip = "10.0.0.2"
      dc = "eqdc10"
    
    [clients]
    data = [ ["gamma", "delta"], [1, 2] ]
    
    # Line breaks are OK when inside arrays
    hosts = [
      "alpha",
      "omega"
    ]

### Default Associations

* MIME: `application/toml`
* MIME: `application/x-toml`
* MIME: `text/toml`
* MIME: `text/x-toml`
* Extension: `toml`


### Supported Hooks

* `Marshaller()`
    Supported
* `ReadIndex()`
    Works against all properties in TOML
* `ReadMap()`
    Works with TOML maps
* `ReadNotIndex()`
    Works against all properties in TOML
* `Unmashaller()`
    Supported

### See Also

* [`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`json` (JSON)](../types/json.md):
  JavaScript Object Notation (JSON) (primitive)
* [`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [`yaml` (YAML)](../types/yaml.md):
  YAML Ain't Markup Language (YAML)
* [element](../commands/element.md):
  
* [format](../commands/format.md):
  
* [jsonl](../types/jsonl.md):
  
* [open](../commands/open.md):
  