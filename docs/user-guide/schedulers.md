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

## Try

This is similar to normal where commands in a pipeline are run in parallel except
Murex validates the stderr and exit status of the last command in any pipeline.

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

* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [pipe-arrow](../user-guide/pipe-arrow.md):
  
* [pipe-err](../user-guide/pipe-err.md):
  
* [pipe-generic](../user-guide/pipe-generic.md):
  
* [pipe-posix](../user-guide/pipe-posix.md):
  
* [runmode](../user-guide/runmode.md):
  
* [try](../user-guide/try.md):
  
* [trypipe](../user-guide/trypipe.md):
  

<hr/>

This document was generated from [gen/user-guide/schedulers_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/user-guide/schedulers_doc.yaml).