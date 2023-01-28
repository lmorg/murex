# _murex_ Shell Docs

## Parser Reference

This section is a glossary of _murex_ tokens and parser behavior.

## Other Reference Material

### Language Guides

1. [GUIDE.builtin-functions](./GUIDE.builtin-functions.md), for docs
on the core builtins.

2. [GUIDE.control-structures](./GUIDE.control-structures.md), which
contains builtins required for building logic.

### _murex_'s Source Code

The parser is located _murex_'s source under the `lang/` path of the project
files.

## Pages

* [And (`&&`) Logical Operator](parser/logical-and.md):
  Continues next operation if previous operation passes
* [Append Pipe (`>>`) Token](parser/pipe-append.md):
  Redirects STDOUT to a file and append its contents
* [Array (`@`) Token](parser/array.md):
  Expand values as an array
* [Arrow Pipe (`->`) Token](parser/pipe-arrow.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [Brace Quote (`%(`, `)`) Tokens](parser/brace-quote.md):
  Initiates or terminates a string (variables expanded)
* [Curly Brace (`{`, `}`) Tokens](parser/curly-brace.md):
  Initiates or terminates a code block
* [Double Quote (`"`) Token](parser/double-quote.md):
  Initiates or terminates a string (variables expanded)
* [Generic Pipe (`=>`) Token](parser/pipe-generic.md):
  Pipes a reformatted STDOUT stream from the left hand command to STDIN of the right hand command
* [Or (`||`) Logical Operator](parser/logical-or.md):
  Continues next operation only if previous operation fails
* [POSIX Pipe (`|`) Token](parser/pipe-posix.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [STDERR Pipe (`?`) Token](parser/pipe-err.md):
  Pipes STDERR from the left hand command to STDIN of the right hand command
* [Single Quote (`'`) Token](parser/single-quote.md):
  Initiates or terminates a string (variables not expanded)
* [String (`$`) Token](parser/string.md):
  Expand values as a string
* [Tilde (`~`) Token](parser/tilde.md):
  Home directory path variable