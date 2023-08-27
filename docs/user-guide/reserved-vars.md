# Reserved Variables

> Special variables reserved by Murex

In Murex, there are four different classes of variables:
1. Local variables (scope limited)
2. Global variables (available to every function within Murex but not shared
    with processes outside of the Murex's runtime)
3. Environmental variables (available to every function and process -- internal
    and external to Murex)
4. Reserved variables

Reserved variables are data that are available to any code running within
Murex and exposed as a variable.

Because reserved variables are dynamic properties of the runtime environment,
they can only be queried and not set:
```
» set SELF="foobar"
Error in `set` (0,1): cannot set a reserved variable: SELF
```

## See Also

* [Modules and Packages](../user-guide/modules.md):
  An introduction to Murex modules and packages
* [Numeric (str)](../variables/numeric.md):
  Variables who's name is a positive integer, eg `0`, `1`, `2`, `3` and above
* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Variable and Config Scoping](../user-guide/scoping.md):
  How scoping works within Murex
* [`ARGV` (json)](../variables/argv.md):
  Array of the command name and parameters within a given scope
* [`COLUMNS` (int)](../variables/columns.md):
  Character width of terminal
* [`HOSTNAME` (str)](../variables/hostname.md):
  Hostname of the current machine
* [`MUREX_ARGV` (json)](../variables/murex_argv.md):
  Array of the command name and parameters passed to the current shell
* [`SELF` (json)](../variables/self.md):
  Meta information about the running scope.
* [`SHELL` (str)](../variables/shell.md):
  Path of current shell
* [`export`](../commands/export.md):
  Define an environmental variable and set it's value
* [`global`](../commands/global.md):
  Define a global variable and set it's value
* [`set`](../commands/set.md):
  Define a local variable and set it's value

<hr/>

This document was generated from [gen/user-guide/reserved_vars_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/user-guide/reserved_vars_doc.yaml).