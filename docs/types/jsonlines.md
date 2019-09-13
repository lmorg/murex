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

### Detail

#### Concatenated JSON

Technically the `jsonl` Marshal() method supports Concatenated JSON, as
described on [Wikipedia]():

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
> **Single line concatenated JSON**
>
>     {"some":"thing\n"}{"may":{"include":"nested","objects":["and","arrays"]}}
>
> **Multi-line concatenated JSON**
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

...however in _murex_'s case, only single line concatenated JSON files
(example 1) are supported; and that is only supported to cover some edge
cases when writing JSON lines and a new line character isn't included. The
primary example might be when generating JSON lines from inside a 'for' loop.

This behavior is also described on GitHub in [issue #141](https://github.com/lmorg/murex/issues/141).

#### More information

This format is sometimes also referred to as LDJSON and NDJSON, as described
on [Wikipedia](https://en.wikipedia.org/wiki/JSON_streaming#Line-delimited_JSON).

_murex_'s [`json` data-type document](json.md) also describes some use
cases for JSON lines.

### Default Associations

* **Extension**: `jsonl`
* **Extension**: `jsonlines`
* **Extension**: `murex_history`
* **MIME**: `application/json-lines`
* **MIME**: `application/json-seq`
* **MIME**: `application/jsonl`
* **MIME**: `application/jsonlines`
* **MIME**: `application/jsonseq`
* **MIME**: `application/ldjson`
* **MIME**: `application/ndjson`
* **MIME**: `application/x-json-lines`
* **MIME**: `application/x-json-seq`
* **MIME**: `application/x-jsonl`
* **MIME**: `application/x-jsonlines`
* **MIME**: `application/x-jsonseq`
* **MIME**: `application/x-ldjson`
* **MIME**: `application/x-ndjson`
* **MIME**: `text/json-lines`
* **MIME**: `text/json-seq`
* **MIME**: `text/jsonl`
* **MIME**: `text/jsonlines`
* **MIME**: `text/jsonseq`
* **MIME**: `text/ldjson`
* **MIME**: `text/ndjson`
* **MIME**: `text/x-json-lines`
* **MIME**: `text/x-json-seq`
* **MIME**: `text/x-jsonl`
* **MIME**: `text/x-jsonlines`
* **MIME**: `text/x-jsonseq`
* **MIME**: `text/x-ldjson`
* **MIME**: `text/x-ndjson`


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
  