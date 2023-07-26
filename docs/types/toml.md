# `toml`

> Tom's Obvious, Minimal Language (TOML)

## Description

TOML support within Murex is pretty mature however it is not considered a
primitive. Which means, while it is a recommended builtin which you should
expect in most deployments of Murex, it's still an optional package and
thus may not be present in some edge cases. This is because it relies on
external source packages for the shell to compile.

## Examples

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

## Default Associations

- **Extension**: `toml`
- **MIME**: `application/toml`
- **MIME**: `application/x-toml`
- **MIME**: `text/toml`
- **MIME**: `text/x-toml`

## Supported Hooks

- `Marshal()`
  Supported
- `ReadArray()`
  Hook supported albeit TOML doesn't support naked arrays
- `ReadArrayWithType()`
  Hook supported albeit TOML doesn't support naked arrays
- `ReadIndex()`
  Works against all properties in TOML
- `ReadNotIndex()`
  Works against all properties in TOML
- `Unmarshal()`
  Supported

## See Also

- [`Marshal()` (type)](/apis/Marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
- [`ReadArray()` (type)](/apis/ReadArray.md):
  Read from a data type one array element at a time
- [`ReadIndex()` (type)](/apis/ReadIndex.md):
  Data type handler for the index, `[`, builtin
- [`ReadMap()` (type)](/apis/ReadMap.md):
  Treat data type as a key/value structure and read its contents
- [`ReadNotIndex()` (type)](/apis/ReadNotIndex.md):
  Data type handler for the bang-prefixed index, `![`, builtin
- [`Unmarshal()` (type)](/apis/Unmarshal.md):
  Converts a structured file format into structured memory
- [`WriteArray()` (type)](/apis/WriteArray.md):
  Write a data type, one array element at a time
- [`[[` (element)](/commands/element.md):
  Outputs an element from a nested structure
- [`[` (index)](/commands/index2.md):
  Outputs an element from an array, map or table
- [`cast`](/commands/cast.md):
  Alters the data type of the previous function without altering it's output
- [`format`](/commands/format.md):
  Reformat one data-type into another data-type
- [`json` ](/types/json.md):
  JavaScript Object Notation (JSON)
- [`jsonl` ](/types/jsonl.md):
  JSON Lines
- [`open`](/commands/open.md):
  Open a file with a preferred handler
- [`runtime`](/commands/runtime.md):
  Returns runtime information on the internal state of Murex
- [`yaml` ](/types/yaml.md):
  YAML Ain't Markup Language (YAML)
