# `expr`

> Expressions: mathematical, string comparisons, logical operators

## Description



## Usage

```
expr expression -> <stdout>
```

## Examples

Order of operations:

```
» 3 * (3 + 1)                                                                                                                                                                                                                         
12
```

JSON array:

```
» %[apples oranges grapes]
[
    "apples",
    "oranges",
    "grapes"
]
```

## See Also

* [`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable (deprecated)
* [`set`](../commands/set.md):
  Define a local variable and set it's value
* [equ](../commands/equ.md):
  

<hr/>

This document was generated from [builtins/core/expressions/expressions_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/expressions/expressions_doc.yaml).