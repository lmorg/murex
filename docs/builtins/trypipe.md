# _murex_ Language Guide

## Command reference: trypipe

> Checks state of each function in a pipeline and exits block on error

### Description

`trypipe` checks the state of each function and exits the block if any of them
fail. Where `trypipe` differs from regular `try` blocks is `trypipe` will check
every process along the pipeline as well as the terminating function (which
`try` only validates against). The downside to this is that piped functions can
no longer run in parallel.

### Usage

    trypipe { code-block } -> <stdout>

### Example

    trypipe {
        out: "Hello, World!" -> grep: "non-existent string" -> cat
        out: "This command will be ignored"
    }

### Detail

A failure is determined by:

* Any process that returns a non-zero exit number
* Any process that returns more output via STDERR than it does via STDOUT

You can see which run mode your functions are executing under via the `fid-list`
command.

### See also

* [`catch`](catch.md): Handles the exception code raised by `try` or `trypipe`
* `evil``
* `fid-list``
* [`try`](try.md): Handles errors inside a block of code
