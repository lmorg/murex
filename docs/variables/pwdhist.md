# `PWDHIST` (json)

> History of each change to the sessions working directory

## Description

`PWDHIST` is a JSON array containing the history of all the working directories
within the current shell session.

It is updated via `cd` however you can overwrite its value manually via `set`.

## Examples

```
» cd ~bob
» cd /tmp
» $PWDHIST
[
    "/Users/bob",
    "/private/tmp"
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

* [Change Directory (`cd`)](../commands/cd.md):
  Change (working) directory
* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
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
* [`PWD` (path)](../variables/pwd.md):
  Current working directory
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)
* [`path`](../types/path.md):
  Structured object for working with file and directory paths
* [`string` (stringing)](../types/str.md):
  string (primitive)

<hr/>

This document was generated from [gen/variables/PWDHIST_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/PWDHIST_doc.yaml).