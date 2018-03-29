# _murex_ Language Guide

## Command reference: rx

> Regexp pattern matching for file system objects (eg '.*\.txt')

### Description

Returns a list of files and directories that match a regexp pattern.

Output is a JSON list.

### Usage

    rx: pattern -> <stdout>

### Examples

    cat: @{ rx: '.*\.txt' }

    g: '.*\.go' -> > filelist.txt

### Details

Unlink globbing (`g`) which can traverse directories, `rx` is only designed to
match file system objects in the current working directory.

`rx` uses Go (lang)'s standard regexp engine.

### See also

* [`f`](f.md): Lists objects (eg files) in the current working directory
* [`g`](g.md): Glob pattern matching for file system objects (eg *.txt)
