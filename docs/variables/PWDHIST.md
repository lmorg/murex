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

* [Array (`@`) Token](../parser/array.md):
  Expand values as an array
* [Modules and Packages](../user-guide/modules.md):
  An introduction to Murex modules and packages
* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [String (`$`) Token](../parser/string.md):
  Expand values as a string
* [Variable and Config Scoping](../user-guide/scoping.md):
  How scoping works within Murex
* [`cd`](../commands/cd.md):
  Change (working) directory
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)
* [`path`](../types/path.md):
  Structured object for working with file and directory paths
* [`set`](../commands/set.md):
  Define a local variable and set it's value

<hr/>

This document was generated from [gen/variables/PWDHIST_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/PWDHIST_doc.yaml).