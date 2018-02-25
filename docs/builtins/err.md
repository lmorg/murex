# _murex_ Language Guide

## Command reference: err

> `echo` a string to the STDERR

### Description

Write parameters to STDERR.

### Usage

    err: string to write -> <stderr>

### Examples

    » err Hello, World!
    Hello, World!

### Detail

`err` outputs as `string` data-type. This can be changed by casting

    err { "Code": 404, "Message": "Page not found" } ? cast json

However passing structured data-types along the STDERR stream is not recommended
as any other function within your code might also pass error messages along the
same stream and thus taint your structured data. This is why _murex_ does not
supply a `tout` function for STDERR. The recommended solution for passing
messages like these which you want separate from your STDOUT stream is to create
a new _murex_ named pipe.

    » pipe: --create messages
    » fork { <messages> -> pretty }
    » tout: <messages> json { "Code": 404, "Message": "Page not found" }
    » pipe: --close messages
    {
        "Code": 404,
        "Message": "Page not found"
    }

### See also

* \>
* \>\>
* cast
* [err](err.md): `echo` a string to the STDERR
* fork
* pipe
* pretty
* [print](print.md): Write a string to the OS STDOUT (bypassing _murex_ pipelines)
* sprintf
* [tout](tout.md): `echo` a string to the STDOUT and set it's data-type
* ttyfd
