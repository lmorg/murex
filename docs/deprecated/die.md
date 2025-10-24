# Exit after error: `die`

**This feature has been deprecated** and thus the following documentation is
provided for historical reference rather than recommendations for new code.

While we do make every effort to maintain backwards compatibility, sometimes
deprecations need to be made in order to keep Murex focused and maintainable.

Please read our [compatibility commitment](https://murex.rocks/compatibility.html)
for more information on how we approach such changes. Visit the [deprecated section](https://github.com/lmorg/murex/tree/master/docs/deprecated)
if you need to view other deprecated features.


> Terminate murex with an exit number of 1 (removed 7.0)

## Description

Terminate Murex with an exit number of 1.

> This builtin has now been deprecated. The same behaviour can be achieved via
> `exit 1`

## Usage

```
die
```

## Examples

```
» die
```

## See Also

* [Exit Murex: `exit`](../commands/exit.md):
  Exit murex
* [Exit Scope: `break`](../commands/break.md):
  Terminate execution of a block within your processes scope
* [Null: `null`](../commands/devnull.md):
  null function. Similar to /dev/null

<hr/>

This document was generated from [builtins/core/typemgmt/types_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/typemgmt/types_doc.yaml).
