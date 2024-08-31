# Function / Module Defaults (`runmode`)

> Alter the scheduler's behaviour at higher scoping level

## Description

Due to dynamic nature in which blocks are compiled on demand, traditional `try`
and `trypipe` blocks cannot affect the runtime behaviour of schedulers already
invoked (eg for function blocks and modules which `try` et al would sit inside).
To solve this we need an additional command that is executed by the compiler
prior to the block being executed which can define the runmode of the scheduler.
This is the purpose of `runmode`.

The caveat of being a compiler command rather than a builtin is that `runmode`
needs be the first command in a block.

## Usage

```
runmode try|trypipe function|module
```

## Examples

```
function hello {
    # Short conversation, exit on error
    
    runmode try function

    read name "What is your name? "
    out "Hello $name, pleased to meet you"
    
    read mood "How are you feeling? "
    out "I'm feeling $mood too"
}
```

## Detail

`runmode`'s parameters are ordered:

### 1st parameter

#### unsafe

Always return a zero exit number.

#### try

Checks only the last command in the pipeline for errors. However still allows
commands in a pipeline to run in parallel.

#### trypipe

Checks every command in the pipeline before executing the next. However this
blocks pipelines from running every command in parallel.

#### tryerr

Checks last command in the pipeline for errors (still allowing commands to run
in parallel) plus also checks if stderr contains excessive output.

#### trypipeerr

Checks every command in the pipeline before executing the next (blocking
commands from running in parallel) plus also checks if stderr contains
excessive output.

### 2nd parameter

#### function

Sets the runmode for all blocks within the function when `runmode` is placed at
the start of the function. This includes privates, autocompletes, events, etc.

#### module

Sets the runmode for all blocks within that module when placed at the start of
the module. This include any functions, privates, autocompletes, events, etc
that are inside that module. They do not need a separate `runmode ... function`
if `runmode ... module` is set.

## See Also

* [Caught Error Block (`catch`)](../commands/catch.md):
  Handles the exception code raised by `try` or `trypipe`
* [Disable Error Handling In Block (`unsafe`)](../commands/unsafe.md):
  Execute a block of code, always returning a zero exit number
* [Display Running Functions (`fid-list`)](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Pipe Fail (`trypipe`)](../commands/trypipe.md):
  Checks for non-zero exits of each function in a pipeline
* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Private Function (`private`)](../commands/private.md):
  Define a private function block
* [Public Function (`function`)](../commands/function.md):
  Define a function block
* [Read User Input (`read`)](../commands/read.md):
  `read` a line of input from the user and store as a variable
* [Schedulers](../user-guide/schedulers.md):
  Overview of the different schedulers (or 'run modes') in Murex
* [Stderr Checking In Pipes (`trypipeerr`)](../commands/trypipeerr.md):
  Checks state of each function in a pipeline and exits block on error
* [Stderr Checking In TTY (`tryerr`)](../commands/tryerr.md):
  Handles errors inside a block of code
* [Tab Autocompletion (`autocomplete`)](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [Try Block (`try`)](../commands/try.md):
  Handles non-zero exits inside a block of code
* [`event`](../commands/event.md):
  Event driven programming for shell scripts

<hr/>

This document was generated from [builtins/core/structs/try_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/structs/try_doc.yaml).