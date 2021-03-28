# _murex_ Shell Docs

## Parser Reference: Brace Quote (`(`, `)`) Token

> Initiates or terminates a string (variables expanded)

## Description

Brace quote is used to initiate and terminate strict strings where variables
can be expanded.

While brace quotes are untraditional compared to your typical string quotations
in POSIX shells, brace quotes have one advantage in that the open and close
grapheme differ (ie `(` is a different character to `)`). This brings benefits
when nesting quotes as it saves the developer from having to carefully escape
the nested quotation marks just the right number of times.

Commands cannot be quoted using double quotes because `(` is recognized as its
own command.



## Examples

The open brace character is only recognized as a brace quote token if it is the
start of a parameter.

    » set: example=(World!)
    
    » out: (Hello $example)
    Hello (World!)

## Detail

Quotes can also work over multiple lines

    » out: (foo
    » bar)
    foo
    bar

## See Also

* [parser/Array (`@`) Token](../parser/array.md):
  Expand values as an array
* [parser/Brace Quote (`(`, `)`) Token](../parser/brace-quote.md):
  Initiates or terminates a string (variables expanded)
* [parser/Single Quote (`'`) Token](../parser/single-quote.md):
  Initiates or terminates a string (variables not expanded)
* [parser/String (`@`) Token](../parser/string.md):
  Expand values as a string
* [parser/Tilde (`~`) Token](../parser/tilde.md):
  Home directory path variable
* [commands/`(` (brace quote)](../commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [commands/`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [commands/`set`](../commands/set.md):
  Define a local variable and set it's value