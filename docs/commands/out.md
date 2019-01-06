# _murex_ Language Guide

## Command Reference: `out`

> `echo` a string to the STDOUT with a trailing new line character

### Description

Write parameters to STDOUT with a trailing new line character.

### Usage

    out: string to write -> <stdout>

### Examples

    » out Hello, World!
    Hello, World!
    
    # For compatibility with other shells, `echo` is also supported:
    » echo Hello, World!
    Hello, World!

### Detail

`out` / `echo` output as `string` data-type. This can be changed by casting or
using the `tout` function.

#### ANSI Constants

`out` supports ANSI constants.

### Synonyms

* `out`
* `echo`


### See Also

* [`(` (brace quote)](../docs/commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [`>>` (write to new or appended file)](../docs/commands/greater-than-greater-than.md):
  Writes STDIN to disk - appending contents if file already exists
* [`>` (write to new or truncated file)](../docs/commands/greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists    
* [`err`](../docs/commands/err.md):
  Print a line to the STDERR
* [`pt`](../docs/commands/pt.md):
  Pipe telemetry. Writes data-types and bytes written
* [`read`](../docs/commands/read.md):
  `read` a line of input from the user and store as a variable
* [`tout`](../docs/commands/tout.md):
  Print a string to the STDOUT and set it's data-type
* [`tread`](../docs/commands/tread.md):
  `read` a line of input from the user and store as a user defined *typed* variable    
* [cast](../docs/commands/commands/cast.md):
  
* [sprintf](../docs/commands/commands/sprintf.md):
  