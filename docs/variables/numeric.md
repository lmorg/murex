# Numeric (str)

> Variables who's name is a positive integer, eg `0`, `1`, `2`, `3` and above

## Description

Variables named `0` and above are the equivalent index value of `@ARGV`.

These are reserved variable so they cannot be changed.

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

## Other Reserved Variables

* [Numeric (str)](../variables/numeric.md):
  Variables who's name is a positive integer, eg `0`, `1`, `2`, `3` and above
* [`$.`, Meta Values (json)](../variables/meta-values.md):
  State information for iteration blocks
* [`ARGV` (json)](../variables/argv.md):
  Array of the command name and parameters within a given scope
* [`COLUMNS` (int)](../variables/columns.md):
  Character width of terminal
* [`EVENT_RETURN` (json)](../variables/event_return.md):
  Return values for events
* [`HOSTNAME` (str)](../variables/hostname.md):
  Hostname of the current machine
* [`MUREX_ARGV` (json)](../variables/murex_argv.md):
  Array of the command name and parameters passed to the current shell
* [`MUREX_EXE` (path)](../variables/murex_exe.md):
  Absolute path to running shell
* [`PARAMS` (json)](../variables/params.md):
  Array of the parameters within a given scope
* [`PWDHIST` (json)](../variables/pwdhist.md):
  History of each change to the sessions working directory
* [`PWD` (path)](../variables/pwd.md):
  Current working directory
* [`SELF` (json)](../variables/self.md):
  Meta information about the running scope.
* [`SHELL` (str)](../variables/shell.md):
  Path of current shell

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
* [`string` (stringing)](../types/str.md):
  string (primitive)

<hr/>

This document was generated from [gen/variables/numeric_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/numeric_doc.yaml).