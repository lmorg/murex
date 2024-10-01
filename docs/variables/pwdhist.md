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

## See Also

* [Change Directory (`cd`)](../commands/cd.md):
  Change (working) directory
* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Modules And Packages](../user-guide/modules.md):
  Modules and packages: An Introduction
* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [Variable And Config Scoping](../user-guide/scoping.md):
  How scoping works within Murex
* [`@Array` Sigil](../parser/array.md):
  Expand values as an array
* [`OLDPWD` (path)](../variables/oldpwd.md):
  Return the home directory for the current session user
* [`PWD` (path)](../variables/pwd.md):
  Current working directory
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)
* [`path`](../types/path.md):
  Structured object for working with file and directory paths

<hr/>

This document was generated from [gen/variables/PWDHIST_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/PWDHIST_doc.yaml).