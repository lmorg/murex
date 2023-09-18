# `?` STDERR Pipe

> Pipes STDERR from the left hand command to STDIN of the right hand command

## Description

This token swaps the STDOUT and STDERR streams of the left hand command.

Please note that this token is only effective when it is prefixed by white
space. 

## Examples

```
» err Hello, world! ? regexp s/world/Earth/
Hello, Earth!
```

In following example the first command is writing to STDOUT rather than STDERR
so `Hello, world!` doesn't get pipelined and thus isn't affected by `regexp`:

```
» out Hello, world! ? regexp s/world/Earth/
Hello, world!
```

In following example the STDERR token isn't whitespace padded so is treated
like any ordinary printable character:

```
» err Hello, world!? regexp s/world/Earth/
Hello, world!? regexp s/world/Earth/
```

## See Also

* [`->` Arrow Pipe](../parser/pipe-arrow.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [`<read-named-pipe>`](../parser/namedpipe.md):
  Reads from a Murex named pipe
* [`=>` Generic Pipe](../parser/pipe-generic.md):
  Pipes a reformatted STDOUT stream from the left hand command to STDIN of the right hand command
* [`|` POSIX Pipe](../parser/pipe-posix.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [err](../parser/err.md):
  
* [out](../parser/out.md):
  
* [pipeline](../parser/pipeline.md):
  
* [regexp](../parser/regexp.md):
  

<hr/>

This document was generated from [gen/parser/pipes_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/pipes_doc.yaml).