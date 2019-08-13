# _murex_ Language Guide

## Data-Type Reference

This section is a glossary of data-types which _murex_ is natively aware.

### Definitions

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

### Feature Sets

Since not all data formats are equal (for example the TOML file format
doesn't support naked arrays where as JSON does), you may find some
features missing in some data-types which are present in others. If in
doubt then refer to the manual here or check the API manual for more
details about specific hooks.

### Pages

* [`float` (floating point number)](types/float.md):
  Floating point number (primitive)
* [`hcl` (HCL)](types/hcl.md):
  HashiCorp Configuration Language (HCL)
* [`int` (integer)](types/int.md):
  Whole number (primitive)
* [`json` (JSON)](types/json.md):
  JavaScript Object Notation (JSON) (primitive)
* [`num` (number)](types/num.md):
  Floating point number (primitive)
* [`toml` (TOML)](types/toml.md):
  Tom's Obvious, Minimal Language (TOML)
* [`yaml` (YAML)](types/yaml.md):
  YAML Ain't Markup Language (YAML)
* [mxjson](types/mxjson.md):
  Murex-flavoured JSON (primitive)