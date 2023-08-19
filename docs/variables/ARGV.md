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

* `MUREX_ARGV`
* `MUREX_ARGS`


## See Also

* [Array (`@`) Token](../parser/array.md):
  Expand values as an array
* [`ARGV` (json)](../variables/ARGV.md):
  Array of the command name and parameters within a given scope
* [`MUREX_EXE` (path)](../variables/MUREX_EXE.md):
  Absolute path to running shell
* [`PARAMS` (json)](../variables/PARAMS.md):
  Array of the parameters within a given scope
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)