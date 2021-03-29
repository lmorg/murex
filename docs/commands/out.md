# _murex_ Shell Docs

## Command Reference: `out`

> Print a string to the STDOUT with a trailing new line character

## Description

Write parameters to STDOUT with a trailing new line character.

## Usage

    out: string to write -> <stdout>

## Examples

    » out Hello, World!
    Hello, World!
    
For compatibility with other shells, `echo` is also supported:

    » echo Hello, World!
    Hello, World!

## Detail

`out` / `echo` output as `string` data-type. This can be changed by casting
(`cast`) or using the `tout` function.

### ANSI Constants

`out` supports ANSI constants.

## Synonyms

* `out`
* `echo`


## See Also

* [user-guide/ANSI Constants](../user-guide/ansi.md):
  Infixed constants that return ANSI escape sequences
* [commands/`(` (brace quote)](../commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [commands/`>>` (append file)](../commands/greater-than-greater-than.md):
  Writes STDIN to disk - appending contents if file already exists
* [commands/`>` (truncate file)](../commands/greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists
* [commands/`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [commands/`err`](../commands/err.md):
  Print a line to the STDERR
* [commands/`pt`](../commands/pt.md):
  Pipe telemetry. Writes data-types and bytes written
* [commands/`read`](../commands/read.md):
  `read` a line of input from the user and store as a variable
* [commands/`tout`](../commands/tout.md):
  Print a string to the STDOUT and set it's data-type
* [commands/`tread`](../commands/tread.md):
  `read` a line of input from the user and store as a user defined *typed* variable
* [commands/sprintf](../commands/sprintf.md):
  