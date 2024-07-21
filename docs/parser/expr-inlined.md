# ( expression )

> Inline expressions

## Description

Expressions can be inlined via parenthesis. For example:

```
» echo (1 + 2)
3
```

The opening brace has to be the first character of that parameter, otherwise
that parameter will be treated like a string.

```
» echo 1+(1 + 2)
1+(1 + 2)
```



## See Also

* [`echo`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [`expr`](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [operations-and-tokens](../parser/operations-and-tokens.md):
  

<hr/>

This document was generated from [gen/parser/expr-inlined_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/expr-inlined_doc.yaml).