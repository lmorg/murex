# `->` Arrow Pipe

> Pipes stdout from the left hand command to stdin of the right hand command

## Description

This token behaves much like pipe would in Bash or similar shells. It passes
stdout along the pipeline while merging stderr stream with the parents stderr
stream.

`->` differs from `|` in the interactive terminal where it produces different
autocompletion suggestion. It returns a list of "methods". That is, commands
that are known to support the output type of the previous command. `->` helps
with the discovery of command line tools.

In shell scripts, `->` and `|` can be used interchangeably.



## Examples

### Piping stdout

```
» out Hello, world! -> regexp s/world/Earth/
Hello, Earth!

» out Hello, world!->regexp s/world/Earth/
Hello, Earth!
```

### Piping stderr

In following example the first command is writing to stderr rather than stdout
so `Hello, world!` doesn't get pipelined and thus isn't affected by `regexp`:

```
» err Hello, world! -> regexp s/world/Earth/
Hello, world!
```

To pipe stderr you'd need to use the `<!>` syntax. For example `<!out>` to
write stderr to stdout:

```
» err <!out> Hello, world! -> regexp s/world/Earth/
Hello, Earth!
```

## See Also

* [Error String (`err`)](../commands/err.md):
  Print a line to the stderr
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Read / Write To A Named Pipe (`<pipe>`)](../parser/namedpipe.md):
  Reads from a Murex named pipe
* [Regex Operations (`regexp`)](../commands/regexp.md):
  Regexp tools for arrays / lists of strings
* [`=>` Generic Pipe](../parser/pipe-generic.md):
  Pipes a reformatted stdout stream from the left hand command to stdin of the right hand command
* [`?` stderr Pipe](../parser/pipe-err.md):
  Pipes stderr from the left hand command to stdin of the right hand command (DEPRECATED)
* [`|` POSIX Pipe](../parser/pipe-posix.md):
  Pipes stdout from the left hand command to stdin of the right hand command

<hr/>

This document was generated from [gen/parser/pipes_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/pipes_doc.yaml).