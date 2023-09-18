# Create object (`%{}`) constructor

> Quickly generate objects and maps

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

**Object passed as a JSON string:**

```
» echo %{foo: bar}
{"foo":"bar"}
```

**The `%` prefix for the nested object is optional:**

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

## See Also

* [Create array (`%[]`) constructor](../parser/create-array.md):
  Quickly generate arrays
* [Double Quote (`"`) Token](../parser/double-quote.md):
  Initiates or terminates a string (variables expanded)
* [Single Quote (`'`) Token](../parser/single-quote.md):
  Initiates or terminates a string (variables not expanded)
* [Special Ranges](../mkarray/special.md):
  Create arrays from ranges of dictionary terms (eg weekdays, months, seasons, etc)
* [`(` (brace quote)](../commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [`expr`](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators

<hr/>

This document was generated from [gen/parser/create_object_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/create_object_doc.yaml).