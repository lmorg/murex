# `ARGV` (json)

> Array of the command name and parameters within a given scope

## Description

`ARGV` returns an array of the command name and parameters within a given
scope. eg `function`, `private`, `autocomplete` or shell script.

Unlike `$PARAMS`, `$ARGV` includes the function name.

This is a [reserved variable](/docs/user-guide/reserved-vars.md) so it cannot be changed.

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

## See Also

* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Modules And Packages](../user-guide/modules.md):
  Modules and packages: An Introduction
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Private Function (`private`)](../commands/private.md):
  Define a private function block
* [Public Function (`function`)](../commands/function.md):
  Define a function block
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [Tab Autocompletion (`autocomplete`)](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [Variable And Config Scoping](../user-guide/scoping.md):
  How scoping works within Murex
* [`@Array` Sigil](../parser/array.md):
  Expand values as an array
* [`PARAMS` (json)](../variables/params.md):
  Array of the parameters within a given scope
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)
* [`str` (string)](../types/str.md):
  string (primitive)

<hr/>

This document was generated from [gen/variables/ARGV_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/ARGV_doc.yaml).