# Or (`||`) Logical Operator - Parser Reference

> Continues next operation only if previous operation fails

## Description

When in the **normal** run mode (see "schedulers" link below) this will only
run the command on the right hand side if the command on the left hand side
does not error. Neither STDOUT nor STDERR are piped.

This has no effect in `try` nor `trypipe` run modes because they automatically
apply stricter error handling. See detail below.

## Examples

Second command does not run because the first command doesn't error:

    » out: one || out: two
    one
    
Second command does run because the first command produces an error:

    » err: one || out: two
    one
    two

## Detail

This has no effect in `try` nor `trypipe` run modes because they automatically
apply stricter error handling. You can achive a similiar behavior in `try` with
the following code:

    try {
        err: one -> !if { out: two }
    }
    
There is no workaround for `trypipe`.

## See Also

* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [STDERR Pipe (`?`) Token](../parser/pipe-err.md):
  Pipes STDERR from the left hand command to STDIN of the right hand command
* [Schedulers](../user-guide/schedulers.md):
  Overview of the different schedulers (or 'run modes') in Murex
* [`err`](../commands/err.md):
  Print a line to the STDERR
* [`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [`try`](../commands/try.md):
  Handles errors inside a block of code
* [`trypipe`](../commands/trypipe.md):
  Checks state of each function in a pipeline and exits block on error