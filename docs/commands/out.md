# `out`

> Print a string to the STDOUT with a trailing new line character

## Description

Write parameters to STDOUT with a trailing new line character.

## Usage

```
out string to write -> <stdout>
```

## Examples

```
» out Hello, World!
Hello, World!
```

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

* [`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`err`](../commands/err.md):
  Print a line to the STDERR
* [`pt`](../commands/pt.md):
  Pipe telemetry. Writes data-types and bytes written
* [`read`](../commands/read.md):
  `read` a line of input from the user and store as a variable
* [`tout`](../commands/tout.md):
  Print a string to the STDOUT and set it's data-type
* [`tread`](../commands/tread.md):
  `read` a line of input from the user and store as a user defined *typed* variable (deprecated)
* [ansi](../commands/ansi.md):
  
* [brace-quote](../commands/brace-quote.md):
  
* [greater-than](../commands/greater-than.md):
  
* [greater-than-greater-than](../commands/greater-than-greater-than.md):
  

<hr/>

This document was generated from [builtins/core/io/echo_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/io/echo_doc.yaml).