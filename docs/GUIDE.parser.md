# _murex_ Shell Docs

## Parser Reference

This section is a glossary of _murex_ tokens and parser behavior.

## Other Reference Material

### Language Guides

1. [GUIDE.builtin-functions.md](./GUIDE.builtin-functions.md), for docs
on the core builtins.

2. [GUIDE.control-structures.md](./GUIDE.control-structures.md), which
contains builtins required for building logic.

### _murex_'s Source Code

The parser is located _murex_'s source under the `lang/` path of the project
files.

## Pages

* [Array (`@`) Token](parser/array.md):
  Expand values as an array
* [Arrow Pipe (`->`) Token](parser/pipearrow.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [Brace Quote (`(`, `)`) Token](parser/brace-quote.md):
  Initiates or terminates a string (variables expanded)
* [Double Quote (`"`) Token](parser/double-quote.md):
  Initiates or terminates a string (variables expanded)
* [Formatted Pipe (`=>`) Token](parser/pipeformat.md):
  Pipes a reformatted STDOUT stream from the left hand command to STDIN of the right hand command
* [POSIX Pipe (`|`) Token](parser/pipeposix.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [STDERR Pipe (`?`) Token](parser/pipeerr.md):
  Pipes STDERR from the left hand command to STDIN of the right hand command
* [Single Quote (`'`) Token](parser/single-quote.md):
  Initiates or terminates a string (variables not expanded)
* [String (`@`) Token](parser/string.md):
  Expand values as a string
* [Tilde (`~`) Token](parser/tilde.md):
  Home directory path variable