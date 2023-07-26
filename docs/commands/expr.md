# `expr`

> Expressions: mathematical, string comparisons, logical operators

## Description

## Usage

    expr: expression -> `<stdout>`

## Examples

Order of operations:

    » 3 * (3 + 1)
    12

JSON array:

    » %[apples oranges grapes]
    [
        "apples",
        "oranges",
        "grapes"
    ]

## See Also

- [`=` (arithmetic evaluation)](./equ.md):
  Evaluate a mathematical function (deprecated)
- [`let`](./let.md):
  Evaluate a mathematical function and assign to variable (deprecated)
- [`set`](./set.md):
  Define a local variable and set it's value
