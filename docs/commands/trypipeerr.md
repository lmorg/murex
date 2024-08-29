# `trypipeerr`

> Checks state of each function in a pipeline and exits block on error

## Description

`trypipeerr` checks the state of each function and exits the block if any of them
fail. Where `trypipeerr` differs from regular `tryerr` blocks is `trypipeerr` will
check every process along the pipeline as well as the terminating function (which
`tryerr` only validates against). The downside to this is that piped functions can
no longer run in parallel.

## Usage

```
trypipeerr { code-block } -> <stdout>

<stdin> -> trypipeerr { -> code-block } -> <stdout>
```

## Examples

```
trypipeerr {
    out "Hello, World!" -> grep: "non-existent string" -> cat
    out "This command will be ignored"
}
```

Formated pager (`less`) where the pager isn't called if the formatter (`pretty`) fails (eg input isn't valid JSON):

```
func pless {
    -> trypipeerr { -> pretty -> less }
}
```

## Detail

A failure is determined by:

* Any process that returns a non-zero exit number
* Any process that returns more output via stderr than it does via stdout

You can see which run mode your functions are executing under via the `fid-list`
command.

## See Also

* [Schedulers](../user-guide/schedulers.md):
  Overview of the different schedulers (or 'run modes') in Murex
* [`catch`](../commands/catch.md):
  Handles the exception code raised by `try` or `trypipe`
* [`if`](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [`runmode`](../commands/runmode.md):
  Alter the scheduler's behaviour at higher scoping level
* [`switch`](../commands/switch.md):
  Blocks of cascading conditionals
* [`try`](../commands/try.md):
  Handles non-zero exits inside a block of code
* [`tryerr`](../commands/tryerr.md):
  Handles errors inside a block of code
* [`trypipe`](../commands/trypipe.md):
  Checks for non-zero exits of each function in a pipeline
* [`unsafe`](../commands/unsafe.md):
  Execute a block of code, always returning a zero exit number
* [proc.list](../commands/fid-list.md):
  Lists all running functions within the current Murex session

<hr/>

This document was generated from [builtins/core/structs/tryerr_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/structs/tryerr_doc.yaml).