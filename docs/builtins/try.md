# _murex_ reference documents

## builtin function: try

> Exit block when error found. Similar in function to Bash's `set -e`

### Description

`try` forces a different execution behaviour where a failed process at the end
of a pipeline will cause the block to terminate regardless of any functions that
might follow.

### example

    try {
        out: "Hello, World!" -> grep: "non-existent string"
        out: "This process will be ignored"
    }

### Detail

A failure is determined by:

* Any process that returns a non-zero exit number
* Any process that returns more output via STDERR than it does via STDOUT

You can see which run mode your functions are executing under via the `fid-list`
command:

### See also

* [trypipe](trypipe.md): Checks state of each function in a pipeline and exits block on error
* evil
* catch
* fid-list
