# _murex_ Language Guide

## Command Reference: `murex-docs`

> Displays the man pages for _murex_ builtins

### Description

Displays the man pages for _murex_ builtins.

### Usage

    murex-docs: [ flag ] command -> <stdout>

### Examples

    # Output this man page
    murex-docs: murex-docs

### Flags

* `--digest`
    returns an abridged description of the command rather than the entire help page.

### Detail

These man pages are compiled into the _murex_ executable.

### See Also

* [`(` (brace quote)](../docs/commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [`>>` (write to new or appended file)](../docs/commands/greater-than-greater-than.md):
  Writes STDIN to disk - appending contents if file already exists
* [`>` (write to new or truncated file)](../docs/commands/greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists    
* [`err`](../docs/commands/err.md):
  Print a line to the STDERR
* [`out`](../docs/commands/out.md):
  `echo` a string to the STDOUT with a trailing new line character
* [`tout`](../docs/commands/tout.md):
  Print a string to the STDOUT and set it's data-type
* [`tread`](../docs/commands/tread.md):
  `read` a line of input from the user and store as a user defined *typed* variable    
* [cast](../docs/commands/commands/cast.md):
  
* [sprintf](../docs/commands/commands/sprintf.md):
  