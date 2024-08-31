# False (`false`)

> Returns a `false` value

## Description

Returns a `false` value.

## Usage

```
false -> <stdout>
```

## Examples

### No flags

By default, `false` also outputs the term "false":

```
» false
false
```

### Silent

You can suppress that with the silent flag:

```
» false -s
```

## Flags

* `-s`
    silent - don't output the term "false"

## See Also

* [If Conditional (`if`)](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [Logic And Statements (`and`)](../commands/and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
* [Logic Or Statements (`or`)](../commands/or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.
* [Not (`!`)](../commands/not-func.md):
  Reads the stdin and exit number from previous process and not's it's condition
* [True (`true`)](../commands/true.md):
  Returns a `true` value

<hr/>

This document was generated from [builtins/core/typemgmt/types_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/typemgmt/types_doc.yaml).