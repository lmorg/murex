# Bang Prefix

> Bang prefixing to reverse default actions

## Description

Some builtins support a bang prefix, `!`, which provides a shorthand negative
action to default behavior. For example, `set` defines a variable where as
`!set` will undefine a variable.

Sometimes the shortcut will be logical, like a **not** operator, as is the case
with `and` where typically each result has to equal **true** normally or
**false** if used in `!and`.

Sometimes the shortcut will be more philosophical, such as with `config` where
normal operations is to query or set configuration but `!config` resets the
configuration to defaults (thus operating the same as `config default`).

Please read the respecting commands doc for details on whether it supports a
bang prefix and what the behavior of that prefix is.

## See Also

* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [and](../user-guide/and.md):
  
* [config](../user-guide/config.md):
  
* [not](../user-guide/bang-prefix.md):
  Bang prefixing to reverse default actions
* [or](../user-guide/or.md):
  
* [set](../user-guide/set.md):
  

<hr/>

This document was generated from [gen/user-guide/bang-prefix_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/user-guide/bang-prefix_doc.yaml).