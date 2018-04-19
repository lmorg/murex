# _murex_ Language Guide

## Command reference: tout

> `echo` a string to the STDOUT and set it's data-type

### Description

Write parameters to STDOUT and cast the output's data-type.

### Usage

    out: data-type "string to write" -> <stdout>

### Examples

    » tout json { "Code": 404, "Message": "Page not found" } -> pretty
    {
        "Code": 404,
        "Message": "Page not found"
    }

### Details

`tout` supports ANSI constants.

### See also

* [`brace-quote`](brace-quote.md): Write a string to the STDOUT without new line
* `cast`
* [`err`](err.md): `echo` a string to the STDERR
* `format`
* [`out`](out.md): `echo` a string to the STDOUT with a trailing new line character
* `pretty`
* `sprintf`
* [`ttyfd`](ttyfd.md): Returns the TTY device of the parent.
