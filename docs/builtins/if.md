# _murex_ Language Guide

## Command reference: if

> Conditional statement to execute different blocks of code depending on the
result of the condition

### Description

`if` can be called in two different ways:

1. Method If: `conditional -> if: { true } { false }`
2. Function If: `if: { conditional } { true } { false }`

The conditional is evaluated based on the output produced by the
function and the exit number. Any non-zero exit numbers are an automatic
"false". Any functions returning no data are also classed as a "false".
For a full list of conditions that are evaluated to determine a true or
false state of a function, please read the documentation on the `boolean`
data type in [GUIDE.syntax.md](../GUIDE.syntax.md#boolean).

Please also note that while the last parameter is optional, if it is
left off and `if` or `!if` would have otherwise called it, then `if` /
`!if` will return a non-zero exit number. The significance of this is
important when using `if` or `!if` inside a `try` block.

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

* [`catch`](catch.md): Handles the exception code raised by `try` or `trypipe`
* [`try`](try.md): Handles errors inside a block of code
* [`trypipe`](trypipe.md): Checks state of each function in a pipeline and exits block on error
