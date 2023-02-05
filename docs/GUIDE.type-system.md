# _murex_ Shell Docs

## Data-Type Reference

This section is a glossary of data-types which _murex_ is natively aware.

Most of the time you will not need to worry about typing in_murex_ as the
shell is designed around productivity as opposed to strictness despite
generally following a strictly typed design.

Read the [Language Tour](GUIDE.quick-start.md) for more detail on this topic.

## Definitions

For clarity, it is worth explaining a couple of terms:

1. "Data-types" in _murex_ are a description of the format of data. This
means that while any stdio stream in UNIX will by "bytes", _murex_ might
label that data as being a JSON string or CSV file (for example) which
means any builtins that parse that stdio stream, for example to return
the first 8 items, would need to parse those types differently. Thus a
"data-type" in _murex_ is actually more than just a description of a data
structure; it is a series of APIs to marshall and unmarshall data from
complex file formats. This enables you to use the same command line tools
to query any type of output.

2. "Primitive" data-types refer to types that are the required by _murex_
to function. These will be `int`, `float` / `number`, `bool`, `string`,
`generic`, `json`, and `null`. All other data-types are optional albeit
still recommended (unless described otherwise).

## Feature Sets

Since not all data formats are equal (for example the TOML file format
doesn't support naked arrays where as JSON does), you may find some
features missing in some data-types which are present in others. If in
doubt then refer to the manual here or check the API manual for more
details about specific hooks.

## Pages

* [`*` (generic) ](types/generic.md):
  generic (primitive)
* [`bool` ](types/bool.md):
  Boolean (primitive)
* [`commonlog` ](types/commonlog.md):
  Apache httpd "common" log format
* [`csv` ](types/csv.md):
  CSV files (and other character delimited tables)
* [`float` (floating point number)](types/float.md):
  Floating point number (primitive)
* [`hcl` ](types/hcl.md):
  HashiCorp Configuration Language (HCL)
* [`int` ](types/int.md):
  Whole number (primitive)
* [`json` ](types/json.md):
  JavaScript Object Notation (JSON) (primitive)
* [`jsonc` ](types/jsonc.md):
  Concatenated JSON
* [`jsonl` ](types/jsonl.md):
  JSON Lines (primitive)
* [`num` (number)](types/num.md):
  Floating point number (primitive)
* [`path` (string) ](types/path.md):
  path data type
* [`str` (string) ](types/str.md):
  string (primitive)
* [`toml` ](types/toml.md):
  Tom's Obvious, Minimal Language (TOML)
* [`yaml` ](types/yaml.md):
  YAML Ain't Markup Language (YAML)
* [mxjson](types/mxjson.md):
  Murex-flavoured JSON (primitive)