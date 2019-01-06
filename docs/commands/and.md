# _murex_ Language Guide

## Command Reference: `and`

> Returns `true` or `false` depending on whether multiple conditions are met

### Description

Returns a boolean results (`true` or `false`) depending on whether all of the
code-blocks included as parameters are successful or not.

### Usage

    and: { code-block } { code-block } -> <stdout>
    
    !and: { code-block } { code-block } -> <stdout>
    
`and` supports as many or as few code-blocks as you wish.

### Examples

    if { and { = 1+1==2 } { = 2+2==4 } { = 3+3==6 } } then {
        out: The laws of mathematics still exist in this universe.
    }

### Detail

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

* `and`
* `!and`


### See Also

* [`catch`](../docs/commands/catch.md):
  Handles the exception code raised by `try` or `trypipe
* [`if`](../docs/commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [`or`](../docs/commands/or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.
* [`try`](../docs/commands/try.md):
  Handles errors inside a block of code
* [`trypipe`](../docs/commands/trypipe.md):
  Checks state of each function in a pipeline and exits block on error
* [false](../docs/commands/commands/false.md):
  
* [not](../docs/commands/commands/not.md):
  
* [true](../docs/commands/commands/true.md):
  