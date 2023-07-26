# `false`

> Returns a `false` value

## Description

Returns a `false` value.

## Usage

    false -> `<stdout>`

## Examples

By default, `false` also outputs the term "false":

    » false
    false

However you can suppress that with the silent flag:

    » false -s

## Flags

- `-s`
  silent - don't output the term "false"

## See Also

- [`!` (not)](./not.md):
  Reads the STDIN and exit number from previous process and not's it's condition
- [`and`](./and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
- [`if`](./if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
- [`or`](./or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.
- [`true`](./true.md):
  Returns a `true` value
