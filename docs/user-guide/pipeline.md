# Pipeline

> Overview of what a "pipeline" is

## Description

In the Murex docs you'll often see the term "pipeline". This refers to any
commands sequenced together.

A pipeline can be joined via any pipe token (eg `|`, `->`, `=>`, `?`). But,
for the sake of documentation, a pipeline might even be a solitary command.

## Examples

Typical Murex pipeline:

```
open example.json -> [[ /node/0 ]]
```

Example of a single command pipeline:

```
top
```

Pipeline you might see in Bash / Zsh (this is also valid in Murex):

```
cat names.txt | sort | uniq
```

Pipeline filtering out a specific error from `example-cmd`

```
example-cmd ? grep "File not found"
```

## Detail

A pipeline isn't a Murex specific construct but rather something inherited
from Unix. Where Murex differs is that it can support sending typed
information to compatible functions (unlike standard Unix pipes which are
dumb-byte streams).

Wikipedia has a page on [Pipeline (Unix)](https://en.wikipedia.org/wiki/Pipeline_(Unix)):

> In Unix-like computer operating systems, a pipeline is a mechanism for
> inter-process communication using message passing. A pipeline is a set of
> processes chained together by their standard streams, so that the output
> text of each process (stdout) is passed directly as input (stdin) to the
> next one. The second process is started as the first process is still
> executing, and they are executed concurrently. The concept of pipelines was
> championed by Douglas McIlroy at Unix's ancestral home of Bell Labs, during
> the development of Unix, shaping its toolbox philosophy. It is named by
> analogy to a physical pipeline. A key feature of these pipelines is their
> "hiding of internals" (Ritchie & Thompson, 1974). This in turn allows for
> more clarity and simplicity in the system. 

## Named Pipes

The drawback with pipes is that it assumes each command runs sequentially one
after another and that everything fits neatly into the concept of "output" and
"errors". The moment you need to use background (`bg`) processes, do anything
more specific with data streams (even if just ignore them entirely), or use
more than one data stream, then this concept breaks down. This is where named
pipes come to the rescue. Named pipes are out of scope for this specific
document but you can read more on them in links the links below.

## See Also

* [Background Process (`bg`)](../commands/bg.md):
  Run processes in the background
* [Bang Prefix](../user-guide/bang-prefix.md):
  Bang prefixing to reverse default actions
* [Schedulers](../user-guide/schedulers.md):
  Overview of the different schedulers (or 'run modes') in Murex
* [`->` Arrow Pipe](../parser/pipe-arrow.md):
  Pipes stdout from the left hand command to stdin of the right hand command
* [`=>` Generic Pipe](../parser/pipe-generic.md):
  Pipes a reformatted stdout stream from the left hand command to stdin of the right hand command
* [`?` stderr Pipe](../parser/pipe-err.md):
  Pipes stderr from the left hand command to stdin of the right hand command (DEPRECATED)
* [`|` POSIX Pipe](../parser/pipe-posix.md):
  Pipes stdout from the left hand command to stdin of the right hand command

<hr/>

This document was generated from [gen/user-guide/pipeline_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/user-guide/pipeline_doc.yaml).