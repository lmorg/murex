# `>>` Append File

> Writes stdin to disk - appending contents if file already exists

## Description

This is used to redirect the stdout of a command and append it to a file. If
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
out "foobar" >> example.txt
```

it will compile an abstract syntax tree which would reflect the following code
instead:

```
out "foobar" | >> example.txt
```

### Truncating a file

To truncate a file (ie overwrite its contents) use `|>` instead.

## Synonyms

* `>>`
* `fappend`


## See Also

* [Create Named Pipe (`pipe`)](../commands/pipe.md):
  Manage Murex named pipes
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Read / Write To A Named Pipe (`<pipe>`)](../parser/namedpipe.md):
  Reads from a Murex named pipe
* [Truncate File (`>`)](../parser/file-truncate.md):
  Writes stdin to disk - overwriting contents if file already exists
* [`->` Arrow Pipe](../parser/pipe-arrow.md):
  Pipes stdout from the left hand command to stdin of the right hand command
* [`|` POSIX Pipe](../parser/pipe-posix.md):
  Pipes stdout from the left hand command to stdin of the right hand command

<hr/>

This document was generated from [gen/parser/pipes_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/pipes_doc.yaml).