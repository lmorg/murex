# Data-Type Reference

This section is a glossary of data-types which Murex is natively aware.

Most of the time you will not need to worry about typing in Murex as the
shell is designed around productivity as opposed to strictness despite
generally following a strictly typed design.

Read the [Language Tour](/docs/tour.md) for more detail on this topic.

## Definitions

For clarity, it is worth explaining a couple of terms:

1. _Data-types_ in Murex _data-types_ are an annotation describing the format
   of data contained in a pipe or variable. _Data-types_ can be _primitives_ or
   structured documents like JSON, CSV, and s-expressions.

   Objects like maps and arrays are just documents (typically JSON) however
   because Murex's builtin commands and expressions work consistently across a
   multitude of different document types, those JSON objects and CSV tables
   (et al) feel as native as Murex _data-types_, as strings do in Bash, s-expr
   in LISP and JSON in JavaScript.

2. _Primitives_ refer to the atomic component of a _data-type_. In other words,
   the smallest possible format for a piece of data. Where a JSON file might
   arrays and maps, the values for those objects cannot be divided any smaller
   than numbers, strings or a small number of constants like `true`, `false`,
   and `null`.

   In Murex, these are defined as _primitives_ and the following _data-types_
   are considered to be _primitive types_:

   * Numeric: `int`, `float` and `num`
  
   * Boolean: `bool`
  
   * Text: `string` and `*` (generic)
  
   * Null: `null`

## Feature Sets

Since not all data formats are equal (for example the TOML file format
doesn't support naked arrays where as JSON does), you may find some
features missing in some data-types which are present in others. If in
doubt then refer to the manual here or check the API manual for more
details about specific hooks.

## Pages

* [`*` (generic)](../types/generic.md):
  generic (primitive)
* [`bool`](../types/bool.md):
  Boolean (primitive)
* [`commonlog`](../types/commonlog.md):
  Apache httpd "common" log format
* [`csv`](../types/csv.md):
  CSV files (and other character delimited tables)
* [`float` (floating point number)](../types/float.md):
  Floating point number (primitive)
* [`hcl`](../types/hcl.md):
  HashiCorp Configuration Language (HCL)
* [`int`](../types/int.md):
  Whole number (primitive)
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)
* [`jsonc`](../types/jsonc.md):
  Concatenated JSON
* [`jsonl`](../types/jsonl.md):
  JSON Lines
* [`num` (number)](../types/num.md):
  Floating point number (primitive)
* [`path`](../types/path.md):
  Structured object for working with file and directory paths
* [`paths`](../types/paths.md):
  Structured array for working with `$PATH` style data
* [`str` (string)](../types/str.md):
  string (primitive)
* [`toml`](../types/toml.md):
  Tom's Obvious, Minimal Language (TOML)
* [`xml`](../types/xml.md):
  Extensible Markup Language (XML) (experimental)
* [`yaml`](../types/yaml.md):
  YAML Ain't Markup Language (YAML)
* [mxjson](../types/mxjson.md):
  Murex-flavoured JSON (deprecated)