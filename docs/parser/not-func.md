# `!` (not)

> Reads the STDIN and exit number from previous process and not's it's condition

## Description

Reads the STDIN and exit number from previous process and not's it's condition.

## Usage

```
<stdin> -> ! -> <stdout>
```## Examples

```
» echo "Hello, world!" -> !
false
```

```
» false -> !
true
```

## Synonyms

* `!`
* `not`


## See Also

* [`and`](../commands/and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
* [`false`](../commands/false.md):
  Returns a `false` value
* [`if`](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [`or`](../commands/or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.
* [`true`](../commands/true.md):
  Returns a `true` value

<hr/>

This document was generated from [builtins/core/typemgmt/types_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/typemgmt/types_doc.yaml).