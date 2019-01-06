# _murex_ Language Guide

## Command Reference: `(` (brace quote)

> Write a string to the STDOUT without new line

### Description

Write parameters to STDOUT (does not include a new line)

### Usage

    (string to write) -> <stdout>

### Examples

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

### Detail

The `(` function performs exactly like the `(` token for quoting so you do not
need to escape other tokens (eg single / double quotes, `'`/`"`, nor curly
braces, `{}`). However the braces are nestable so you will need to escape those
characters if you don't want them nested.

#### ANSI Constants

`(` supports ANSI constants.

### Synonyms

* `(`


### See Also

* [`>>` (write to new or appended file)](../docs/commands/greater-than-greater-than.md):
  Writes STDIN to disk - appending contents if file already exists
* [`>` (write to new or truncated file)](../docs/commands/greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists    
* [`err`](../docs/commands/err.md):
  Print a line to the STDERR
* [`out`](../docs/commands/out.md):
  `echo` a string to the STDOUT with a trailing new line character
* [`pt`](../docs/commands/pt.md):
  Pipe telemetry. Writes data-types and bytes written
* [`tout`](../docs/commands/tout.md):
  Print a string to the STDOUT and set it's data-type
* [cast](../docs/commands/commands/cast.md):
  
* [sprintf](../docs/commands/commands/sprintf.md):
  