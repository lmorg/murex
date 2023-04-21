# `jsonc`  - Data-Type Reference

> Concatenated JSON

## Description

The following description is taken from [Wikipedia](https://en.wikipedia.org/wiki/JSON_streaming#Concatenated_JSON):

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

## Examples

Because of the similiaries with jsonlines (`jsonl`), the examples here will
focus on jsonlines examples. However concatenated JSON doesn't need a new line
separator. So the examples below could all be concatenated into one long line.

Example JSON lines documents taken from [jsonlines.org](http://jsonlines.org/examples/)

### Tabulated data

    ["Name", "Session", "Score", "Completed"]
    ["Gilbert", "2013", 24, true]
    ["Alexa", "2013", 29, true]
    ["May", "2012B", 14, false]
    ["Deloise", "2012A", 19, true] 
    
This format is equatable to `generic` and `csv`.

### Nested objects

    {"name": "Gilbert", "wins": [["straight", "7♣"], ["one pair", "10♥"]]}
    {"name": "Alexa", "wins": [["two pair", "4♠"], ["two pair", "9♠"]]}
    {"name": "May", "wins": []}
    {"name": "Deloise", "wins": [["three of a kind", "5♣"]]}

## Detail

### Similarities with `jsonl`

The advantage of concatenated JSON is that it supports everything jsonlines
supports but without the dependency of a new line as a separator.

Eventually it is planned that this Murex data-type will replace jsonlines
and possibly even the regular JSON parser. However this concatenated JSON
parser currently requires reading the entire file first before parsing whereas
jsonlines can read one line at a time. Which makes jsonlines a better data-
type for pipelining super large documents. For this reason (and that this
parser is still in beta), it is shipped as an additional data-type.

## Default Associations

* **Extension**: `concatenated-json`
* **Extension**: `json-seq`
* **Extension**: `jsonc`
* **Extension**: `jsonconcat`
* **Extension**: `jsons`
* **Extension**: `jsonseq`
* **MIME**: `application/concatenated-json`
* **MIME**: `application/json-seq`
* **MIME**: `application/jsonc`
* **MIME**: `application/jsonconcat`
* **MIME**: `application/jsonseq`
* **MIME**: `application/x-concatenated-json`
* **MIME**: `application/x-json-seq`
* **MIME**: `application/x-jsonc`
* **MIME**: `application/x-jsonconcat`
* **MIME**: `application/x-jsonseq`
* **MIME**: `text/concatenated-json`
* **MIME**: `text/concatenated-json`
* **MIME**: `text/json-seq`
* **MIME**: `text/jsonc`
* **MIME**: `text/jsonconcat`
* **MIME**: `text/jsonseq`
* **MIME**: `text/x-json-seq`
* **MIME**: `text/x-jsonc`
* **MIME**: `text/x-jsonconcat`
* **MIME**: `text/x-jsonseq`


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

* [`*` (generic) ](../types/generic.md):
  generic (primitive)
* [`Marshal()` (type)](../apis/Marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [`ReadArray()` (type)](../apis/ReadArray.md):
  Read from a data type one array element at a time
* [`ReadIndex()` (type)](../apis/ReadIndex.md):
  Data type handler for the index, `[`, builtin
* [`ReadMap()` (type)](../apis/ReadMap.md):
  Treat data type as a key/value structure and read its contents
* [`ReadNotIndex()` (type)](../apis/ReadNotIndex.md):
  Data type handler for the bang-prefixed index, `![`, builtin
* [`Unmarshal()` (type)](../apis/Unmarshal.md):
  Converts a structured file format into structured memory
* [`WriteArray()` (type)](../apis/WriteArray.md):
  Write a data type, one array element at a time
* [`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`csv` ](../types/csv.md):
  CSV files (and other character delimited tables)
* [`foreach`](../commands/foreach.md):
  Iterate through an array
* [`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [`hcl` ](../types/hcl.md):
  HashiCorp Configuration Language (HCL)
* [`json` ](../types/json.md):
  JavaScript Object Notation (JSON) (primitive)
* [`jsonl` ](../types/jsonl.md):
  JSON Lines (primitive)
* [`open`](../commands/open.md):
  Open a file with a preferred handler
* [`pretty`](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of Murex
* [`toml` ](../types/toml.md):
  Tom's Obvious, Minimal Language (TOML)
* [`yaml` ](../types/yaml.md):
  YAML Ain't Markup Language (YAML)
* [mxjson](../types/mxjson.md):
  Murex-flavoured JSON (deprecated)