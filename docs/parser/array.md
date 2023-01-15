# _murex_ Shell Docs

## Parser Reference: Array (`@`) Token

> Expand values as an array

## Description

The array token is used to tell _murex_ to expand the string as multiple
parameters (an array) rather than as a single parameter string.

## Examples

    » set: example="foo\nbar"
    
    » out: $example
    foo
    bar
    
    » out: @example
    foo bar
    
In this example the second command is passing `foo\nbar` (`\n` escaped as a new
line) to `out`. The third command is passing an array of two values: `foo` and
`bar`.

The string and array tokens also works for subshells

    » out: ${ ja: [Mon..Fri] }
    ["Mon","Tue","Wed","Thu","Fri"]
    
    » out: @{ ja: [Mon..Fri] }
    Mon Tue Wed Thu Fri

## Detail

Since arrays are expanded over multiple parameters, you cannot expand an array
inside quoted strings like you can with a string variable:

    » out: "foo ${ ja: [1..5] } bar"
    foo ["1","2","3","4","5"] bar
    
    » out: "foo @{ ja: [1..5] } bar"
    foo  1 2 3 4 5  bar
    
    » (${ ja: [1..5] })
    ["1","2","3","4","5"]   
    
    » (@{ ja: [1..5] })
    @{ ja: [1..5] } 

## See Also

* [parser/Brace Quote (`(`, `)`) Tokens](../parser/brace-quote.md):
  Initiates or terminates a string (variables expanded)
* [parser/Double Quote (`"`) Token](../parser/double-quote.md):
  Initiates or terminates a string (variables expanded)
* [parser/Single Quote (`'`) Token](../parser/single-quote.md):
  Initiates or terminates a string (variables not expanded)
* [parser/String (`$`) Token](../parser/string.md):
  Expand values as a string
* [parser/Tilde (`~`) Token](../parser/tilde.md):
  Home directory path variable
* [commands/`(` (brace quote)](../commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [commands/`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [commands/`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [commands/`set`](../commands/set.md):
  Define a local variable and set it's value