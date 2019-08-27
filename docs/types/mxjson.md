# _murex_ Shell Guide

## Data-Type Reference: mxjson

> Murex-flavoured JSON (primitive)

### Description

mxjson is an extension to JSON designed to integrate more seamlessly when
use as a configuration file. Thus mxjson supports comments and _murex_ code
blocks embedded into the JSON schema.

> mxjson is a format that is pre-parsed into a valid JSON format.

mxjson isn't currently a proper _murex_ data-type in that you cannot marshal
and unmarshal mxjson files. Currently it is a format that is only supported
by a small subset of _murex_ builtins (eg `config` and `autocomplete`) where
config might embed _murex_ code blocks.

### mxjson Features The Following Enhancements

#### Line Comments

Line comments are prefixed with a 'hash', `#`, just like with regular _murex_
code.

#### Block Quotation

Code blocks are quoted with `(`, `)`. For example, below "ExampleFunction"
uses the `({ block quote })` method.

    {
        "ExampleFunction": ({
            out: "This is an example _murex_ function"
            if { =1==2 } then {
                err: "The laws of the universe are broken"
            }
        })
    }
    
Any block quoted by this method will be converted to the following valid JSON:

    {
        "ExampleFunction": "\n    out: \"This is an example _murex_ function\"\n    if { =1==2 } then {\n        err: \"The laws of the universe are broken\"\n    }"
    }



### Default Associations





### See Also

* [`Marshal()` ](../apis/marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [`Unmarshal()` ](../apis/unmarshal.md):
  Converts a structured file format into structured memory
* [`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [`autocomplete`](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`config`](../commands/config.md):
  Query or define _murex_ runtime settings
* [`hcl` (HCL)](../types/hcl.md):
  HashiCorp Configuration Language (HCL)
* [`json` (JSON)](../types/json.md):
  JavaScript Object Notation (JSON) (primitive)
* [`pretty`](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [`toml` (TOML)](../types/toml.md):
  Tom's Obvious, Minimal Language (TOML)
* [`yaml` (YAML)](../types/yaml.md):
  YAML Ain't Markup Language (YAML)
* [element](../commands/element.md):
  
* [format](../commands/format.md):
  
* [jsonl](../types/jsonl.md):
  
* [open](../commands/open.md):
  