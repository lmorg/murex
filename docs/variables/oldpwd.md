# `OLDPWD` (path)

> Return the home directory for the current session user

## Description

`OLDPWD` return the previous directory.

This variable exists to support POSIX, however the idiomatic way to access this
same data is via `$PWDHIST`.

This is a [reserved variable](/docs/user-guide/reserved-vars.md) so it cannot be changed.

## Detail

### Comparison With PWDHIST

`PWDHIST` is an array that holds the entire `PWD` history rather than just the
previously accessed path.

`OLDPWD` reads `PWDHIST`, so if `PWDHIST` is overwritten, this will affect the
value of `OLDPWD` as well.

### Error Handling

If a previous directory cannot be determined, then `OLDPWD` will error. For
example:

```
» cd $OLDPWD
Error in `cd` (0,1): cannot assign value to $OLDPWD: already at oldest entry in $PWDHIST
                   > Expression: cd $OLDPWD
                   >           :          ^
                   > Character : 9
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
* [`COLUMNS` (int)](../variables/lines.md):
  Character height of terminal
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
* [`RANDOM` (int)](../variables/random.md):
  Return a random 32-bit integer (historical)
* [`SELF` (json)](../variables/self.md):
  Meta information about the running scope.
* [`SHELL` (str)](../variables/shell.md):
  Path of current shell
* [`TMPDIR` (path)](../variables/tmpdir.md):
  Return the temporary directory
* [`USER` (str)](../variables/user.md):
  Username for the current session

## See Also

* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [`PWDHIST` (json)](../variables/pwdhist.md):
  History of each change to the sessions working directory
* [`PWD` (path)](../variables/pwd.md):
  Current working directory
* [`path`](../types/path.md):
  Structured object for working with file and directory paths

<hr/>

This document was generated from [gen/variables/OLDPWD_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/OLDPWD_doc.yaml).