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

### See also

* >
* >>
* cast
* [err](err.md): `echo` a string to the STDERR
* [print](print.md): Write a string to the OS STDOUT (bypassing _murex_ pipelines)
* [pt](pt.md): Pipe telemetry. Writes data-types and bytes written
* sprintf
* [tout](tout.md): `echo` a string to the STDOUT and set it's data-type
* ttyfd
