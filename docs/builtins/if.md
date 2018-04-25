# _murex_ Language Guide

## Command reference: if

> Conditional statement to execute different blocks of code depending on the
result of the condition

### Description

Conditional control flow

`if` can be utilised both as a method as well as a standalone function. As a
method, the conditional state is derived from the calling function (eg if the
previous function succeeds then the condition is `true`).

### Usage

    # function if
    if { code-block } then {
        # true
    } else {
        # false
    }

    # method if
    command -> if {
        # true
    } else {
        # false
    }

    # negative function if
    !if { code-block } then {
        # false
    }

    # negative method if
    command -> !if {
        # false
    }

Note: the `then` and `else` statements are optional. So the first usage could
also be written as:

    if { code-block } {
        # true
    } {
        # false
    }

However the practice of omitting those statements isn't recommended beyond
writing short one liners in the command line.

### example

    # Check if a file exists
    if { g somefile.txt } then {
        out "File exists"
    }

    # ...or does not exist
    !if { g somefile.txt } then {
        out "File does not exist"
    }


### Detail

The conditional block can contain entire pipelines - even multiple lines of code
let alone a single pipeline - as well as solitary commands as demonstrated in
the examples above. However the conditional block does not output STDOUT nor
STDERR to the rest of the pipeline so you don't have to worry about redirecting
the output streams to `null`.

### Synonyms

* !if

### See also

* `and`
* [`catch`](catch.md): Handles the exception code raised by `try` or `trypipe`
* `false`
* `not`
* `or`
* `true`
* [`try`](try.md): Handles errors inside a block of code
* [`trypipe`](trypipe.md): Checks state of each function in a pipeline and exits block on error
