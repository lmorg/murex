# `>>` Append Pipe

> Redirects STDOUT to a file and append its contents

## Description

This is used to redirect the STDOUT of a command and append it to a file. If
that file does not exist, then the file is created.

This behaves similarly to the [Bash (et al) token](https://www.gnu.org/software/bash/manual/bash.html#Appending-Redirected-Output)
except it doesn't support adding alternative file descriptor numbers. Instead
you will need to use named pipes to achieve the same effect in Murex.

## Examples

```
» out "Hello" >> example.txt
» out "World!" >> example.txt
» open example.txt
Hello
World!
```

## Detail

This is just syntactic sugar for `-> >>`. Thus when the parser reads code like
the following:

```
echo "foobar" >> example.txt
```

it will compile an abstract syntax tree which would reflect the following code
instead:

```
echo "foobar" | >> example.txt
```

### Truncating a file

To truncate a file (ie overwrite its contents) use `|>` instead.

## See Also

* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [`->` Arrow Pipe](../parser/pipe-arrow.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [`<read-named-pipe>`](../parser/namedpipe.md):
  Reads from a Murex named pipe
* [`>>` (append file)](../parser/greater-than-greater-than.md):
  Writes STDIN to disk - appending contents if file already exists
* [`?` STDERR Pipe](../parser/pipe-err.md):
  Pipes STDERR from the left hand command to STDIN of the right hand command
* [`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [`|>` (truncate file)](../parser/greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists
* [`|` POSIX Pipe](../parser/pipe-posix.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command

<hr/>

This document was generated from [gen/parser/pipes_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/pipes_doc.yaml).