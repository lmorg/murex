# _murex_ Language Guide

## Command Reference: `>` (write to new or truncated file)

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

* [`>>` (write to new or appended file)](../commands/greater-than-greater-than.md):
  Writes STDIN to disk - appending contents if file already exists
* [`g`](../commands/g.md):
  Glob pattern matching for file system objects (eg *.txt)
* [pipe](../commands/pipe.md):
  
* [tmp](../commands/tmp.md):
  