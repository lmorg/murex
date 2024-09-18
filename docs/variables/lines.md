# `LINES` (int)

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