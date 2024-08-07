# `PARAMS` (json)

> Array of the parameters within a given scope

## Description

`PARAMS` returns an array of the parameters within a given scope.
eg `function`, `private`, `autocomplete` or shell script.

Unlike `$ARGV`, `$PARAMS` does not include the function name.

This is a reserved variable so it cannot be changed.

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

* [Modules And Packages](../user-guide/modules.md):
  An introduction to Murex modules and packages
* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [Variable and Config Scoping](../user-guide/scoping.md):
  How scoping works within Murex
* [`@Array` Sigil](../parser/array.md):
  Expand values as an array
* [`ARGV` (json)](../variables/argv.md):
  Array of the command name and parameters within a given scope
* [`autocomplete`](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [`function`](../commands/function.md):
  Define a function block
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)
* [`out`](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [`private`](../commands/private.md):
  Define a private function block
* [`set`](../commands/set.md):
  Define a local variable and set it's value
* [`string` (stringing)](../types/str.md):
  string (primitive)

<hr/>

This document was generated from [gen/variables/PARAMS_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/PARAMS_doc.yaml).