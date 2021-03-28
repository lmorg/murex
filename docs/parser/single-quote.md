# _murex_ Shell Docs

## Parser Reference: Single Quote (`'`) Token

> Initiates or terminates a string (variables not expanded)

## Description

Single quote is used to initiate and terminate strict strings where variables
cannot be expanded.

Commands can also be quoted using single quotes (eg where a command might
contain a space character in it's name)



## Examples

    » set: example='World!'
    
    » out: 'Hello $example'
    Hello $example

## Detail

Quotes can also work over multiple lines

    » out: 'foo
    » bar'
    foo
    bar

## See Also

* [parser/Array (`@`) Token](../parser/array.md):
  Expand values as an array
* [parser/Brace Quote (`(`, `)`) Token](../parser/brace-quote.md):
  Initiates or terminates a string (variables expanded)
* [parser/Double Quote (`"`) Token](../parser/double-quote.md):
  Initiates or terminates a string (variables expanded)
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