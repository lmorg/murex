# _murex_ Shell Docs

## Command Reference: `g`

> Glob pattern matching for file system objects (eg *.txt)

## Description

Returns a list of files and directories that match a glob pattern.

Output is a JSON list.

## Usage

    g: pattern -> <stdout>
    
    [ <stdin> -> ] @g command pattern [ -> <stdout> ]
    
    !g: pattern -> <stdout>

## Examples

Inline globbing:

    cat: @{ g: *.txt }
    
Writing a JSON array of files to disk:

    g: *.txt -> > filelist.json
    
Writing a list of files to disk:

    g: *.txt -> format str -> > filelist.txt
    
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
    
Auto-globbing (eg for Bash compatibility):

    @g ls *.txt

## Detail

### Pattern Reference

* `*` matches any number of (including zero) characters
* `?` matches any single character

### Auto-Globbing

Any command prefixed with `@g` will be auto-globbed. For example, the two
following commands will produce the same output:

    » ls @{g: *.go}
    benchmarks_test.go  defaults_test.go  flags.go  godoc.go  main.go  murex_test.go
    
    » @g ls: *.go
    benchmarks_test.go  defaults_test.go  flags.go  godoc.go  main.go  murex_test.go
    
The rational behind the ugly `@g` syntax is simply to make one-liners a bit
less painful when coming from more traditional POSIX-like shells (eg Bash)
where wildcards are automatically expanded. So if you type `ls *` (for example)
then realise you've forgotten to subshell, you can just recall the last command
with auto-globbing enabled:

    @g ^!!
    
### Inverse Matches

If you want to exclude any matches based on wildcards, rather than include
them, then you can use the bang prefix. eg

    » g: READ*
    [
        "README.md"
    ]
    
    » !g: *
    Error in `!g` (1,1): No data returned.

## Synonyms

* `g`
* `@g`
* `!g`


## See Also

* [commands/`f`](../commands/f.md):
  Lists objects (eg files) in the current working directory
* [commands/`rx`](../commands/rx.md):
  Regexp pattern matching for file system objects (eg '.*\.txt')