# `COLUMNS` (int)

> Character width of terminal

## Description

`COLUMNS` returns the cell width of the terminal.

Some characters might be more than or less than 1 (one) cell in width, such
as Chinese logograms and zero-width joiners. Whereas one ASCII character is
the same width as one terminal cell.

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
* [Interactive Shell](../user-guide/interactive-shell.md):
  What's different about Murex's interactive shell?
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [`str` (string)](../types/str.md):
  string (primitive)
* [`string` (stringing)](../types/str.md):
  string (primitive)

<hr/>

This document was generated from [gen/variables/COLUMNS_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/COLUMNS_doc.yaml).