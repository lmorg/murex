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
    !if { code-block } else {
        # false
    }

    # negative method if
    command -> if {
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

    if { }

#### Method If

This is where the conditional is evaluated from the result of the piped
function. The last parameter is optional.

    # if / then
    out: hello world | grep: world -> if: { out: world found }

    # if / then / else
    out: hello world | grep: world -> if: { out: world found } { out: world missing }

    if / else
    out: hello world | grep: world -> !if: { out: world missing }

#### Function If

This is where the conditional is evaluated from the first parameter. The
last parameter is optional.

    # if / then / else
    if: { out: hello world | grep: world } { out: world found }

    # if / then / else
    if: { out: hello world | grep: world } { out: world found } { out: world missing }

    if / else
    !if: { out: hello world | grep: world } { out: world missing }

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
