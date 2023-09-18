# mxjson

> Murex-flavoured JSON (deprecated)

## Description

> This format has been deprecated in favour of `%{}` constructors.

mxjson is an extension to JSON designed to integrate more seamlessly when
use as a configuration file. Thus mxjson supports comments and Murex code
blocks embedded into the JSON schema.

> mxjson is a format that is pre-parsed into a valid JSON format.

mxjson isn't a Murex data-type in that you cannot marshal
and unmarshal mxjson files. Currently it is a format that is only supported
by a small subset of Murex builtins (eg `config` and `autocomplete`) where
config might embed Murex code blocks.

**mxjson features the following enhancements:**

### Line Comments

Line comments are prefixed with a 'hash', `#`, just like with regular Murex
code.

### Block Quotation

Code blocks are quoted with `(`, `)`. For example, below "ExampleFunction"
uses the `({ block quote })` method.

```
{
    "ExampleFunction": ({
        out "This is an example Murex function"
        if { =1==2 } then {
            err "The laws of the universe are broken"
        }
    })
}
```

Any block quoted by this method will be converted to the following valid JSON:

```
{
    "ExampleFunction": "\n    out \"This is an example Murex function\"\n    if { =1==2 } then {\n        err \"The laws of the universe are broken\"\n    }"
}
```

## See Also

* [`hcl`](../types/hcl.md):
  HashiCorp Configuration Language (HCL)
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)
* [`jsonc`](../types/jsonc.md):
  Concatenated JSON
* [`jsonl`](../types/jsonl.md):
  JSON Lines
* [`toml`](../types/toml.md):
  Tom's Obvious, Minimal Language (TOML)
* [`yaml`](../types/yaml.md):
  YAML Ain't Markup Language (YAML)
* [autocomplete](../types/autocomplete.md):
  
* [brace-quote](../types/brace-quote.md):
  
* [cast](../types/cast.md):
  
* [code-block](../types/code-block.md):
  
* [config](../types/config.md):
  
* [create-array](../types/create-array.md):
  
* [create-object](../types/create-object.md):
  
* [curly-brace](../types/curly-brace.md):
  
* [element](../types/element.md):
  
* [format](../types/format.md):
  
* [index](../types/index.md):
  
* [open](../types/open.md):
  
* [pretty](../types/pretty.md):
  
* [runtime](../types/runtime.md):
  

### Read more about type hooks

- [`ReadIndex()` (type)](../apis/ReadIndex.md): Data type handler for the index, `[`, builtin
- [`ReadNotIndex()` (type)](../apis/ReadNotIndex.md): Data type handler for the bang-prefixed index, `![`, builtin
- [`ReadArray()` (type)](../apis/ReadArray.md): Read from a data type one array element at a time
- [`WriteArray()` (type)](../apis/WriteArray.md): Write a data type, one array element at a time
- [`ReadMap()` (type)](../apis/ReadMap.md): Treat data type as a key/value structure and read its contents
- [`Marshal()` (type)](../apis/Marshal.md): Converts structured memory into a structured file format (eg for stdio)
- [`Unmarshal()` (type)](../apis/Unmarshal.md): Converts a structured file format into structured memory

<hr/>

This document was generated from [builtins/types/json/mxjson_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/types/json/mxjson_doc.yaml).