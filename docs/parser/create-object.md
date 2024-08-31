# `%{}` Object Builder

> Quickly generate objects (dictionaries / maps)

## Description

`%{}` is a way of defining objects in expressions and statements. Whenever an
`%{}` object is outputted as a string, it will be converted to minified JSON.

Object elements inside `%{}` can be new line and/or comma delimited. This
allows for compatibility with JSON but without the pain of accidentally invalid
comma management breaking JSON parsers. However a colon is still required to
separate keys from values.

Like with YAML, strings in `%[]` do not need to be quoted unless you need to
force numeric or boolean looking values to be stored as strings.



## Examples

### Object passed as a JSON string

```
» echo %{foo: bar}
{"foo":"bar"}
```

### Nested objects

The `%` prefix for the nested object is optional:

```
» %{foo: bar, baz: [1 2 3]}
{
    "baz": [
        1,
        2,
        3
    ],
    "foo": "bar"
}
```

## Detail

### Syntax

The syntax is pretty flexible, albeit all Murex objects are displayed as JSON
objects when printed to screen or otherwise handled like a string.

#### The creation syntax follows these rules:

Each **key** needs to be followed by a colon, `:`.

Each **value** needs to be followed by either a comma, `,`, new line or closing
curly brace, `}`.

Strings can be quoted or unquoted (like with statement parameters). However any
unquoted values will first be tested to see if it is a number, boolean (`true`
of `false`) or null (`null`). Keys are always strings (`str`), even if they
look like a number.

## See Also

* [Expressions (`expr`)](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [Special Ranges](../mkarray/special.md):
  Create arrays from ranges of dictionary terms (eg weekdays, months, seasons, etc)
* [`"Double Quote"`](../parser/double-quote.md):
  Initiates or terminates a string (variables expanded)
* [`%(Brace Quote)`](../parser/brace-quote.md):
  Initiates or terminates a string (variables expanded)
* [`%[]` Array Builder](../parser/create-array.md):
  Quickly generate arrays
* [`'Single Quote'`](../parser/single-quote.md):
  Initiates or terminates a string (variables not expanded)

<hr/>

This document was generated from [gen/parser/create_object_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/create_object_doc.yaml).