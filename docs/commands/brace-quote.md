# _murex_ Shell Docs

## Command Reference: `(` (brace quote)

> Write a string to the STDOUT without new line

## Description

Write parameters to STDOUT (does not include a new line)

## Usage

    (string to write) -> <stdout>

## Examples

    » (Hello, World!)
    Hello, World!
    
    » (Hello,\nWorld!)
    Hello,
    World!
    
    » ((Hello,) (World!))
    (Hello,) (World!)
    
    # Print "Hello, World!" in red text
    » {RED}Hello, World!{RESET}
    Hello, World!

## Detail

The `(` function performs exactly like the `(` token for quoting so you do not
need to escape other tokens (eg single / double quotes, `'`/`"`, nor curly
braces, `{}`). However the braces are nestable so you will need to escape those
characters if you don't want them nested.

### ANSI Constants

`(` supports ANSI constants.

## Synonyms

* `(`


## See Also

* [commands/`>>` (append file)](../commands/greater-than-greater-than.md):
  Writes STDIN to disk - appending contents if file already exists
* [commands/`>` (truncate file)](../commands/greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists
* [commands/`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [commands/`err`](../commands/err.md):
  Print a line to the STDERR
* [commands/`out`](../commands/out.md):
  `echo` a string to the STDOUT with a trailing new line character
* [commands/`pt`](../commands/pt.md):
  Pipe telemetry. Writes data-types and bytes written
* [commands/`tout`](../commands/tout.md):
  Print a string to the STDOUT and set it's data-type
* [userguide/ansi](../userguide/ansi.md):
  
* [commands/sprintf](../commands/sprintf.md):
  