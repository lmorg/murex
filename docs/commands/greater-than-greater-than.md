# `>>` (append file)

> Writes STDIN to disk - appending contents if file already exists

## Description

Redirects output to file.

If a file already exists, the contents will be appended to existing contents.
Otherwise a new file is created.

## Usage

    `<stdin>` >> filename

## Examples

    g * >> files.txt

## Synonyms

- `>>`
- `fappend`

## See Also

- [Arrow Pipe (`->`) Token](/parser/pipe-arrow.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
- [POSIX Pipe (`|`) Token](/parser/pipe-posix.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
- [STDERR Pipe (`?`) Token](/parser/pipe-err.md):
  Pipes STDERR from the left hand command to STDIN of the right hand command
- [`<>` / `read-named-pipe`](./namedpipe.md):
  Reads from a Murex named pipe
- [`>` (truncate file)](./greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists
- [`g`](./g.md):
  Glob pattern matching for file system objects (eg `*.txt`)
- [`pipe`](./pipe.md):
  Manage Murex named pipes
- [`tmp`](./tmp.md):
  Create a temporary file and write to it
