# `csv`

> CSV files (and other character delimited tables)

## Description

This data type can be used for not only CSV files but also TSV (tab separated)
or any other exotic characters used as a delimiter.

## Detail

The CSV parser is configurable via `config` (see link below for docs on how to
use `config`)

```
» config -> [csv]      
{
    "comment": {
        "Data-Type": "str",
        "Default": "#",
        "Description": "The prefix token for comments in a CSV table.",
        "Dynamic": false,
        "Global": false,
        "Value": "#"
    },
    "separator": {
        "Data-Type": "str",
        "Default": ",",
        "Description": "The delimiter for records in a CSV file.",
        "Dynamic": false,
        "Global": false,
        "Value": ","
    }
}
```

## Default Associations

* **Extension**: `csv`
* **MIME**: `application/csv`
* **MIME**: `application/x-csv`
* **MIME**: `text/csv`
* **MIME**: `text/x-csv`


## Supported Hooks

* `Marshal()`
    Supported
* `ReadArray()`
    Treats each new line as a new array element
* `ReadArrayWithType()`
    Treats each new line as a new array element, each element is mini `csv` file
* `ReadIndex()`
    Indexes treated as table coordinates
* `ReadMap()`
    Works against tables such as the output from `ps -fe` 
* `ReadNotIndex()`
    Indexes treated as table coordinates
* `Unmarshal()`
    Supported
* `WriteArray()`
    Writes a new line per array element

## See Also

* [Define Type (`cast`)](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [Get Nested Element (`[[ Element ]]`)](../parser/element.md):
  Outputs an element from a nested structure
* [Inline SQL (`select`)](../optional/select.md):
  Inlining SQL into shell pipelines
* [Reformat Data type (`format`)](../commands/format.md):
  Reformat one data-type into another data-type
* [Shell Configuration And Settings (`config`)](../commands/config.md):
  Query or define Murex runtime settings
* [`*` (generic)](../types/generic.md):
  generic (primitive)
* [`int`](../types/int.md):
  Whole number (primitive)
* [`jsonl`](../types/jsonl.md):
  JSON Lines
* [`str` (string)](../types/str.md):
  string (primitive)
* [index](../parser/item-index.md):
  Outputs an element from an array, map or table

### Read more about type hooks

- [`ReadIndex()` (type)](../apis/ReadIndex.md): Data type handler for the index, `[`, builtin
- [`ReadNotIndex()` (type)](../apis/ReadNotIndex.md): Data type handler for the bang-prefixed index, `![`, builtin
- [`ReadArray()` (type)](../apis/ReadArray.md): Read from a data type one array element at a time
- [`WriteArray()` (type)](../apis/WriteArray.md): Write a data type, one array element at a time
- [`ReadMap()` (type)](../apis/ReadMap.md): Treat data type as a key/value structure and read its contents
- [`Marshal()` (type)](../apis/Marshal.md): Converts structured memory into a structured file format (eg for stdio)
- [`Unmarshal()` (type)](../apis/Unmarshal.md): Converts a structured file format into structured memory

<hr/>

This document was generated from [builtins/types/csv/csv_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/types/csv/csv_doc.yaml).