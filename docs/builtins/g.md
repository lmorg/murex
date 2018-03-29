# _murex_ Language Guide

## Command reference: g

> Glob pattern matching for file system objects (eg *.txt)

### Description

Returns a list of files and directories that match a glob pattern.

Output is a JSON list.

### Usage

    g: pattern -> <stdout>

### Examples

    cat @{ g: *.txt }

    g: *.go -> > filelist.txt

### Details

* `*` matches any number of (including zero) characters
* `?` matches any single character

### See also

* [`f`](f.md): Lists objects (eg files) in the current working directory
* [`rx`](rx.md): Regexp pattern matching for file system objects (eg '.*\.txt')
