# _murex_ Language Guide

## Command Reference: `f`

> Lists objects (eg files) in the current working directory

### Description

Lists objects (eg files) in the current working directory.

### Usage

    f: options -> <stdout>

### Examples

    # return only directories:
    f: +d
    
    # return file and directories but exclude symlinks:
    f: +d +f -s

### Flags

* `d`
    directories
* `f`
    files
* `s`
    symlinks (automatically included with files and directories)

### Detail

By default `f` will return no results. It is then your responsibility to select
which types of objects to be included or excluded from the results.

### See Also

* [`g`](../docs/commands/g.md):
  Glob pattern matching for file system objects (eg *.txt)
* [`rx`](../docs/commands/rx.md):
  Regexp pattern matching for file system objects (eg '.*\.txt')