# `{Curly Brace}`

> Initiates or terminates a code block

## Description

Curly braces are used to denote the start and end of a code block. Like with
the single quotation marks (`'`), any code inside a curly brace is not parsed.
Also unlike any other quotation tokens, the curly brace is included as part
of the parsed string.

```
» out {example}
{example}
```

Also like the brace quote (`(`, `)`), the curly brace character is only
recognized as a curly brace token if it is the start of a parameter.

Curly braces are also used for other fields besides code blocks. For example
inlining JSON.



## Detail

### Multiline Blocks

Curly braces can work over multiple lines

```
» out {foo
» bar}
{foo
bar}
```

### Code Golfing

Curly braces can be used to terminate the parsing of the command name / start
the parsing of the first parameter however each new parameter would still need
to be separated by whitespace:

```
# Valid
if{true} {out "Yipee"}

# Invalid
if{true}{out "Yipee"}
```

### ANSI Constants

Some builtins (like `out`) also support infixing using the curly brace. eg

```
out "{GREEN}PASSED{RESET}"
```

This is a separate layer of parsing and happens at the parameter level for
specific builtins which opt to support ANSI constants. See the ANSI Constant
user guide (link below) for more information on supporting builtins and which
constants are available.

## See Also

* [Array (`@`) Token](../parser/array.md):
  Expand values as an array
* [Tilde (`~`) Token](../parser/tilde.md):
  Home directory path variable
* [`"Double Quote"`](../parser/double-quote.md):
  Initiates or terminates a string (variables expanded)
* [`$variable`](../parser/string.md):
  Expand values as a string
* [`'Single Quote'`](../parser/single-quote.md):
  Initiates or terminates a string (variables not expanded)
* [`(brace quote)`](../parser/brace-quote.md):
  Write a string to the STDOUT without new line
* [`(brace quote)`](../parser/brace-quote.md):
  Write a string to the STDOUT without new line
* [ansi](../parser/ansi.md):
  
* [code-block](../parser/code-block.md):
  
* [err](../parser/err.md):
  
* [out](../parser/out.md):
  
* [set](../parser/set.md):
  
* [tout](../parser/tout.md):
  

<hr/>

This document was generated from [gen/parser/codeblock_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/codeblock_doc.yaml).