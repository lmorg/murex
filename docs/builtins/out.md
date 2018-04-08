# _murex_ Language Guide

## Command reference: out

> `echo` a string to the STDOUT

### Description

Write parameters to STDOUT.

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
* `cast`
* [`err`](err.md): `echo` a string to the STDERR
* [`print`](print.md): Write a string to the OS STDOUT (bypassing _murex_ pipelines)
* [`pt`](pt.md): Pipe telemetry. Writes data-types and bytes written
* `sprintf`
* [`tout`](tout.md): `echo` a string to the STDOUT and set it's data-type
* [`ttyfd`](ttyfd.md): Returns the TTY device of the parent.
