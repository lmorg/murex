# _murex_ Language Guide

## Command reference: and

> Returns `true` or `false` depending on whether multiple conditions are met

### Description

Returns a boolean results (`true` or `false`) depending on whether all of the
code-blocks included as parameters are successful or not.

### Usage

    and: { code-block } { code-block } -> <stdout>

    !and: { code-block } { code-block } -> <stdout>

`and` supports as many or as few code-blocks as you wish.

### Examples

    if { and { = 1+1==2 } { = 2+2==4 } } then {
        out: The laws of mathmatics still exist in this universe.
    }

### Details

`and` does not set the exit number on failure so it is safe to use inside a `try`
or `trypipe` block.

If `and` is prefixed by a bang then it returns `true` only when all code-blocks
are unsuccessful.

#### Code-Block Testing

* `and` tests all code-blocks up until one of the code-blocks is unsuccessful,
  then `and` exits and returns `false`.

* `!and` tests all code-blocks up until one of the code-blocks is successful,
  then `!and` exits and returns `false` (ie `!and` is `not`ing every code-block).

### Synonyms

* !and

### See also

* [`catch`](catch.md): Handles the exception code raised by `try` or `trypipe`
* [`if`](if.md): Conditional statement to execute different blocks of code depending on the
result of the condition
* `not`
* [`or`](or.md): Returns `true` or `false` depending on whether one code-block out of multiple
ones supplied is successful or unsuccessful.
* [`try`](try.md): Handles errors inside a block of code
* [`trypipe`](trypipe.md): Checks state of each function in a pipeline and exits block on error
