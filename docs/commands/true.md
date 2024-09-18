# True (`true`)

> Returns a `true` value

## Description

Returns a `true` value.

## Usage

```
true -> <stdout>
```

## Examples

### No flags

By default, `true` also outputs the term "true":

```
» true
true
```

### Silent

You can suppress that with the silent flag:

```
» true -s
```

## Flags

* `-s`
    silent - don't output the term "true"

## See Also

* [False (`false`)](../commands/false.md):
  Returns a `false` value
* [If Conditional (`if`)](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [Logic And Statements (`and`)](../commands/and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
* [Logic Or Statements (`or`)](../commands/or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.
* [Not (`!`)](../commands/not-func.md):
  Reads the stdin and exit number from previous process and not's it's condition

<hr/>

This document was generated from [builtins/core/typemgmt/types_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/typemgmt/types_doc.yaml).