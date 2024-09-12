# `$.`, Meta Values (json)

> State information for iteration blocks

## Description

Meta Values, `$.`, provides state information for blocks like `foreach`,
`formap`, `while` and lambdas.

Meta Values are a specific to the block, so you will need to refer to each
iteration structure's documentation to check what information is exposed via
`$.`

## Examples

```
» %[Monday..Friday] -> foreach day { out "$.i: $day" }
1: Monday
2: Tuesday
3: Wednesday
4: Thursday
5: Friday
```

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

* [For Each In List (`foreach`)](../commands/foreach.md):
  Iterate through an array
* [For Each In Map (`formap`)](../commands/formap.md):
  Iterate through a map or other collection of data
* [Loop While (`while`)](../commands/while.md):
  Loop until condition false
* [`[{ Lambda }]`](../parser/lambda.md):
  Iterate through structured data

<hr/>

This document was generated from [gen/variables/meta-values_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/meta-values_doc.yaml).