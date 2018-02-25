# _murex_ Language Guide

## Command reference: print

> Write a string to the OS STDOUT (bypassing _murex_ pipelines)

### Description

Write parameters to the OS STDOUT (bypassing STDOUT along the _murex_ pipeline).

### Usage

    print: string to write

### Examples

    » print Hello, World!
    Hello, World!

### Detail

This is a throwaway function for edge cases. Generally all problems should be
solvable using `out` or _murex_ named pipes (`pipe`).

> Please note: because this function writes directly to the OS's STDOUT,
  redirection will not work with `print`.

### See also

* [tout](tout.md): `echo` a string to the STDOUT and set it's data-type
* [err](err.md): `echo` a string to the STDERR
* [print](print.md): Write a string to the OS STDOUT (bypassing _murex_ pipelines)
* sprintf
* ttyfd
