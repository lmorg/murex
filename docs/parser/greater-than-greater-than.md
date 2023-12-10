# `>>` Append File

> Writes STDIN to disk - appending contents if file already exists

## Description

Redirects output to file.

If a file already exists, the contents will be appended to existing contents.
Otherwise a new file is created.

## Usage

```
<stdin> >> filename
```

## Examples

```
g * >> files.txt
```

## Synonyms

* `>>`
* `fappend`


## See Also

* [`->` Arrow Pipe](../parser/pipe-arrow.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [`<read-named-pipe>`](../parser/namedpipe.md):
  Reads from a Murex named pipe
* [`?` STDERR Pipe](../parser/pipe-err.md):
  Pipes STDERR from the left hand command to STDIN of the right hand command (DEPRECATED)
* [`g`](../commands/g.md):
  Glob pattern matching for file system objects (eg `*.txt`)
* [`pipe`](../commands/pipe.md):
  Manage Murex named pipes
* [`tmp`](../commands/tmp.md):
  Create a temporary file and write to it
* [`|>` Truncate File](../parser/greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists
* [`|` POSIX Pipe](../parser/pipe-posix.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command

<hr/>

This document was generated from [builtins/core/io/file_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/io/file_doc.yaml).