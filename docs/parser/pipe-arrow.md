# `->` Arrow Pipe

> Pipes STDOUT from the left hand command to STDIN of the right hand command

## Description

This token behaves much like pipe would in Bash or similar shells. It passes
STDOUT along the pipeline while merging STDERR stream with the parents STDERR
stream.

`->` differs from `|` in the interactive terminal where it produces different
autocompletion suggestion. It returns a list of "methods". That is, commands
that are known to support the output type of the previous command. `->` helps
with the discovery of commandline tools.

In shell scripts, `->` and `|` can be used interchangeably.

## Examples

```
» out Hello, world! -> regexp s/world/Earth/
Hello, Earth!

» out Hello, world!->regexp s/world/Earth/
Hello, Earth!
```

In following example the first command is writing to STDERR rather than STDOUT
so `Hello, world!` doesn't get pipelined and thus isn't affected by `regexp`:

```
» err Hello, world! -> regexp s/world/Earth/
Hello, world!
```

## See Also

* [`<read-named-pipe>`](../parser/namedpipe.md):
  Reads from a Murex named pipe
* [`=>` Generic Pipe](../parser/pipe-generic.md):
  Pipes a reformatted STDOUT stream from the left hand command to STDIN of the right hand command
* [`?` STDERR Pipe](../parser/pipe-err.md):
  Pipes STDERR from the left hand command to STDIN of the right hand command
* [`|` POSIX Pipe](../parser/pipe-posix.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [err](../parser/err.md):
  
* [out](../parser/out.md):
  
* [pipeline](../parser/pipeline.md):
  
* [regexp](../parser/regexp.md):
  

<hr/>

This document was generated from [gen/parser/pipes_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/pipes_doc.yaml).