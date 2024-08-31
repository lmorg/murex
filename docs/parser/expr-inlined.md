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

* [Expressions (`expr`)](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [Operators And Tokens](../user-guide/operators-and-tokens.md):
  A table of all supported operators and tokens
* [Output String (`echo`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character

<hr/>

This document was generated from [gen/parser/expr-inlined_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/expr-inlined_doc.yaml).