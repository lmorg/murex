# `or` - Command Reference

> Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.

## Description

Returns a boolean results (`true` or `false`) depending on whether any of the
code-blocks included as parameters are successful or not.

## Usage

    or: { code-block } { code-block } -> <stdout>
    
    !or: { code-block } { code-block } -> <stdout>
    
`or` supports as many or as few code-blocks as you wish.

## Examples

    if { or { = 1+1==2 } { = 2+2==5 } { = 3+3==6 } } then {
        out: At least one of those equations are correct
    }

## Detail

`or` does not set the exit number on failure so it is safe to use inside a `try`
or `trypipe` block.

If `or` is prefixed by a bang (`!or`) then it returns `true` when one or more
code-blocks are unsuccessful (ie the opposite of `or`).

### Code-Block Testing

* `or` only executes code-blocks up until one of the code-blocks is successful
  then it exits the function and returns `true`.

* `!or` only executes code-blocks while the code-blocks are successful. Once one
  is unsuccessful `!or` exits and returns `true` (ie it `not`s every code-block).

## Synonyms

* `or`
* `!or`


## See Also

* [`!` (not)](../commands/not.md):
  Reads the STDIN and exit number from previous process and not's it's condition
* [`and`](../commands/and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
* [`catch`](../commands/catch.md):
  Handles the exception code raised by `try` or `trypipe` 
* [`false`](../commands/false.md):
  Returns a `false` value
* [`if`](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [`true`](../commands/true.md):
  Returns a `true` value
* [`try`](../commands/try.md):
  Handles errors inside a block of code
* [`trypipe`](../commands/trypipe.md):
  Checks state of each function in a pipeline and exits block on error