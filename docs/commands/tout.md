# Output With Type Annotation (`tout`)

> Print a string to the stdout and set it's data-type

## Description

Write parameters to stdout without a trailing new line character. Cast the
output's data-type to the value of the first parameter.

## Usage

```
tout data-type "string to write" -> <stdout>
```

## Examples

```
» tout json { "Code": 404, "Message": "Page not found" } -> pretty
{
    "Code": 404,
    "Message": "Page not found"
}
```

## Detail

`tout` supports ANSI constants.

Unlike `out`, `tout` does not append a carriage return / line feed.

## Synonyms

* `tout`


## See Also

* [ANSI Constants](../user-guide/ansi.md):
  Infixed constants that return ANSI escape sequences
* [Define Type (`cast`)](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [Error String (`err`)](../commands/err.md):
  Print a line to the stderr
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Prettify JSON](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [Reformat Data type (`format`)](../commands/format.md):
  Reformat one data-type into another data-type
* [`(brace quote)`](../parser/brace-quote-func.md):
  Write a string to the stdout without new line (deprecated)

<hr/>

This document was generated from [builtins/core/io/echo_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/io/echo_doc.yaml).