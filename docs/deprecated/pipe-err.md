# `?` stderr Pipe

**This feature has been deprecated** and thus the following documentation is
provided for historical reference rather than recommendations for new code.

While we do make every effort to maintain backwards compatibility, sometimes
deprecations need to be made in order to keep Murex focused and maintainable.

Please read our [compatibility commitment](https://murex.rocks/compatibility.html)
for more information on how we approach such changes. Visit the [deprecated section](https://github.com/lmorg/murex/tree/master/docs/deprecated)
if you need to view other deprecated features.


> Pipes stderr from the left hand command to stdin of the right hand command (removed 8.0)

## Description

This token swaps the stdout and stderr streams of the left hand command.

Please note that this token is only effective when it is prefixed by white
space.

> This feature has been deprecated. Please use `<err> <!out>` instead. For example:
> ```
> command <err> <!out> parameter-1 parameter-2 -> next-command parameter-1
> ```



## Examples

```
» err Hello, world! ? regexp s/world/Earth/
Hello, Earth!
```

In following example the first command is writing to stdout rather than stderr
so `Hello, world!` doesn't get pipelined and thus isn't affected by `regexp`:

```
» out Hello, world! ? regexp s/world/Earth/
Hello, world!
```

In following example the stderr token isn't whitespace padded so is treated
like any ordinary printable character:

```
» err Hello, world!? regexp s/world/Earth/
Hello, world!? regexp s/world/Earth/
```

## See Also

* [Error String, strerr: `err`](../commands/err.md):
  Print a line to the stderr
* [Output String, stdout: `out`](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Read / Write To A Named Pipe: `<pipe>`](../parser/namedpipe.md):
  Reads from a Murex named pipe
* [Regex Patterns: `regexp`](../commands/regexp.md):
  Regexp tools for arrays / lists of strings
* [`->` Arrow Pipe](../parser/pipe-arrow.md):
  Pipes stdout from the left hand command to stdin of the right hand command
* [`=>` Generic Pipe](../parser/pipe-generic.md):
  Pipes a reformatted stdout stream from the left hand command to stdin of the right hand command
* [`|` POSIX Pipe](../parser/pipe-posix.md):
  Pipes stdout from the left hand command to stdin of the right hand command

<hr/>

This document was generated from [gen/parser/pipes_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/pipes_doc.yaml).