# _murex_ Language Guide

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

* [`(` (brace quote)](../docs/commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [`err`](../docs/commands/err.md):
  Print a line to the STDERR
* [`out`](../docs/commands/out.md):
  `echo` a string to the STDOUT with a trailing new line character
* [cast](../docs/commands/commands/cast.md):
  
* [format](../docs/commands/commands/format.md):
  
* [pretty](../docs/commands/commands/pretty.md):
  
* [sprintf](../docs/commands/commands/sprintf.md):
  