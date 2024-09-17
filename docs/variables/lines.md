# `COLUMNS` (int)

> Character height of terminal

## Description

`LINES` returns the cell height of the terminal.

This is a [reserved variable](/docs/user-guide/reserved-vars.md) so it cannot be changed.

## Detail

The Murex controlling terminal is assumed to be the terminal. If Stdout cannot
be successfully queried for its dimensions, for example Murex is a piped rather
than controlling a TTY, then `$LINES` will generate an error:

```
» exec $MUREX_EXE -c 'out $LINES' -> cat
Error in `out` ( 1,1):
      Command: out $LINES
      Error: cannot assign value to $LINES: inappropriate ioctl for device
          > Expression: out $LINES
          >           :          ^
          > Character : 9
Error in `/Users/laurencemorgan/dev/go/src/github.com/lmorg/murex/murex` (0,1): exit status 1
```

This error can be caught via `||`, `try` et al.

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

* [Execute External Command (`exec`)](../commands/exec.md):
  Runs an executable
* [Interactive Shell](../user-guide/interactive-shell.md):
  What's different about Murex's interactive shell?
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [Try Block (`try`)](../commands/try.md):
  Handles non-zero exits inside a block of code
* [`COLUMNS` (int)](../variables/columns.md):
  Character width of terminal
* [`MUREX_EXE` (path)](../variables/murex_exe.md):
  Absolute path to running shell
* [`int`](../types/int.md):
  Whole number (primitive)
* [`||` Or Logical Operator](../parser/logical-or.md):
  Continues next operation only if previous operation fails

<hr/>

This document was generated from [gen/variables/LINES_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/LINES_doc.yaml).