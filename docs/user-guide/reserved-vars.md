# Reserved Variables

> Special variables reserved by Murex

In Murex, there are different classes of variables:

- Murex variables, of varying scope:
  1. Local variables (scope limited to a function et al)
  2. Module variables (scoped to a module)
  3. Global variables (available to every function within Murex but not shared
    with processes outside of the Murex's runtime)

- Environmental variables (available to every function and process -- internal
    and external to Murex)

- Reserved variables

Reserved variables are Murex variables which contain read-only runtime data and
thus made available via Murex's runtime rather than assigned by any running
Murex code.

Reserved variables are called _reserved_ because they are read only. 

Reserved variables are often also dynamic, returning different values based on
contextual circumstances.

```
» set SELF="foobar"
Error in `set` (0,1): cannot set a reserved variable: SELF
```

## See Also

* [Define Environmental Variable (`export`)](../commands/export.md):
  Define an environmental variable and set it's value
* [Define Global (`global`)](../commands/global.md):
  Define a global variable and set it's value
* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Modules And Packages](../user-guide/modules.md):
  Modules and packages: An Introduction
* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Variable And Config Scoping](../user-guide/scoping.md):
  How scoping works within Murex

<hr/>

This document was generated from [gen/user-guide/reserved_vars_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/user-guide/reserved_vars_doc.yaml).