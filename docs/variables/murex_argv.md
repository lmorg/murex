# `MUREX_ARGV` (json)

> Array of the command name and parameters passed to the current shell

## Description

`MUREX_ARGV` returns an array of the command name and parameters passed to
the current running Murex shell



## Examples

```
» murex -trypipe -c '$MUREX_ARGV'
[
    "murex",
    "-trypipe",
    "-c",
    "$MUREX_ARGV"
]
```

## Synonyms

* `murex_argv`
* `MUREX_ARGV`
* `MUREX_ARGS`


## See Also

* [Array (`@`) Token](../parser/array.md):
  Expand values as an array
* [`ARGV` (json)](../variables/argv.md):
  Array of the command name and parameters within a given scope
* [`MUREX_EXE` (path)](../variables/murex_exe.md):
  Absolute path to running shell
* [`PARAMS` (json)](../variables/params.md):
  Array of the parameters within a given scope
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)

<hr/>

This document was generated from [gen/variables/MUREX_ARGV_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/MUREX_ARGV_doc.yaml).