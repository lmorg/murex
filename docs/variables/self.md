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

Please see scoping document (link below) for more information on scoping.

This is a reserved variable so it cannot be changed.

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

## Other Reserved Variables

* [Numeric (str)](../variables/numeric.md):
  Variables who's name is a positive integer, eg `0`, `1`, `2`, `3` and above
* [`$.`, Meta Values (json)](../variables/meta-values.md):
  State information for iteration blocks
* [`ARGV` (json)](../variables/argv.md):
  Array of the command name and parameters within a given scope
* [`COLUMNS` (int)](../variables/columns.md):
  Character width of terminal
* [`EVENT_RETURN` (json)](../variables/event_return.md):
  Return values for events
* [`HOSTNAME` (str)](../variables/hostname.md):
  Hostname of the current machine
* [`MUREX_ARGV` (json)](../variables/murex_argv.md):
  Array of the command name and parameters passed to the current shell
* [`MUREX_EXE` (path)](../variables/murex_exe.md):
  Absolute path to running shell
* [`PARAMS` (json)](../variables/params.md):
  Array of the parameters within a given scope
* [`PWDHIST` (json)](../variables/pwdhist.md):
  History of each change to the sessions working directory
* [`PWD` (path)](../variables/pwd.md):
  Current working directory
* [`SELF` (json)](../variables/self.md):
  Meta information about the running scope.
* [`SHELL` (str)](../variables/shell.md):
  Path of current shell

## See Also

* [Interactive Shell](../user-guide/interactive-shell.md):
  What's different about Murex's interactive shell?
* [Modules And Packages](../user-guide/modules.md):
  An introduction to Murex modules and packages
* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Public Function (`function`)](../commands/function.md):
  Define a function block
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [Variable and Config Scoping](../user-guide/scoping.md):
  How scoping works within Murex
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)
* [`string` (stringing)](../types/str.md):
  string (primitive)

<hr/>

This document was generated from [gen/variables/SELF_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/SELF_doc.yaml).