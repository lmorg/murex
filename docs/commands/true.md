# `true`

> Returns a `true` value

## Description

Returns a `true` value.

## Usage

```
true -> <stdout>
```

## Examples

By default, `true` also outputs the term "true":

```
» true
true
```

However you can suppress that with the silent flag:

```
» true -s
```

## Flags

- `-s`
  silent - don't output the term "true"

## See Also

- [`!` (not)](./not.md):
  Reads the STDIN and exit number from previous process and not's it's condition
- [`and`](./and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
- [`false`](./false.md):
  Returns a `false` value
- [`if`](./if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
- [`or`](./or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.
