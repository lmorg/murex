# `esccli` - Command Reference

> Escapes an array so output is valid shell code

## Description

`esccli` takes an array and escapes any characters that might cause problems
when pasted back into the terminal. Typically you'd want to use this against
command parameters.

## Usage

    <stdin> -> esccli -> <stdout>
    
    esccli @array -> <stdout>

## Examples

As a method

    » alias foobar=out 'foo$b@r'
    » alias -> [foobar]
    [
        "out",
        "foo$b@r"
    ]
    » alias -> [foobar] -> esccli
    out foo\$b\@r
    
As a function

    » alias -> [foobar] -> set: fb
    » $fb
    ["out","foo$b@r"]
    » esccli: @fb
    out foo\$b\@r

## See Also

* [`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [`alias`](../commands/alias.md):
  Create an alias for a command
* [`escape`](../commands/escape.md):
  Escape or unescape input 
* [`eschtml`](../commands/eschtml.md):
  Encode or decodes text for HTML
* [`escurl`](../commands/escurl.md):
  Encode or decodes text for the URL
* [`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character