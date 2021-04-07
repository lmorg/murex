# _murex_ Shell Docs

## Command Reference: `trypipe`

> Checks state of each function in a pipeline and exits block on error

## Description

`trypipe` checks the state of each function and exits the block if any of them
fail. Where `trypipe` differs from regular `try` blocks is `trypipe` will check
every process along the pipeline as well as the terminating function (which
`try` only validates against). The downside to this is that piped functions can
no longer run in parallel.

## Usage

    trypipe { code-block } -> <stdout>
    
    <stdin> -> trypipe { -> code-block } -> <stdout>

## Examples

    trypipe {
        out: "Hello, World!" -> grep: "non-existent string" -> cat
        out: "This command will be ignored"
    }
    
Formated pager (`less`) where the pager isn't called if the formatter (`pretty`) fails (eg input isn't valid JSON):

    func pless {
        -> trypipe { -> pretty -> less }
    }

## Detail

A failure is determined by:

* Any process that returns a non-zero exit number
* Any process that returns more output via STDERR than it does via STDOUT

You can see which run mode your functions are executing under via the `fid-list`
command.

## See Also

* [user-guide/Schedulers](../user-guide/schedulers.md):
  Overview of the different schedulers (or 'run modes') in _murex_
* [commands/`catch`](../commands/catch.md):
  Handles the exception code raised by `try` or `trypipe` 
* [commands/`fid-list`](../commands/fid-list.md):
  Lists all running functions within the current _murex_ session
* [commands/`if`](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [commands/`switch`](../commands/switch.md):
  Blocks of cascading conditionals
* [commands/`try`](../commands/try.md):
  Handles errors inside a block of code
* [commands/evil](../commands/evil.md):
  