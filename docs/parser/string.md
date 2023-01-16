# _murex_ Shell Docs

## Parser Reference: String (`$`) Token

> Expand values as a string

## Description

The string token is used to tell _murex_ to expand variables and subshells as a
string (ie one single parameter) irrespective of the data that is stored in the
string. One handy common use case is file names where traditional POSIX shells
would treat spaces as a new file, whereas _murex_ treats spaces as a printable
character unless explicitly told to do otherwise.

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
    
The string token can also be used as a command too

    » set: example="Hello World!"
    
    » $example
    Hello World!

## Detail

Variables and subshells can be expanded inside double quotes, brace quotes as
well as used naked. But they cannot be expanded inside single quotes.

    » set: example="World!"
    
    » out: Hello $example
    Hello World!
    
    » out: 'Hello $example'
    Hello $example
    
    » out: "Hello $example"
    Hello World!
    
    » out: (Hello $example)
    Hello World!

## See Also

* [parser/Array (`@`) Token](../parser/array.md):
  Expand values as an array
* [parser/Brace Quote (`(`, `)`) Tokens](../parser/brace-quote.md):
  Initiates or terminates a string (variables expanded)
* [parser/Double Quote (`"`) Token](../parser/double-quote.md):
  Initiates or terminates a string (variables expanded)
* [user-guide/Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by _murex_
* [parser/Single Quote (`'`) Token](../parser/single-quote.md):
  Initiates or terminates a string (variables not expanded)
* [parser/Tilde (`~`) Token](../parser/tilde.md):
  Home directory path variable
* [commands/`(` (brace quote)](../commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [commands/`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [commands/`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable (deprecated)
* [commands/`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [commands/`set`](../commands/set.md):
  Define a local variable and set it's value