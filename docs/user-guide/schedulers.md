# User Guide: Schedulers

> Overview of the different schedulers (or 'run modes') in _murex_

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
_murex_ validates the stderr and exit status of the last command in any pipeline.

If stderr is greater than stdout (per bytes written) **OR** the exit status is
non-zero then the scheduler exits that entire block.

## Try Pipe

This runs the commands sequentially because the stderr and the exit status of
each command is checked irrespective of whether that command is at the start of
the pipeline (eg `start -> middle -> end`), or anywhere else.

Like with `try`, if stderr is greater than stdout (per bytes written) **OR**
the exit status is non-zero then the scheduler exits that entire block. Unlike
with `try`, this check happens on every command rather than the last command in
the pipeline. 

## See Also

* [Arrow Pipe (`->`) Token](../parser/pipe-arrow.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [Generic Pipe (`=>`) Token](../parser/pipe-generic.md):
  Pipes a reformatted STDOUT stream from the left hand command to STDIN of the right hand command
* [POSIX Pipe (`|`) Token](../parser/pipe-posix.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [STDERR Pipe (`?`) Token](../parser/pipe-err.md):
  Pipes STDERR from the left hand command to STDIN of the right hand command
* [`runmode`](../commands/runmode.md):
  Alter the scheduler's behaviour at higher scoping level
* [`try`](../commands/try.md):
  Handles errors inside a block of code
* [`trypipe`](../commands/trypipe.md):
  Checks state of each function in a pipeline and exits block on error