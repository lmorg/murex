# `MUREX_EXE` (path)

> Absolute path to running shell

## Description

`MUREX_EXE` is very similar to the `$SHELL` environmental variable in that it
holds the full path to the running shell. The reason for defining a reserved
variable is so that the shell path cannot be overridden.

This is a reserved variable so it cannot be changed.

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
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [`SHELL` (str)](../variables/shell.md):
  Path of current shell
* [`path`](../types/path.md):
  Structured object for working with file and directory paths

<hr/>

This document was generated from [gen/variables/MUREX_EXE_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/MUREX_EXE_doc.yaml).