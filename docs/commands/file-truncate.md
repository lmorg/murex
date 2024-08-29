# fs.truncate (`>`)

> Writes stdin to disk - overwriting contents if file already exists

## Description

Redirects output to file.

If a file already exists, the contents will be truncated (overwritten).
Otherwise a new file is created.

## Usage

```
<stdin> |> filename
```

## Examples

```
g * |> files.txt
```

## Synonyms

* `>`
* `fs.truncate`
* `|>`
* `fwrite`


## See Also

* [`->` Arrow Pipe](../parser/pipe-arrow.md):
  Pipes stdout from the left hand command to stdin of the right hand command
* [`>>` Append File](../parser/file-append.md):
  Writes stdin to disk - appending contents if file already exists
* [`?` stderr Pipe](../parser/pipe-err.md):
  Pipes stderr from the left hand command to stdin of the right hand command (DEPRECATED)
* [`|` POSIX Pipe](../parser/pipe-posix.md):
  Pipes stdout from the left hand command to stdin of the right hand command
* [fs.glob (`g`)](../commands/g.md):
  Glob pattern matching for file system objects (eg `*.txt`)
* [fs.tmpfile (`tmp`)](../commands/tmp.md):
  Create a temporary file and write to it
* [io.new.pipe](../commands/pipe.md):
  Manage Murex named pipes
* [io.pipe (`<pipe>`)](../commands/namedpipe.md):
  Reads from a Murex named pipe

<hr/>

This document was generated from [builtins/core/io/file_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/io/file_doc.yaml).