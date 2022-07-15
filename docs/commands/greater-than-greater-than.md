# _murex_ Shell Docs

## Command Reference: `>>` (append file)

> Writes STDIN to disk - appending contents if file already exists

## Description

Redirects output to file.

If a file already exists, the contents will be appended to existing contents.
Otherwise a new file is created.

## Usage

    <stdin> >> filename

## Examples

    g * >> files.txt

## Synonyms

* `>>`
* `fappend`


## See Also

* [parser/Arrow Pipe (`->`) Token](../parser/pipe-arrow.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [parser/POSIX Pipe (`|`) Token](../parser/pipe-posix.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [parser/STDERR Pipe (`?`) Token](../parser/pipe-err.md):
  Pipes STDERR from the left hand command to STDIN of the right hand command
* [commands/`<>` / `read-named-pipe`](../commands/namedpipe.md):
  Reads from a _murex_ named pipe
* [commands/`>` (truncate file)](../commands/greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists
* [commands/`g`](../commands/g.md):
  Glob pattern matching for file system objects (eg *.txt)
* [commands/`pipe`](../commands/pipe.md):
  Manage _murex_ named pipes
* [commands/`tmp`](../commands/tmp.md):
  Create a temporary file and write to it