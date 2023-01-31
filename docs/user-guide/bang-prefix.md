# User Guide: Bang Prefix

> Bang prefixing to reverse default actions

Some builtins support a bang prefix, `!`, which provides a shorthand negative
action to default behavior. For example, `set` defines a variable where as
`!set` will undefine a variable.

Sometimes the shortcut will be logical, like a **not** operator, as is the case
with `and` where typically each result has to equal **true** normally or
**false** if used in `!and`. Sometimes the shortcut will be more philosophical,
such as with `config` where normal operations is to query or set configuration
but `!config` resets the configuration to defaults (thus operating the same as
`config default`).

Please read the respecting commands doc for details on whether it supports a
bang prefix and what the behavior of that prefix is.

## See Also

* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by _murex_
* [`!` (not)](../commands/not.md):
  Reads the STDIN and exit number from previous process and not's it's condition
* [`and`](../commands/and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
* [`config`](../commands/config.md):
  Query or define _murex_ runtime settings
* [`or`](../commands/or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.
* [`set`](../commands/set.md):
  Define a local variable and set it's value