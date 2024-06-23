# `PWD` (path)

> Current working directory

## Description

`PWD` is an environmental variable containing the current working directory.

It is updated via `cd` however you can overwrite its value manually via `export`.

## Other Reserved Variables

* [Numeric (str)](../variables/numeric.md):
  Variables who's name is a positive integer, eg `0`, `1`, `2`, `3` and above
* [`$.`, Meta Values (json)](../variables/meta-values.md):
  State information for iteration blocks
* [`ARGV` (json)](../variables/argv.md):
  Array of the command name and parameters within a given scope
* [`COLUMNS` (int)](../variables/columns.md):
  Character width of terminal
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

* [`PWDHIST` (json)](../variables/pwdhist.md):
  History of each change to the sessions working directory
* [`cd`](../commands/cd.md):
  Change (working) directory
* [`export`](../commands/export.md):
  Define an environmental variable and set it's value
* [`path`](../types/path.md):
  Structured object for working with file and directory paths

<hr/>

This document was generated from [gen/variables/PWD_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/PWD_doc.yaml).