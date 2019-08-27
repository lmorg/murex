# _murex_ Shell Guide

## Command Reference: `rx`

> Regexp pattern matching for file system objects (eg '.*\.txt')

### Description

Returns a list of files and directories that match a regexp pattern.

Output is a JSON list.

### Usage

    rx: pattern -> <stdout>

### Examples

    # inline regex file matching
    cat: @{ rx: '.*\.txt' }
    
    # writing a list of files to disk
    rx: '.*\.go' -> > filelist.txt
    
    # checking if any files exist
    if { rx: somefiles.* } then {
        # files exist
    }
    
    # checking if no files exist
    !if { rx: somefiles.* } then {
        # files do not exist
    }

### Detail

Unlike globbing (`g`) which can traverse directories (eg `g: /path/*`), `rx` is
only designed to match file system objects in the current working directory.

`rx` uses Go (lang)'s standard regexp engine.

### See Also

* [`f`](../commands/f.md):
  Lists objects (eg files) in the current working directory
* [`g`](../commands/g.md):
  Glob pattern matching for file system objects (eg *.txt)
* [`match`](../commands/match.md):
  Match an exact value in an array
* [`regexp`](../commands/regexp.md):
  Regexp tools for arrays / lists of strings