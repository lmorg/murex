# Numeric (str)

> Variables who's name is a positive integer, eg `0`, `1`, `2`, `3` and above

## Description

Variables named `0` and above are the equivalent index value of `@ARGV`.

These are [reserved variables](/docs/user-guide/reserved-vars.md) so they cannot be changed.

## Examples

```
» function example { out $0 $2 }
» example 1 2 3
example 2
```

## Detail

### `0` (str)

This returns the name of the executable (like `$ARGS[0]`)

### `1`, `2`, `3`... (str)

This returns parameter _n_ (like `$ARGS[n]`). If there is no parameter _n_
then the variable will not be set thus the upper limit variable is determined
by how many parameters are set. For example if you have 19 parameters passed
then variables `$1` through to `$19` (inclusive) will all be set.

## See Also

* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Private Function (`private`)](../commands/private.md):
  Define a private function block
* [Public Function (`function`)](../commands/function.md):
  Define a function block
* [Tab Autocompletion (`autocomplete`)](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [`ARGV` (json)](../variables/argv.md):
  Array of the command name and parameters within a given scope
* [`PARAMS` (json)](../variables/params.md):
  Array of the parameters within a given scope
* [`str` (string)](../types/str.md):
  string (primitive)

<hr/>

This document was generated from [gen/variables/numeric_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/numeric_doc.yaml).