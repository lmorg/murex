# `jsonl`

> JSON Lines

## Description

The following description is taken from [jsonlines.org](http://jsonlines.org/):

> JSON Lines is a convenient format for storing structured data that may be
> processed one record at a time. It works well with unix-style text
> processing tools and shell pipelines. It's a great format for log files.
> It's also a flexible format for passing messages between cooperating
> processes.

## Examples

Example JSON lines documents taken from [jsonlines.org](http://jsonlines.org/examples/)

### Tabulated data

```
["Name", "Session", "Score", "Completed"]
["Gilbert", "2013", 24, true]
["Alexa", "2013", 29, true]
["May", "2012B", 14, false]
["Deloise", "2012A", 19, true] 
```

This format is equatable to `generic` and `csv`.

### Nested objects

```
{"name": "Gilbert", "wins": [["straight", "7♣"], ["one pair", "10♥"]]}
{"name": "Alexa", "wins": [["two pair", "4♠"], ["two pair", "9♠"]]}
{"name": "May", "wins": []}
{"name": "Deloise", "wins": [["three of a kind", "5♣"]]}
```

## Detail

### Concatenated JSON

Technically the `jsonl` Unmarshal() method supports Concatenated JSON, as
described on [Wikipedia](https://en.wikipedia.org/wiki/JSON_streaming#Concatenated_JSON):

> Concatenated JSON streaming allows the sender to simply write each JSON
> object into the stream with no delimiters. It relies on the receiver using
> a parser that can recognize and emit each JSON object as the terminating
> character is parsed. Concatenated JSON isn't a new format, it's simply a
> name for streaming multiple JSON objects without any delimiters.
>
> The advantage of this format is that it can handle JSON objects that have
> been formatted with embedded newline characters, e.g., pretty-printed for
> human readability. For example, these two inputs are both valid and produce
> the same output:
>
> #### Single line concatenated JSON
>
>     {"some":"thing\n"}{"may":{"include":"nested","objects":["and","arrays"]}}
>
> #### Multi-line concatenated JSON
>
>     {
>       "some": "thing\n"
>     }
>     {
>       "may": {
>         "include": "nested",
>         "objects": [
>           "and",
>           "arrays"
>         ]
>       }
>     }

...however in Murex's case, only single line concatenated JSON files
(example 1) are supported; and that is only supported to cover some edge
cases when writing JSON lines and a new line character isn't included. The
primary example might be when generating JSON lines from inside a `for` loop.

This is resolved in the new data-type parser `jsonc` (Concatenated JSON). See
line below.

### More information

This format is sometimes also referred to as LDJSON and NDJSON, as described
on [Wikipedia](https://en.wikipedia.org/wiki/JSON_streaming#Line-delimited_JSON).

Murex's [`json` data-type document](json.md) also describes some use
cases for JSON lines.

## Default Associations

* **Extension**: `json-lines`
* **Extension**: `jsonl`
* **Extension**: `jsonlines`
* **Extension**: `ldjson`
* **Extension**: `murex_history`
* **Extension**: `ndjson`
* **MIME**: `application/json-lines`
* **MIME**: `application/jsonl`
* **MIME**: `application/jsonlines`
* **MIME**: `application/ldjson`
* **MIME**: `application/ndjson`
* **MIME**: `application/x-json-lines`
* **MIME**: `application/x-jsonl`
* **MIME**: `application/x-jsonlines`
* **MIME**: `application/x-ldjson`
* **MIME**: `application/x-ndjson`
* **MIME**: `text/json-lines`
* **MIME**: `text/jsonl`
* **MIME**: `text/jsonlines`
* **MIME**: `text/ldjson`
* **MIME**: `text/ndjson`
* **MIME**: `text/x-json-lines`
* **MIME**: `text/x-jsonl`
* **MIME**: `text/x-jsonlines`
* **MIME**: `text/x-ldjson`
* **MIME**: `text/x-ndjson`


## Supported Hooks

* `Marshal()`
    Supported
* `ReadArray()`
    Works with JSON arrays. Maps are converted into arrays
* `ReadArrayWithType()`
    Works with JSON arrays. Maps are converted into arrays. Element data type is `json` 
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

## See Also

* [Define Type (`cast`)](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [For Each In List (`foreach`)](../commands/foreach.md):
  Iterate through an array
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
* [`*` (generic)](../types/generic.md):
  generic (primitive)
* [`csv`](../types/csv.md):
  CSV files (and other character delimited tables)
* [`hcl`](../types/hcl.md):
  HashiCorp Configuration Language (HCL)
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)
* [`jsonc`](../types/jsonc.md):
  Concatenated JSON
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

This document was generated from [builtins/types/jsonlines/jsonlines_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/types/jsonlines/jsonlines_doc.yaml).