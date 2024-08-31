# Exit Murex (`exit`)

> Exit murex

## Description

Exit's Murex with either a exit number of 0 (by default if no parameters
supplied) or a custom value specified by the first parameter.

`exit` is not scope aware; if it is included in a function then the whole
shell will still exist and not just that function.

## Usage

```
exit

exit number
```

## Examples

```
» exit
```

```
» exit 42
```

## See Also

* [Exit Block (`break`)](../commands/break.md):
  Terminate execution of a block within your processes scope
* [Null (`null`)](../commands/devnull.md):
  null function. Similar to /dev/null
* [`die`](../commands/die.md):
  Terminate murex with an exit number of 1 (deprecated)

<hr/>

This document was generated from [builtins/core/typemgmt/types_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/typemgmt/types_doc.yaml).