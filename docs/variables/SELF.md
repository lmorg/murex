# `SELF` (json)

> Meta information about the running scope.

## Description

`SELF` returns information about the functional scope that the code is running
inside. Such as whether that functions STDOUT is a TTY, running in the
background or a method.

A 'scope' in Murex is a collection of code blocks to which variables and
config are persistent within. In Murex, a variable declared inside an `if` or
`foreach` block will be persistent outside of their blocks as long as you're
still inside the same function.

Please see scoping document (link below) for more information on scoping.

This is a reserved variable so it cannot be changed.



## Examples

```
» function example { out $SELF }
» example
{
    "Background": false,
    "Interactive": true,
    "Method": false,
    "Module": "murex/shell",
    "Not": false,
    "Parent": 834,
    "Scope": 834,
    "TTY": true
}
```

## See Also

* [Interactive Shell](../user-guide/interactive-shell.md):
  What's different about Murex's interactive shell?
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
* [`function`](../commands/function.md):
  Define a function block
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)