# _murex_ Shell Docs

## Command Reference: `true`

> Returns a `true` value

### Description

Returns a `true` value.

### Usage

    true -> <stdout>

### Examples

By default, `true` also outputs the term "true":

    » true
    true
    
However you can suppress that with the silent flag:

    » true -s

### Flags

* `-s`
    silent - don't output the term "true"

### See Also

* [commands/`!` (not)](../commands/not.md):
  Reads the STDIN and exit number from previous process and not's it's condition
* [commands/`and`](../commands/and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
* [commands/`false`](../commands/false.md):
  Returns a `false` value
* [commands/`if`](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [commands/`or`](../commands/or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.