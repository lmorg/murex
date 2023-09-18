# `|>` (truncate file)

> Writes STDIN to disk - overwriting contents if file already exists

## Description

Redirects output to file.

If a file already exists, the contents will be truncated (overwritten).
Otherwise a new file is created.

## Usage

```
<stdin> |> filename
```## Examples

```
g * |> files.txt
```

## Synonyms

* `>`
* `fwrite`


## See Also

* [`->` Arrow Pipe](../parser/pipe-arrow.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [`<read-named-pipe>`](../parser/namedpipe.md):
  Reads from a Murex named pipe
* [`>>` (append file)](../parser/greater-than-greater-than.md):
  Writes STDIN to disk - appending contents if file already exists
* [`?` STDERR Pipe](../parser/pipe-err.md):
  Pipes STDERR from the left hand command to STDIN of the right hand command
* [`|` POSIX Pipe](../parser/pipe-posix.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [g](../parser/g.md):
  
* [pipe](../parser/pipe.md):
  
* [tmp](../parser/tmp.md):
  

<hr/>

This document was generated from [builtins/core/io/file_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/io/file_doc.yaml).