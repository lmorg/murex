# _murex_ Shell Docs

## Parser Reference: Schedulers

> Overview of the different schedulers (or run modes) in _murex_

## Description

There are a few distinct schedulers (or run modes) in _murex_ which are invoked
by builtin commands. This means you can alter the way commands are executed
dynamically within _murex_ shell scripts.

## Normal

This is a traditional shell where anything in a pipeline (eg `cmd1 -> cmd2 -> cmd3`)
is executed in parallel. The scheduler only pauses launching new commands when
the last command in any pipeline is still executing. A pipeline could be multiple
commands (like above) or a single command (eg `top`).

## Try

This is similar to normal where commands in a pipeline are run in parallel except
_murex_ validates the STDERR and exit status of the last command in any pipeline.

If STDERR is greater than STDOUT (per bytes written) **OR** the exit status is
non-zero then the scheduler exits that entire block.

## Try Pipe

This runs the commands sequentially because the STDERR and the exit status of
each command is checked irrespective of whether that command is at the start of
the pipeline (eg `start -> middle -> end`), or anywhere else.

Like with `try`, if STDERR is greater than STDOUT (per bytes written) **OR**
the exit status is non-zero then the scheduler exits that entire block. Unlike
with `try`, this check happens on every command rather than the last command in
the pipeline. 



## See Also

* [parser/Pipeline](../parser/pipeline.md):
  Overview of what a "pipeline" is
* [commands/`try`](../commands/try.md):
  Handles errors inside a block of code
* [commands/`trypipe`](../commands/trypipe.md):
  Checks state of each function in a pipeline and exits block on error