# `%[]` Array Builder

> Quickly generate arrays

## Description

`%[]` is a way of defining arrays in expressions and statements. Whenever a
`%[]` array is outputted as a string, it will be converted to minified JSON.

Array elements inside `%[]` can be whitespace and/or comma delimited. This
allows for compatibility with both Bash muscle memory, and people more
familiar with JSON.

Additionally you can also embed `a` style parameters inside `%[]` arrays too.

Like with YAML, strings in `%[]` do not need to be quoted unless you need to
force numeric or boolean looking values to be stored as strings.



## Examples

### Arrays passed as a JSON string:

```
» echo %[1..3]
[1,2,3]

» %[1..3] -> cat
[1,2,3]
```

### Different supported syntax for creating a numeric array:

#### As a range

```
» %[1..3]
[
    1,
    2,
    3
]
```

#### JSON formatted

```
» %[1,2,3]
[
    1,
    2,
    3
]
```

#### Whitespace separated

```
» %[1 2 3]
[
    1,
    2,
    3
]
```

#### Values and ranges

```
» %[1,2..3]
[
    1,
    2,
    3
]
```

### Strings

#### barewords and whitespace separated

This will allow you to copy/paste lists from traditional shells like Bash

```
» %[foo bar]
[
    "foo",
    "bar"
]
```

#### JSON formatted

```
» %["foo", "bar"]
[
    "foo",
    "bar"
]
```

### Special ranges

```
» %[June..August]
[
    "June",
    "July",
    "August"
]
```

A full list of special ranges are available at [docs/mkarray/special](../mkarray/special.md)

### Multiple expansion blocks:

```
» %[[A,B]:[1..4]]
[
    "A:1",
    "A:2",
    "A:3",
    "A:4",
    "B:1",
    "B:2",
    "B:3",
    "B:4"
]
```

### Nested arrays:

```
» %[foo [bar]]
[
    "foo",
    [
        "bar"
    ]
]
```

The `%` prefix for the nested array is optional.

### JSON objects within arrays

```
» %[foo {bar: baz}]
[
    "foo",
    {
        "bar": "baz"
    }
]
```

The `%` prefix for the nested object is optional.

## Detail

Murex supports a number of different formats that can be used to generate
arrays. For more details on these please refer to the documents for each format

* [Calendar Date Ranges](../mkarray/date.md):
  Create arrays of dates
* [Character arrays](../mkarray/character.md):
  Making character arrays (a to z)
* [Decimal Ranges](../mkarray/decimal.md):
  Create arrays of decimal integers
* [Non-Decimal Ranges](../mkarray/non-decimal.md):
  Create arrays of integers from non-decimal number bases
* [Special Ranges](../mkarray/special.md):
  Create arrays from ranges of dictionary terms (eg weekdays, months, seasons, etc)

## See Also

* [Create JSON Array (`ja`)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [Create New Array (`ta`)](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type
* [Expressions (`expr`)](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [Special Ranges](../mkarray/special.md):
  Create arrays from ranges of dictionary terms (eg weekdays, months, seasons, etc)
* [Stream New List (`a`)](../commands/a.md):
  A sophisticated yet simple way to stream an array or list (mkarray)
* [`"Double Quote"`](../parser/double-quote.md):
  Initiates or terminates a string (variables expanded)
* [`%(Brace Quote)`](../parser/brace-quote.md):
  Initiates or terminates a string (variables expanded)
* [`%{}` Object Builder](../parser/create-object.md):
  Quickly generate objects (dictionaries / maps)
* [`'Single Quote'`](../parser/single-quote.md):
  Initiates or terminates a string (variables not expanded)

<hr/>

This document was generated from [gen/parser/create_array_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/create_array_doc.yaml).