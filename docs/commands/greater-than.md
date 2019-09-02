# _murex_ Shell Guide

## Command Reference: `>` (truncate file)

> Writes STDIN to disk - overwriting contents if file already exists

### Description

Redirects output to file.

If a file already exists, the contents will be truncated (overwritten).
Otherwise a new file is created.

### Usage

    <stdin> -> > filename

### Examples

    g * -> > files.txt

### Synonyms

* `>`


### See Also

* [commands/`>>` (append file)](../commands/greater-than-greater-than.md):
  Writes STDIN to disk - appending contents if file already exists
* [commands/`g`](../commands/g.md):
  Glob pattern matching for file system objects (eg *.txt)
* [commands/pipe](../commands/pipe.md):
  
* [commands/tmp](../commands/tmp.md):
  