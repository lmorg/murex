# `SHELL` (str)

> Path of current shell

## Description

`SHELL` is an environmental variable containing the full path to the current
shell -- which in our case is Murex.

This variable is set by Murex to conform to POSIX standards. However, being an
environmental variable, it can be overwritten via `export`.

For Murex specific code, you are recommended to use `MUREX_EXE`, which is a
reserved variable, and thus read only, if you require precision and safety.

## See Also

* [Define Environmental Variable (`export`)](../commands/export.md):
  Define an environmental variable and set it's value
* [`MUREX_EXE` (path)](../variables/murex_exe.md):
  Absolute path to running shell
* [`str` (string)](../types/str.md):
  string (primitive)

<hr/>

This document was generated from [gen/variables/SHELL_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/SHELL_doc.yaml).