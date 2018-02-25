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

### See also

* [tout](tout.md): `echo` a string to the STDOUT and set it's data-type
* [err](err.md): `echo` a string to the STDERR
* [print](print.md): Write a string to the OS STDOUT (bypassing _murex_ pipelines)
* sprintf
* cast
* pretty
* format
* ttyfd
