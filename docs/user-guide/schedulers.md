# Schedulers

> Overview of the different schedulers (or 'run modes') in Murex

There are a few distinct schedulers (or run modes) in Murex which are invoked
by builtin commands. This means you can alter the way commands are executed
dynamically within Murex shell scripts.

## Normal

This is a traditional shell where anything in a pipeline (eg `cmd1 -> cmd2 -> cmd3`)
is executed in parallel. The scheduler only pauses launching new commands when
the last command in any pipeline is still executing. A pipeline could be multiple
commands (like above) or a single command (eg `top`).

**Normal** is the default run mode. When running in a stricter mode, you can not
reset the run mode back to **normal**. You can, however, switch to `unsafe`.

## Unsafe

This is the same as **normal** except that `unsafe` blocks always return zero
exit numbers. The purpose for this is to enable "normal" like scheduling inside
stricter code blocks that might exit if the last function was a non-zero exit
number.

## Try

This is the weakest of all the stricter run modes. It does check the exit number
to confirm if the last function was successful, but only the last function in
any given pipeline. So in `cmd1 -> cmd2 -> cmd3`, if `cmd1` or `cmd2` fail, the
`try` block doesn't exit.

The benefit of run mode is that it still supports commands running in parallel.

## Try Pipe

This runs the commands sequentially because the stderr and the exit number of
each command is checked irrespective of whether that command is at the start of
the pipeline (eg `start -> middle -> end`), or anywhere else.

This offers better coverage of exit numbers but at the cost of parallelisation.

## Try Err

This is similar to `try` and **normal** where commands in a pipeline are run in
parallel. The key difference with `tryerr` is that  Murex validates the stderr
as well as the exit number of the last command in any pipeline.

If stderr is greater than stdout (per bytes written) **OR** the exit number is
non-zero then the scheduler exits that entire block.

## Try Pipe Err

This runs the commands sequentially because the stderr and the exit number of
each command is checked irrespective of whether that command is at the start of
the pipeline (eg `start -> middle -> end`), or anywhere else.

Like with `tryerr`, if stderr is greater than stdout (per bytes written) **OR**
the exit number is non-zero then the scheduler exits that entire block. Unlike
with `tryerr`, this check happens on every command rather than the last command
in the pipeline. 

## See Also

* [Disable Error Handling In Block (`unsafe`)](../commands/unsafe.md):
  Execute a block of code, always returning a zero exit number
* [Function / Module Defaults (`runmode`)](../commands/runmode.md):
  Alter the scheduler's behaviour at higher scoping level
* [Pipe Fail (`trypipe`)](../commands/trypipe.md):
  Checks for non-zero exits of each function in a pipeline
* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Stderr Checking In Pipes (`trypipeerr`)](../commands/trypipeerr.md):
  Checks state of each function in a pipeline and exits block on error
* [Stderr Checking In TTY (`tryerr`)](../commands/tryerr.md):
  Handles errors inside a block of code
* [Try Block (`try`)](../commands/try.md):
  Handles non-zero exits inside a block of code
* [`->` Arrow Pipe](../parser/pipe-arrow.md):
  Pipes stdout from the left hand command to stdin of the right hand command
* [`=>` Generic Pipe](../parser/pipe-generic.md):
  Pipes a reformatted stdout stream from the left hand command to stdin of the right hand command
* [`?` stderr Pipe](../parser/pipe-err.md):
  Pipes stderr from the left hand command to stdin of the right hand command (DEPRECATED)
* [`|` POSIX Pipe](../parser/pipe-posix.md):
  Pipes stdout from the left hand command to stdin of the right hand command

<hr/>

This document was generated from [gen/user-guide/schedulers_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/user-guide/schedulers_doc.yaml).