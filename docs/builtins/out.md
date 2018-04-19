# _murex_ Language Guide

## Command reference: out

> `echo` a string to the STDOUT with a trailing new line character

### Description

Write parameters to STDOUT with a trailing new line character.

### Usage

    out: string to write -> <stdout>

### Examples

    » out Hello, World!
    Hello, World!

    » echo Hello, World!
    Hello, World!

(for compatibility with other shells, `echo` is also supported)

### Detail

`out` / `echo` output as `string` data-type. This can be changed by casting or
using the `tout` function.

#### ANSI Constants

`out` supports ANSI constants.

### Synonyms

* echo

### See also

* [`>`](>.md): Writes STDIN to disk - overwriting contents if file already exists
* [`>>`](>>.md): Writes STDIN to disk - appending contents if file already exists
* [`brace-quote`](brace-quote.md): Write a string to the STDOUT without new line
* `cast`
* [`err`](err.md): Print a line to the STDERR
* [`pt`](pt.md): Pipe telemetry. Writes data-types and bytes written
* [`read`](read.md): `read` a line of input from the user and store as a variable
* `sprintf`
* [`tout`](tout.md): Print a string to the STDOUT and set it's data-type
* [`tread`](tread.md): `read` a line of input from the user and store as a user defined typed variable
* [`ttyfd`](ttyfd.md): Returns the TTY device of the parent.
