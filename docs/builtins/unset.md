# _murex_ Language Guide

## Command reference: unset

> Deallocates an environmental variable (aliased to `!export`)

### Description

`unset` internally points to the same function as `!export` and exists purely
for compatability with other shells (eg Bash).

### Usage

    unset var_name

### Details

Please read the docs on `export` (link below).

### See also

* `export`
