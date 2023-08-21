# `MUREX_EXE` (path)

> Absolute path to running shell

## Description

`MUREX_EXE` is very similar to the `$SHELL` environmental variable in that it
holds the full path to the running shell. The reason for defining a reserved
variable is so that the shell path cannot be overridden.

This is a reserved variable so it cannot be changed.



## See Also

* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [`SHELL` (str)](../variables/SHELL.md):
  Path of current shell
* [`path`](../types/path.md):
  Structured object for working with file and directory paths
* [`set`](../commands/set.md):
  Define a local variable and set it's value

<hr/>

This document was generated from [gen/variables/MUREX_EXE_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/MUREX_EXE_doc.yaml).