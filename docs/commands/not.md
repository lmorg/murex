# _murex_ Shell Guide

## Command Reference: `!` (not)

> Reads the STDIN and exit number from previous process and not's it's condition

### Description

Reads the STDIN and exit number from previous process and not's it's condition.

### Usage

    <stdin> -> ! -> <stdout>

### Examples

    » echo "Hello, world!" -> !
    false
    
    » false -> !
    true

### Synonyms

* `!`


### See Also

* commands/[`and`](../commands/and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
* commands/[`false`](../commands/false.md):
  Returns a `false` value
* commands/[`if`](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* commands/[`or`](../commands/or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.
* commands/[`true`](../commands/true.md):
  Returns a `true` value