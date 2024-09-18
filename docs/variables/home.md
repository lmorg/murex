# `HOME` (path)

> Return the home directory for the current session user

## Description

`$HOME` returns the home directory for the current session user.

This variable duplicates functionality from `~` and thus is only provided for
POSIX support.

This is a [reserved variable](/docs/user-guide/reserved-vars.md) so it cannot be changed.

## See Also

* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [`path`](../types/path.md):
  Structured object for working with file and directory paths
* [`~` Home Sigil](../parser/tilde.md):
  Home directory path variable

<hr/>

This document was generated from [gen/variables/HOME_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/HOME_doc.yaml).