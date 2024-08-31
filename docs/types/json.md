# `json`

> JavaScript Object Notation (JSON)

## Description

JSON is a structured data-type within Murex. It is the standard format for all
structured data within Murex however other formats such as YAML, TOML and CSV
are equally first class citizens.

## Examples

Example JSON document taken from [Wikipedia](https://en.wikipedia.org/wiki/JSON)

```
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
```

## Detail

### Tips when writing JSON inside for loops

One of the drawbacks (or maybe advantages, depending on your perspective) of
JSON is that parsers generally expect a complete file for processing in that
the JSON specification requires closing tags for every opening tag. This means
it's not always suitable for streaming. For example

```
» ja [1..3] -> foreach i { out ({ "$i": $i }) }
{ "1": 1 }
{ "2": 2 }
{ "3": 3 }
```

**What does this even mean and how can you build a JSON file up sequentially?**

One answer if to write the output in a streaming file format and convert back
to JSON

```
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
```

**What if I'm returning an object rather than writing one?**

The problem with building JSON structures from existing structures is that you
can quickly end up with invalid JSON due to the specifications strict use of
commas.

For example in the code below, each item block is it's own object and there are
no `[ ... ]` encapsulating them to denote it is an array of objects, nor are
the objects terminated by a comma.

```
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
```

Luckily JSON also has it's own streaming format: JSON lines (`jsonl`). We can
`cast` this output as `jsonl` then `format` it back into valid JSON:

```
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
```

#### `foreach` will automatically cast it's output as `jsonl` _if_ it's stdin type is `json`

```
» ja [Tom,Dick,Sally] -> foreach name { out Hello $name }
Hello Tom
Hello Dick
Hello Sally

» ja [Tom,Dick,Sally] -> foreach name { out Hello $name } -> debug -> [[ /Data-Type/Murex ]]
jsonl

» ja [Tom,Dick,Sally] -> foreach name { out Hello $name } -> format json
[
    "Hello Tom",
    "Hello Dick",
    "Hello Sally"
]
```

## Default Associations

* **Extension**: `json`
* **MIME**: `application/json`
* **MIME**: `application/x-json`
* **MIME**: `text/json`
* **MIME**: `text/x-json`


## Supported Hooks

* `Marshal()`
    Writes minified JSON when no TTY detected and human readable JSON when stdout is a TTY
* `ReadArray()`
    Works with JSON arrays. Maps are converted into arrays
* `ReadArrayWithType()`
    Works with JSON arrays. Maps are converted into arrays. Elements data-type in Murex mirrors the JSON type of the element
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

## See Also

* [Define Type (`cast`)](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [Get Nested Element (`[[ Element ]]`)](../parser/element.md):
  Outputs an element from a nested structure
* [Open File (`open`)](../commands/open.md):
  Open a file with a preferred handler
* [Prettify JSON](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [Reformat Data type (`format`)](../commands/format.md):
  Reformat one data-type into another data-type
* [Shell Runtime (`runtime`)](../commands/runtime.md):
  Returns runtime information on the internal state of Murex
* [`hcl`](../types/hcl.md):
  HashiCorp Configuration Language (HCL)
* [`jsonc`](../types/jsonc.md):
  Concatenated JSON
* [`jsonl`](../types/jsonl.md):
  JSON Lines
* [`toml`](../types/toml.md):
  Tom's Obvious, Minimal Language (TOML)
* [`yaml`](../types/yaml.md):
  YAML Ain't Markup Language (YAML)
* [index](../parser/item-index.md):
  Outputs an element from an array, map or table
* [mxjson](../types/mxjson.md):
  Murex-flavoured JSON (deprecated)

### Read more about type hooks

- [`ReadIndex()` (type)](../apis/ReadIndex.md): Data type handler for the index, `[`, builtin
- [`ReadNotIndex()` (type)](../apis/ReadNotIndex.md): Data type handler for the bang-prefixed index, `![`, builtin
- [`ReadArray()` (type)](../apis/ReadArray.md): Read from a data type one array element at a time
- [`WriteArray()` (type)](../apis/WriteArray.md): Write a data type, one array element at a time
- [`ReadMap()` (type)](../apis/ReadMap.md): Treat data type as a key/value structure and read its contents
- [`Marshal()` (type)](../apis/Marshal.md): Converts structured memory into a structured file format (eg for stdio)
- [`Unmarshal()` (type)](../apis/Unmarshal.md): Converts a structured file format into structured memory

<hr/>

This document was generated from [builtins/types/json/json_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/types/json/json_doc.yaml).