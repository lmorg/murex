# `SELF` (json)

> Meta information about the running scope.

## Description

`SELF` returns information about the functional scope that the code is running
inside. Such as whether that functions stdout is a TTY, running in the
background or a method.

A 'scope' in Murex is a collection of code blocks to which variables and
config are persistent within. In Murex, a variable declared inside an `if` or
`foreach` block will be persistent outside of their blocks as long as you're
still inside the same function.

Please read the [scoping document](/docs/user-guide/scoping.md) for more information on scoping.

This is a [reserved variable](/docs/user-guide/reserved-vars.md) so it cannot be changed.

## Examples

```
» function example { $SELF }
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

## Detail

### Background (bool)

A boolean value to identify whether the current scope is running in the
background for foreground.

### Interactive (bool)

A boolean value to describe whether the current scope is running interactively
or not.

An interactive scope is one where the shell prompt is running _and_ the scope
isn't running in the background. Shell scripts are not considered interactive
terminals even though they might have interactive element in their code.

### Method (bool)

A boolean value to describe whether the current scope is a method (ie being
called mid-way or at the end of a pipeline).

### Module (str)

This will be the module string for the current scope.

### Not (bool)

A boolean value which represents whether the function was called with a bang-
prefix or not.

### Parent (num)

This is the function ID of the parent function that created the scope. In
some instances this will be the same value as scope FID. However if in doubt
then please using **Scope** instead.

### Scope (num)

The scope value here returns the function ID of the top level function in the
scope.

### TTY (bool)

A boolean value as to whether stdout is a TTY (ie are we printing to the
terminal (TTY) or a pipe?)

## See Also

* [Interactive Shell](../user-guide/interactive-shell.md):
  What's different about Murex's interactive shell?
* [Modules And Packages](../user-guide/modules.md):
  Modules and packages: An Introduction
* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Public Function (`function`)](../commands/function.md):
  Define a function block
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [Variable And Config Scoping](../user-guide/scoping.md):
  How scoping works within Murex
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)
* [`str` (string)](../types/str.md):
  string (primitive)

<hr/>

This document was generated from [gen/variables/SELF_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/SELF_doc.yaml).