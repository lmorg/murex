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