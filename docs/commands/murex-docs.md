# _murex_ Shell Docs

## Command Reference: `murex-docs`

> Displays the man pages for _murex_ builtins

## Description

Displays the man pages for _murex_ builtins.

## Usage

    murex-docs: [ flags ] command -> <stdout>

## Examples

    # Output this man page
    murex-docs: murex-docs

## Flags

* `--summary`
    Returns an abridged description of the command rather than the entire help page.

## Detail

These man pages are compiled into the _murex_ executable.

## See Also

* [commands/`(` (brace quote)](../commands/brace-quote.md):
  Write a string to the STDOUT without new line
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
* [commands/`tout`](../commands/tout.md):
  Print a string to the STDOUT and set it's data-type
* [commands/`tread`](../commands/tread.md):
  `read` a line of input from the user and store as a user defined *typed* variable
* [commands/sprintf](../commands/sprintf.md):
  