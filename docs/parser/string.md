# _murex_ Shell Docs

## Parser Reference: String (`@`) Token

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
    
In this example the second command is passing **foo\nbar** (`\n` escaped as a new
line) to `out`. The third command is passing an array of two values: **foo** and
**bar**.

The string token also works for subshells

    » out: ${ ja: [Mon..Fri] }
    ["Mon","Tue","Wed","Thu","Fri"]
    
And it can also function as a command too

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

* [commands/`(` (brace quote)](../commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [commands/`ja`](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [commands/`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [commands/`set`](../commands/set.md):
  Define a local variable and set it's value
* [parser/brace-quote](../parser/brace-quote.md):
  
* [parser/double-quote](../parser/double-quote.md):
  
* [parser/single-quote](../parser/single-quote.md):
  
* [parser/variable](../parser/variable.md):
  