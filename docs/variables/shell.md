# `SHELL` (str)

> Path of current shell

## Description

`SHELL` is an environmental variable containing the full path to the current
shell -- which in our case is Murex.

This variable is set by Murex to conform to standards however, being an
environmental variable, it can be overwritten via `export`. Thus you are
recommended to use `MUREX_EXE`, which is a reserved variable, if you require
precision.



## Synonyms

* `shell`
* `SHELL`


## See Also

* [`MUREX_EXE` (path)](../variables/murex_exe.md):
  Absolute path to running shell
* [`export`](../commands/export.md):
  Define an environmental variable and set it's value
* [`string` (stringing)](../types/str.md):
  string (primitive)

<hr/>

This document was generated from [gen/variables/SHELL_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/SHELL_doc.yaml).