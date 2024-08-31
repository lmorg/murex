# Reserved Variables

> Special variables reserved by Murex

In Murex, there are five different classes of variables:
1. Local variables (scope limited to a function et al)
2. Module variables (scoped to a module)
3. Global variables (available to every function within Murex but not shared
    with processes outside of the Murex's runtime)
4. Environmental variables (available to every function and process -- internal
    and external to Murex)
5. Reserved variables

Reserved variables are data that are available to any code running within
Murex and exposed as a variable.

Because reserved variables are dynamic properties of the runtime environment,
they can only be queried and not set:

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
  An introduction to Murex modules and packages
* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Variable and Config Scoping](../user-guide/scoping.md):
  How scoping works within Murex

<hr/>

This document was generated from [gen/user-guide/reserved_vars_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/user-guide/reserved_vars_doc.yaml).