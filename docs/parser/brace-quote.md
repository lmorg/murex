# `%(Brace Quote)`

> Initiates or terminates a string (variables expanded)

## Description

Brace quote is used to initiate and terminate strict strings where variables
can be expanded.

While brace quotes are untraditional compared to your typical string quotations
in POSIX shells, brace quotes have one advantage in that the open and close
grapheme differ (ie `(` is a different character to `)`). This brings benefits
when nesting quotes as it saves the developer from having to carefully escape
the nested quotation marks just the right number of times.

Commands cannot be quoted using brace quotes because `%(` is recognized as its
own function.



## Examples

### As a parameter

```
name = %(Bob)
```

### As a function

```
» %(hello world)
hello world
```

### Nested quotes

```
» murex -c %(out %(Hello "${murex -c %(out %(Bob))}"))
Hello "Bob"
```

In this example we are calling Murex to execute code as a command line
parameter (the `-c` flag). That code outputs `Hello "..."` but inside the
double quotes is a name that is generated from a sub-shell. That sub-shell
itself runs another murex instance which also executes another command line
parameter, this time outputting the name **Bob**.

The example is contrived but it does demonstrate how you can heavily nest
quotes and even mix and match that with other quotation marks if desired.

This is something that is extremely difficult to write in traditional shells
because it would require lots of escaping, and even escaping the escape
characters (and so on) the further deep you get in your nest.

## Detail

### Multi-Line Quotes

Quotes can also work over multiple lines

```
» out %(foo
» bar)
foo
bar
```

### Legacy Support

Version 3.x of Murex introduced support for the `%` token, before that brace
quotes worked without it. However to retain backwards compatibility, the older
syntax is still supported...albeit officially classed as "deprecated" and may
be removed from a future release.

Below is a little detail about how the legacy syntax worked:

#### Deprecated Syntax

The open brace character is only recognized as a brace quote token if it is the
start of a parameter.

```
» set example=(World!)
» out (Hello $example)
Hello (World!)
```

## See Also

* [Code Block Parsing](../user-guide/code-block.md):
  Overview of how code blocks are parsed
* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [`"Double Quote"`](../parser/double-quote.md):
  Initiates or terminates a string (variables expanded)
* [`'Single Quote'`](../parser/single-quote.md):
  Initiates or terminates a string (variables not expanded)
* [`(brace quote)`](../parser/brace-quote-func.md):
  Write a string to the stdout without new line (deprecated)
* [`@Array` Sigil](../parser/array.md):
  Expand values as an array
* [`string` (stringing)](../types/str.md):
  string (primitive)
* [`{ Curly Brace }`](../parser/curly-brace.md):
  Initiates or terminates a code block
* [`~` Home Sigil](../parser/tilde.md):
  Home directory path variable

<hr/>

This document was generated from [gen/parser/quotes_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/quotes_doc.yaml).