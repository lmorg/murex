# `out`

> Print a string to the stdout with a trailing new line character

## Description

Write parameters to stdout with a trailing new line character.

## Usage

```
out string to write -> <stdout>
```

## Examples

### out

```
» out Hello, World!
Hello, World!
```

### echo

For compatibility with other shells, `echo` is also supported:

```
» echo Hello, World!
Hello, World!
```

## Detail

`out` / `echo` output as `string` data-type. This can be changed by casting
(`cast`) or using the `tout` function.

### ANSI Constants

`out` supports ANSI constants.

## Synonyms

* `out`
* `echo`


## See Also

* [ANSI Constants](../user-guide/ansi.md):
  Infixed constants that return ANSI escape sequences
* [`(brace quote)`](../parser/brace-quote-func.md):
  Write a string to the stdout without new line (deprecated)
* [`>>` Append File](../parser/file-append.md):
  Writes stdin to disk - appending contents if file already exists
* [`cast`](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [`err`](../commands/err.md):
  Print a line to the stderr
* [`pt`](../commands/pt.md):
  Pipe telemetry. Writes data-types and bytes written
* [`read`](../commands/read.md):
  `read` a line of input from the user and store as a variable
* [`tout`](../commands/tout.md):
  Print a string to the stdout and set it's data-type
* [`tread`](../commands/tread.md):
  `read` a line of input from the user and store as a user defined *typed* variable (deprecated)
* [`|>` Truncate File](../parser/file-truncate.md):
  Writes stdin to disk - overwriting contents if file already exists

<hr/>

This document was generated from [builtins/core/io/echo_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/io/echo_doc.yaml).