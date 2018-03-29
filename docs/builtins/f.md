# _murex_ Language Guide

## Command reference: f

> Lists objects (eg files) in the current working directory

### Description

List objects (eg files) in the current working directory.

Output is a JSON list.

### Usage

    f: options -> <stdout>

### Examples

    # return only directories:
    f: +d

    # return file and directories but exclude symlinks:
    f: +d +f -s

### Details

By default `f` will return no results. It is then your responsibility to select
which types of objects to be included or excluded from the results.

#### Flags

* `f`: files
* `d`: directories
* `s`: symlinks (automatically included with files and directories)

### See also

* [`g`](g.md): Glob pattern matching for file system objects (eg *.txt)
* [`rx`](rx.md): Regexp pattern matching for file system objects (eg '.*\.txt')
