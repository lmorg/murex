# `is-null`

> Checks is a variable is null or undefined

## Description

`is-null` returns 

## Usage

```
is-null variable -> <stdout>
```

## Examples

**Interactive output:**

```
» $baz = ""
» is-null foo bar baz
foo: undefined or null
bar: undefined or null
baz: defined and not null
```

**None interactive output:**

```
if { is-null foobar } then {
    out "baz is undefined"
}
```

## See Also

* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [Variable and Config Scoping](../user-guide/scoping.md):
  How scoping works within Murex
* [`??` Null Coalescing Operator](../parser/null-coalescing.md):
  Returns the right operand if the left operand is empty / undefined
* [`export`](../commands/export.md):
  Define an environmental variable and set it's value
* [`global`](../commands/global.md):
  Define a global variable and set it's value
* [`set`](../commands/set.md):
  Define a local variable and set it's value

<hr/>

This document was generated from [builtins/core/typemgmt/isnull_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/typemgmt/isnull_doc.yaml).