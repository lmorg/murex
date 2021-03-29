# _murex_ Shell Docs

## Parser Reference: STDERR Pipe (`?`) Token

> Pipes STDERR from the left hand command to STDIN of the right hand command

## Description

This token swaps the STDOUT and STDERR streams of the left hand command.

Please note that this token is only effective when it is prefixed by white
space. 

## Examples

    » err Hello, world! ? regexp s/world/Earth/
    Hello, Earth!
    
In following example the first command is writing to STDOUT rather than STDERR
so `Hello, world!` doesn't get pipelined and thus isn't affected by `regexp`:

    » out Hello, world! ? regexp s/world/Earth/
    Hello, world!
    
In following example the STDERR token isn't whitespace padded so is treated
like any ordinary printable character:

    » err Hello, world!? regexp s/world/Earth/
    Hello, world!? regexp s/world/Earth/

## See Also

* [parser/Arrow Pipe (`->`) Token](../parser/pipe-arrow.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [parser/Formatted Pipe (`=>`) Token](../parser/pipe-format.md):
  Pipes a reformatted STDOUT stream from the left hand command to STDIN of the right hand command
* [parser/POSIX Pipe (`|`) Token](../parser/pipe-posix.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [parser/Pipeline](../parser/pipeline.md):
  Overview of what a "pipeline" is
* [commands/`err`](../commands/err.md):
  Print a line to the STDERR
* [commands/`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [commands/`regexp`](../commands/regexp.md):
  Regexp tools for arrays / lists of strings
* [parser/pipe-named](../parser/pipe-named.md):
  