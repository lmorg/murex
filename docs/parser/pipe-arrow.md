# _murex_ Shell Docs

## Parser Reference: Arrow Pipe (`->`) Token

> Pipes STDOUT from the left hand command to STDIN of the right hand command

## Description

This token behaves much like pipe would in Bash or similar shells. It passes
STDOUT along the pipeline while merging STDERR stream with the parents STDERR
stream.

It is identical in purpose to the `|` pipe token.

## Examples

    » out: Hello, world! -> regexp: s/world/Earth/
    Hello, Earth!
    
    » out: Hello, world!->regexp: s/world/Earth/
    Hello, Earth!
    
In following example the first command is writing to STDERR rather than STDOUT
so `Hello, world!` doesn't get pipelined and thus isn't affected by `regexp`:

    » err: Hello, world! -> regexp: s/world/Earth/
    Hello, world!

## See Also

* [parser/Formatted Pipe (`=>`) Token](../parser/pipe-format.md):
  Pipes a reformatted STDOUT stream from the left hand command to STDIN of the right hand command
* [parser/POSIX Pipe (`|`) Token](../parser/pipe-posix.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [user-guide/Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [parser/STDERR Pipe (`?`) Token](../parser/pipe-err.md):
  Pipes STDERR from the left hand command to STDIN of the right hand command
* [commands/`err`](../commands/err.md):
  Print a line to the STDERR
* [commands/`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [commands/`regexp`](../commands/regexp.md):
  Regexp tools for arrays / lists of strings
* [parser/pipe-named](../parser/pipe-named.md):
  