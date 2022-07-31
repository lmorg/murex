# _murex_ Shell Docs

## Command Reference: `rx`

> Regexp pattern matching for file system objects (eg '.*\.txt')

## Description

Returns a list of files and directories that match a regexp pattern.

Output is a JSON list.

## Usage

    rx: pattern -> <stdout>
    
    !rx: pattern -> <stdout>

## Examples

Inline regex file matching:

    cat: @{ rx: '.*\.txt' }
    
Writing a list of files to disk:

    rx: '.*\.go' -> > filelist.txt
    
Checking if files exist:

    if { rx: somefiles.* } then {
        # files exist
    }
    
Checking if files do not exist:

    !if { rx: somefiles.* } then {
        # files do not exist
    }
    
Return all files apart from text files:

    !g: '\.txt$'

## Detail

### Traversing Directories

Unlike globbing (`g`) which can traverse directories (eg `g: /path/*`), `rx` is
only designed to match file system objects in the current working directory.

`rx` uses Go (lang)'s standard regexp engine.

### Inverse Matches

If you want to exclude any matches based on wildcards, rather than include
them, then you can use the bang prefix. eg

    » rx: READ*                                                                                                                                                              
    [
        "README.md"
    ]
    
    murex-dev» !rx: .*
    Error in `!rx` (1,1): No data returned.

## Synonyms

* `rx`
* `!rx`


## See Also

* [commands/`f`](../commands/f.md):
  Lists or filters file system objects (eg files)
* [commands/`g`](../commands/g.md):
  Glob pattern matching for file system objects (eg *.txt)
* [commands/`match`](../commands/match.md):
  Match an exact value in an array
* [commands/`regexp`](../commands/regexp.md):
  Regexp tools for arrays / lists of strings