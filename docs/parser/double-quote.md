# `"Double Quote"`

> Initiates or terminates a string (variables expanded)

## Description

Double quote is used to initiate and terminate strict strings where variables
can be expanded.

Commands can also be quoted using double quotes (eg where a command might
contain a space character in it's name) however variables cannot be used as
part of a command name.

## Examples

```
» set: example="World!"

» out: "Hello $example"
Hello World!
```

## Detail

Quotes can also work over multiple lines

```
» out "foo
» bar"
foo
bar
```

## See Also

* [Array (`@`) Token](../parser/array.md):
  Expand values as an array
* [Tilde (`~`) Token](../parser/tilde.md):
  Home directory path variable
* [`$variable`](../parser/string.md):
  Expand values as a string
* [`'Single Quote'`](../parser/single-quote.md):
  Initiates or terminates a string (variables not expanded)
* [`(brace quote)`](../parser/brace-quote.md):
  Write a string to the STDOUT without new line
* [`(brace quote)`](../parser/brace-quote.md):
  Write a string to the STDOUT without new line
* [`{Curly Brace}`](../parser/curly-brace.md):
  Initiates or terminates a code block
* [out](../parser/out.md):
  
* [set](../parser/set.md):
  

<hr/>

This document was generated from [gen/parser/quotes_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/quotes_doc.yaml).