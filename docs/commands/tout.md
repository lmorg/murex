# _murex_ Shell Guide

## Command Reference: `tout`

> Print a string to the STDOUT and set it's data-type

### Description

Write parameters to STDOUT without a trailing new line character. Cast the
output's data-type to the value of the first parameter.

### Usage

    tout: data-type "string to write" -> <stdout>

### Examples

    » tout: json { "Code": 404, "Message": "Page not found" } -> pretty
    {
        "Code": 404,
        "Message": "Page not found"
    }

### Detail

`tout` supports ANSI constants.

Unlike `out`, `tout` does not append a carriage return / line feed.

### See Also

* [`(` (brace quote)](../commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`err`](../commands/err.md):
  Print a line to the STDERR
* [`out`](../commands/out.md):
  `echo` a string to the STDOUT with a trailing new line character
* [`pretty`](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [format](../commands/format.md):
  
* [sprintf](../commands/sprintf.md):
  