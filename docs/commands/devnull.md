# Null (`null`)

> null function. Similar to /dev/null

## Description

`null` is a function that acts a little like the `null` data type and the
UNIX /dev/null device.

## Usage

```
<stdin> -> null
```

## Examples

```
» out "Hello, world!" -> null
```

## Detail

While this method does exist, a more idiomatic way to suppress stdout is to
use the named pipe property rather than piping to null:

```
» out <null> "Hello, world!"
```

## Synonyms

* `null`


## See Also

* [Exit Block (`break`)](../commands/break.md):
  Terminate execution of a block within your processes scope
* [Exit Murex (`exit`)](../commands/exit.md):
  Exit murex
* [`die`](../commands/die.md):
  Terminate murex with an exit number of 1 (deprecated)

<hr/>

This document was generated from [builtins/core/typemgmt/types_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/typemgmt/types_doc.yaml).