# Is Value Null (`is-null`)

> Checks if a variable is null or undefined

## Description

`is-null` checks if a variable is null or undefined. If multiple variables are
passed in `is-null`'s parameters, the exit number will be a count of the number
of non-null variables checked.

`is-null` is intended to be run non-interactively, where it doesn't write to
stdout but instead communicates its results via exit number. However if stdout
is a TTY `is-null` will additionally write to the terminal.

The following conditions are considered "null" by `is-null`:

* a variable not being defined
* a property not existed, eg an object key or array index
* any other error reading from a variable
* or the value of the variable or property being the data-type `null` or a
  value of null.

Zero length strings, strings containing the word "null" and numeric data types
(eg `num`, `int`, `float` with the value of `0`) are all **not null**.

## Usage

```
is-null variable_name... -> <stdout>
```

## Examples

### Interactive output

```
» $baz = ""
» is-null foo bar baz
foo: undefined or null
bar: undefined or null
baz: defined and not null
```

### None interactive output

```
if { is-null foobar } then {
    out "baz is undefined"
}
```

## See Also

* [Define Environmental Variable (`export`)](../commands/export.md):
  Define an environmental variable and set it's value
* [Define Global (`global`)](../commands/global.md):
  Define a global variable and set it's value
* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [Variable and Config Scoping](../user-guide/scoping.md):
  How scoping works within Murex
* [`??` Null Coalescing Operator](../parser/null-coalescing.md):
  Returns the right operand if the left operand is empty / undefined (expression)

<hr/>

This document was generated from [builtins/core/typemgmt/isnull_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/typemgmt/isnull_doc.yaml).