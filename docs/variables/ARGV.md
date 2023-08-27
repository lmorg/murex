# `ARGV` (json)

> Array of the command name and parameters within a given scope

## Description

`ARGV` returns an array of the command name and parameters within a given
scope. eg `function`, `private`, `autocomplete` or shell script.

Unlike `$PARAMS`, `$ARGV` includes the function name.

This is a reserved variable so it cannot be changed.



## Examples

```
» function example { $ARGV }
» example abc 1 2 3
[
    "example",
    "abc",
    "1",
    "2",
    "3"
]
```

## Detail

### Deprecation of `ARGS`

In Murex versions 4.x and below, this variable was named `ARGS` (with an 'S').
However in Murex 5.x and above it was renamed to `ARGV` (with a 'V') to unify
the name with other languages.

`ARGS` will remain available for compatibility reasons but is considered
deprecated and may be removed from future releases.

## Synonyms

* `argv`
* `ARGV`
* `ARGS`


## See Also

* [Array (`@`) Token](../parser/array.md):
  Expand values as an array
* [Modules and Packages](../user-guide/modules.md):
  An introduction to Murex modules and packages
* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [String (`$`) Token](../parser/string.md):
  Expand values as a string
* [Variable and Config Scoping](../user-guide/scoping.md):
  How scoping works within Murex
* [`PARAMS` (json)](../variables/params.md):
  Array of the parameters within a given scope
* [`autocomplete`](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [`function`](../commands/function.md):
  Define a function block
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)
* [`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [`private`](../commands/private.md):
  Define a private function block
* [`set`](../commands/set.md):
  Define a local variable and set it's value

<hr/>

This document was generated from [gen/variables/ARGV_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/ARGV_doc.yaml).