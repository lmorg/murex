# _murex_ Shell Guide

## Data-Type Reference: `json` 

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

### Detail

#### Tips when writing JSON inside for loops

One of the drawbacks (or maybe advantages, depending on your perspective) of
JSON is that parsers generally expect a complete file for processing in that
the JSON specification requires closing tags for every opening tag. This means
it's not always suitable for streaming. For example

    » ja [1..3] -> foreach i { out ({ "$i": $i }) }
    { "1": 1 }
    { "2": 2 }
    { "3": 3 }
    
**What does this even mean and how can you build a JSON file up sequentially?**

One answer if to write the output in a streaming file format and convert back
to JSON

    » ja [1..3] -> foreach i { out (- "$i": $i) }
    - "1": 1
    - "2": 2
    - "3": 3
    
    » ja [1..3] -> foreach i { out (- "$i": $i) } -> cast yaml -> format json
    [
        {
            "1": 1
        },
        {
            "2": 2
        },
        {
            "3": 3
        }
    ]
    
**What if I'm returning an object rather than writing one?**

The problem with building JSON structures from existing structures is that you
can quickly end up with invalid JSON due to the specifications strict use of
commas.

    » config -> [ shell ] -> formap k v { $v -> alter /Foo Bar }
    {
        "Data-Type": "bool",
        "Default": true,
        "Description": "Display the interactive shell's hint text helper. Please note, even when this is disabled, it will still appear when used for regexp searches and other readline-specific functions",
        "Dynamic": false,
        "Foo": "Bar",
        "Global": true,
        "Value": true
    }
    {
        "Data-Type": "block",
        "Default": "{ progress $PID }",
        "Description": "Murex function to execute when an `exec` process is stopped",
        "Dynamic": false,
        "Foo": "Bar",
        "Global": true,
        "Value": "{ progress $PID }"
    }
    {
        "Data-Type": "bool",
        "Default": true,
        "Description": "ANSI escape sequences in Murex builtins to highlight syntax errors, history completions, {SGR} variables, etc",
        "Dynamic": false,
        "Foo": "Bar",
        "Global": true,
        "Value": true
    }
    ...
    
Luckily JSON also has it's own streaming format: JSON lines (`jsonl`)

    » config -> [ shell ] -> formap k v { $v -> alter /Foo Bar } -> cast jsonl -> format json
    [
        {
            "Data-Type": "bool",
            "Default": true,
            "Description": "Write shell history (interactive shell) to disk",
            "Dynamic": false,
            "Foo": "Bar",
            "Global": true,
            "Value": true
        },
        {
            "Data-Type": "int",
            "Default": 4,
            "Description": "Maximum number of lines with auto-completion suggestions to display",
            "Dynamic": false,
            "Foo": "Bar",
            "Global": true,
            "Value": "6"
        },
        {
            "Data-Type": "bool",
            "Default": true,
            "Description": "Display some status information about the stop process when ctrl+z is pressed (conceptually similar to ctrl+t / SIGINFO on some BSDs)",
            "Dynamic": false,
            "Foo": "Bar",
            "Global": true,
            "Value": true
        },
    ...

### Default Associations

* **Extension**: `json`
* **MIME**: `application/json`
* **MIME**: `application/x-toml`
* **MIME**: `text/toml`
* **MIME**: `text/x-toml`


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
* `Unmarshal()`
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
  