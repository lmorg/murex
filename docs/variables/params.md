# `PARAMS` (json)

> Array of the parameters within a given scope

## Description

`PARAMS` returns an array of the parameters within a given scope.
eg `function`, `private`, `autocomplete` or shell script.

Unlike `$ARGV`, `$PARAMS` does not include the function name.

This is a [reserved variable](/docs/user-guide/reserved-vars.md) so it cannot be changed.

## Examples

```
» function example { $PARAMS }
» example abc 1 2 3
[
    "abc",
    "1",
    "2",
    "3"
]
```

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
* [`ARGV` (json)](../variables/argv.md):
  Array of the command name and parameters within a given scope
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)
* [`str` (string)](../types/str.md):
  string (primitive)

<hr/>

This document was generated from [gen/variables/PARAMS_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/PARAMS_doc.yaml).