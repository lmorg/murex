# `round`  - Command Reference

> Round a number by a user defined precision

## Description

`round` supports a few different levels of precision

## Usage

    round input precision -> <stdout>

## Flags

* `--down`
    Rounds down to the nearest multiple (not supported when precision is to decimal places)
* `--up`
    Rounds up to the nearest multiple (not supported when precision is to decimal places)
* `-d`
    alias of `--down
* `-u`
    alias of `--up

## See Also

* [`expr`](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators