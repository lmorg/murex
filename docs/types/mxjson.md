# _murex_ Shell Docs

## Data-Type Reference: mxjson

> Murex-flavoured JSON (primitive)

## Description

mxjson is an extension to JSON designed to integrate more seamlessly when
use as a configuration file. Thus mxjson supports comments and _murex_ code
blocks embedded into the JSON schema.

> mxjson is a format that is pre-parsed into a valid JSON format.

mxjson isn't currently a proper _murex_ data-type in that you cannot marshal
and unmarshal mxjson files. Currently it is a format that is only supported
by a small subset of _murex_ builtins (eg `config` and `autocomplete`) where
config might embed _murex_ code blocks.

**mxjson features the following enhancements:**

### Line Comments

Line comments are prefixed with a 'hash', `#`, just like with regular _murex_
code.

### Block Quotation

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

## See Also

* [parser/Brace Quote (`(`, `)`) Tokens](../parser/brace-quote.md):
  Initiates or terminates a string (variables expanded)
* [parser/Curly Brace (`{`, `}`) Tokens](../parser/curly-brace.md):
  Initiates or terminates a code block
* [apis/`Marshal()` (type)](../apis/Marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [apis/`Unmarshal()` (type)](../apis/Unmarshal.md):
  Converts a structured file format into structured memory
* [commands/`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [commands/`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [commands/`autocomplete`](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [commands/`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [commands/`config`](../commands/config.md):
  Query or define _murex_ runtime settings
* [commands/`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [types/`hcl` ](../types/hcl.md):
  HashiCorp Configuration Language (HCL)
* [types/`json` ](../types/json.md):
  JavaScript Object Notation (JSON) (primitive)
* [types/`jsonc` ](../types/jsonc.md):
  JSON Lines (primitive)
* [types/`jsonl` ](../types/jsonl.md):
  JSON Lines (primitive)
* [commands/`open`](../commands/open.md):
  Open a file with a preferred handler
* [commands/`pretty`](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [commands/`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [types/`toml` ](../types/toml.md):
  Tom's Obvious, Minimal Language (TOML)
* [types/`yaml` ](../types/yaml.md):
  YAML Ain't Markup Language (YAML)
* [parser/code-block](../parser/code-block.md):
  