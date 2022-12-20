# _murex_ Shell Docs

## Command Reference: `g`

> Glob pattern matching for file system objects (eg `*.txt`)

## Description

Returns a list of files and directories that match a glob pattern.

Output is a JSON list.

## Usage

    g: pattern -> <stdout>
    
    [ <stdin> -> ] @g command pattern [ -> <stdout> ]
    
    !g: pattern -> <stdout>
    
    <stdin> -> g: pattern -> <stdout>
    
    <stdin> -> !g: pattern -> <stdout>

## Examples

Inline globbing:

    cat: @{ g: *.txt }
    
Writing a JSON array of files to disk:

    g: *.txt |> filelist.json
    
Writing a list of files to disk:

    g: *.txt -> format str |> filelist.txt
    
Checking if a file exists:

    if { g: somefile.txt } then {
        # file exists
    }
    
Checking if a file does not exist:

    !if { g: somefile.txt } then {
        # file does not exist
    }
    
Return all files apart from text files:

    !g: *.txt
    
Filtering a file list based on glob matches:

    f: +f -> g: *.md
    
Remove any glob matches from a file list:

    f: +f -> !g: *.md

## Detail

### Pattern Reference

* `*` matches any number of (including zero) characters
* `?` matches any single character

### Inverse Matches

If you want to exclude any matches based on wildcards, rather than include
them, then you can use the bang prefix. eg

    » g: READ*
    [
        "README.md"
    ]
    
    » !g: *
    Error in `!g` (1,1): No data returned.
    
### When Used As A Method

`!g` first looks for files that match its pattern, then it reads the file list
from STDIN. If STDIN contains contents that are not files then `!g` might not
handle those list items correctly. This shouldn't be an issue with `frx` in its
normal mode because it is only looking for matches however when used as `!g`
any items that are not files will leak through.

This is its designed feature and not a bug. If you wish to remove anything that
also isn't a file then you should first pipe into either `g: *`, `rx: .*`, or
`f +f` and then pipe that into `!g`.

The reason for this behavior is to separate this from `!regexp` and `!match`.

## Synonyms

* `g`
* `!g`


## See Also

* [commands/`f`](../commands/f.md):
  Lists or filters file system objects (eg files)
* [commands/`match`](../commands/match.md):
  Match an exact value in an array
* [commands/`regexp`](../commands/regexp.md):
  Regexp tools for arrays / lists of strings
* [commands/`rx`](../commands/rx.md):
  Regexp pattern matching for file system objects (eg `.*\\.txt`)