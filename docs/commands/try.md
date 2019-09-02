# _murex_ Shell Guide

## Command Reference: `try`

> Handles errors inside a block of code

### Description

`try` forces a different execution behavior where a failed process at the end
of a pipeline will cause the block to terminate regardless of any functions that
might follow.

It's usage is similar to try blocks in other languages (eg Java) but a closer
functional example would be `set -e` in Bash.

To maintain concurrency within the pipeline, `try` will only check the last
function in any given pipeline (ie series of functions joined via `|`, `->`, or
similar operators). If you need the entire pipeline checked then use `trypipe`.

### Usage

    try { code-block } -> <stdout>
    
    <stdin> -> try { -> code-block } -> <stdout>

### Examples

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

### See Also

* [commands/`catch`](../commands/catch.md):
  Handles the exception code raised by `try` or `trypipe` 
* [commands/`if`](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [commands/`trypipe`](../commands/trypipe.md):
  Checks state of each function in a pipeline and exits block on error
* [commands/evil](../commands/evil.md):
  
* [commands/fid-list](../commands/fid-list.md):
  