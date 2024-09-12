# Disable Error Handling In Block (`unsafe`)

> Execute a block of code, always returning a zero exit number

## Description

`unsafe` is similar to normal execution except that the exit number for the
last function in the `unsafe` block is ignored. `unsafe` always returns `0`.

This is useful in any situations where you might want errors ignored.

## Usage

```
unsafe { code-block } -> <stdout>

<stdin> -> unsafe { -> code-block } -> <stdout>
```

## Examples

```
try {
    unsafe { err "foobar" }
    out "This message still displays because the error is inside an `unsafe` block"
}
```

## See Also

* [Caught Error Block (`catch`)](../commands/catch.md):
  Handles the exception code raised by `try` or `trypipe`
* [Display Running Functions (`fid-list`)](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [Function / Module Defaults (`runmode`)](../commands/runmode.md):
  Alter the scheduler's behaviour at higher scoping level
* [If Conditional (`if`)](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [Pipe Fail (`trypipe`)](../commands/trypipe.md):
  Checks for non-zero exits of each function in a pipeline
* [Schedulers](../user-guide/schedulers.md):
  Overview of the different schedulers (or 'run modes') in Murex
* [Stderr Checking In Pipes (`trypipeerr`)](../commands/trypipeerr.md):
  Checks state of each function in a pipeline and exits block on error
* [Stderr Checking In TTY (`tryerr`)](../commands/tryerr.md):
  Handles errors inside a block of code
* [Switch Conditional (`switch`)](../commands/switch.md):
  Blocks of cascading conditionals
* [Try Block (`try`)](../commands/try.md):
  Handles non-zero exits inside a block of code

<hr/>

This document was generated from [builtins/core/structs/try_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/structs/try_doc.yaml).