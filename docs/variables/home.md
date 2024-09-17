# `HOME` (path)

> Return the home directory for the current session user

## Description

`$HOME` returns the home directory for the current session user.

This variable duplicates functionality from `~` and thus is only provided for
POSIX support.

This is a [reserved variable](/docs/user-guide/reserved-vars.md) so it cannot be changed.

## Other Reserved Variables

* [Numeric (str)](../variables/numeric.md):
  Variables who's name is a positive integer, eg `0`, `1`, `2`, `3` and above
* [`$.`, Meta Values (json)](../variables/meta-values.md):
  State information for iteration blocks
* [`ARGV` (json)](../variables/argv.md):
  Array of the command name and parameters within a given scope
* [`COLUMNS` (int)](../variables/lines.md):
  Character height of terminal
* [`COLUMNS` (int)](../variables/columns.md):
  Character width of terminal
* [`EVENT_RETURN` (json)](../variables/event_return.md):
  Return values for events
* [`HOME` (path)](../variables/home.md):
  Return the home directory for the current session user
* [`HOSTNAME` (str)](../variables/hostname.md):
  Hostname of the current machine
* [`LOGNAME` (str)](../variables/logname.md):
  Username for the current session (historic)
* [`MUREX_ARGV` (json)](../variables/murex_argv.md):
  Array of the command name and parameters passed to the current shell
* [`MUREX_EXE` (path)](../variables/murex_exe.md):
  Absolute path to running shell
* [`OLDPWD` (path)](../variables/oldpwd.md):
  Return the home directory for the current session user
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
* [`USER` (str)](../variables/user.md):
  Username for the current session

## See Also

* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [`path`](../types/path.md):
  Structured object for working with file and directory paths
* [`~` Home Sigil](../parser/tilde.md):
  Home directory path variable

<hr/>

This document was generated from [gen/variables/HOME_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/HOME_doc.yaml).