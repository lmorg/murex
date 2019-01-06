# _murex_ Language Guide

## Command Reference: `err`

> Print a line to the STDERR

### Description

Write parameters to STDERR with a trailing new line character.

### Usage

    err: string to write -> <stderr>

### Examples

    » err Hello, World!
    Hello, World!

### Detail

`err` outputs as `string` data-type. This can be changed by casting

    err { "Code": 404, "Message": "Page not found" } ? cast json
    
However passing structured data-types along the STDERR stream is not recommended
as any other function within your code might also pass error messages along the
same stream and thus taint your structured data. This is why _murex_ does not
supply a `tout` function for STDERR. The recommended solution for passing
messages like these which you want separate from your STDOUT stream is to create
a new _murex_ named pipe.

    » pipe: --create messages
    » bg { <messages> -> pretty }
    » tout: <messages> json { "Code": 404, "Message": "Page not found" }
    » pipe: --close messages
    {
        "Code": 404,
        "Message": "Page not found"
    }
    
#### ANSI Constants

`err` supports ANSI constants.

### See Also

* [`(` (brace quote)](../docs/commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [`>>` (write to new or appended file)](../docs/commands/greater-than-greater-than.md):
  Writes STDIN to disk - appending contents if file already exists
* [`>` (write to new or truncated file)](../docs/commands/greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists    
* [`out`](../docs/commands/out.md):
  `echo` a string to the STDOUT with a trailing new line character
* [`pt`](../docs/commands/pt.md):
  Pipe telemetry. Writes data-types and bytes written
* [`tout`](../docs/commands/tout.md):
  Print a string to the STDOUT and set it's data-type
* [bg](../docs/commands/commands/bg.md):
  
* [cast](../docs/commands/commands/cast.md):
  
* [pipe](../docs/commands/commands/pipe.md):
  
* [pretty](../docs/commands/commands/pretty.md):
  
* [sprintf](../docs/commands/commands/sprintf.md):
  