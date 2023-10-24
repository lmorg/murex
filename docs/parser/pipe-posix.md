# `|` POSIX Pipe

> Pipes STDOUT from the left hand command to STDIN of the right hand command

## Description

This token behaves much like pipe would in Bash or similar shells. It passes
STDOUT along the pipeline while merging STDERR stream with the parents STDERR
stream.

It can be used interchangeably with the arrow pipe, `->`, in shell scripts.



## Examples

```
» out Hello, world! | regexp s/world/Earth/
Hello, Earth!

» out Hello, world!|regexp s/world/Earth/
Hello, Earth!
```

In this example the first command is writing to STDERR rather than STDOUT so
`Hello, world!` doesn't get pipelined and thus isn't affected by `regexp`:

```
» err Hello, world! | regexp s/world/Earth/
Hello, world!
```

## See Also

* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [`->` Arrow Pipe](../parser/pipe-arrow.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [`<read-named-pipe>`](../parser/namedpipe.md):
  Reads from a Murex named pipe
* [`=>` Generic Pipe](../parser/pipe-generic.md):
  Pipes a reformatted STDOUT stream from the left hand command to STDIN of the right hand command
* [`?` STDERR Pipe](../parser/pipe-err.md):
  Pipes STDERR from the left hand command to STDIN of the right hand command (DEPRECATED)
* [`err`](../commands/err.md):
  Print a line to the STDERR
* [`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [`regexp`](../commands/regexp.md):
  Regexp tools for arrays / lists of strings

<hr/>

This document was generated from [gen/parser/pipes_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/pipes_doc.yaml).