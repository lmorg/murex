# _murex_ Shell Guide

## Command Reference: `g`

> Glob pattern matching for file system objects (eg *.txt)

### Description

Returns a list of files and directories that match a glob pattern.

Output is a JSON list.

### Usage

    g: pattern -> <stdout>

### Examples

    # inline globbing
    cat: @{ g: *.txt }
    
    # writing a JSON array of files to disk
    g: *.txt -> > filelist.json
    
    # writing a list of files to disk
    g: *.txt -> format str -> > filelist.txt
    
    # checking if a file exists
    if { g: somefile.txt } then {
        # file exists
    }
    
    # checking if a file does not exist
    !if { g: somefile.txt } then {
        # file does not exist
    }

### Detail

#### Pattern reference

* `*` matches any number of (including zero) characters
* `?` matches any single character

#### Auto-globbing

Any command prefixed with `@g` will be auto-globbed. For example, the two
following commands will produce the same output:

    » ls @{g: *.go}
    benchmarks_test.go  defaults_test.go  flags.go  godoc.go  main.go  murex_test.go
    
    » @g ls *.go
    benchmarks_test.go  defaults_test.go  flags.go  godoc.go  main.go  murex_test.go
    
The rational behind the ugly `@g` syntax is simply to make one-liners a bit
less painful when coming from more traditional POSIX-like shells (eg Bash)
where wildcards are automatically expanded. So if you type `ls *` (for example)
then realise you've forgotten to subshell, you can just recall the last command
with auto-globbing enabled:

    @g ^!!

### Synonyms

* `g`
* `@g`


### See Also

* commands/[`f`](../commands/f.md):
  Lists objects (eg files) in the current working directory
* commands/[`rx`](../commands/rx.md):
  Regexp pattern matching for file system objects (eg '.*\.txt')