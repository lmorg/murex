# _murex_ Shell Docs

## Command Reference: `!` (not)

> Reads the STDIN and exit number from previous process and not's it's condition

## Description

Reads the STDIN and exit number from previous process and not's it's condition.

## Usage

    <stdin> -> ! -> <stdout>

## Examples

    » echo "Hello, world!" -> !
    false
    
    » false -> !
    true

## Synonyms

* `!`


## See Also

* [`and`](../commands/and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
* [`false`](../commands/false.md):
  Returns a `false` value
* [`if`](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [`or`](../commands/or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.
* [`true`](../commands/true.md):
  Returns a `true` value