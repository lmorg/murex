# `{ Curly Brace }`

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

Curly braces can be used to terminate the parsing of the command name and/or
parameters too:

```
if{true}{out Yipee}
```

### Nesting

Curly braces can be nested:

```
» out {{foo} bar}
{{foo} bar}
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

* [ANSI Constants](../user-guide/ansi.md):
  Infixed constants that return ANSI escape sequences
* [Code Block Parsing](../user-guide/code-block.md):
  Overview of how code blocks are parsed
* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Error String (`err`)](../commands/err.md):
  Print a line to the stderr
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Output With Type Annotation (`tout`)](../commands/tout.md):
  Print a string to the stdout and set it's data-type
* [`"Double Quote"`](../parser/double-quote.md):
  Initiates or terminates a string (variables expanded)
* [`%(Brace Quote)`](../parser/brace-quote.md):
  Initiates or terminates a string (variables expanded)
* [`'Single Quote'`](../parser/single-quote.md):
  Initiates or terminates a string (variables not expanded)
* [`(brace quote)`](../parser/brace-quote-func.md):
  Write a string to the stdout without new line (deprecated)
* [`@Array` Sigil](../parser/array.md):
  Expand values as an array
* [`string` (stringing)](../types/str.md):
  string (primitive)
* [`~` Home Sigil](../parser/tilde.md):
  Home directory path variable

<hr/>

This document was generated from [gen/parser/codeblock_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/codeblock_doc.yaml).