# _murex_ reference documents

## Builtin function: try

> Handles errors inside a block of code

### Description

`try` forces a different execution behaviour where a failed process at the end
of a pipeline will cause the block to terminate regardless of any functions that
might follow.

It's usage is similar to try blocks in other languages (eg Java) but a closer
functional example would be `set -e` in Bash.

To maintain concurrency within the pipeline, `try` will only check the last
function in any given pipeline (ie series of functions joined via `|`, `->`, or
similar operators). If you need the entire pipeline checked then use `trypipe`.

### example

    try {
        out: "Hello, World!" -> grep: "non-existent string"
        out: "This command will be ignored"
    }

### Detail

A failure is determined by:

* Any process that returns a non-zero exit number
* Any process that returns more output via STDERR than it does via STDOUT

You can see which run mode your functions are executing under via the `fid-list`
command.

### See also

* [trypipe](trypipe.md): Checks state of each function in a pipeline and exits block on error
* evil
* [catch](catch.md): Handles the exception code raised by `try` or `trypipe`
* fid-list
* [if](if.md): Conditional statement to execute different blocks of code depending on the
result of the condition
