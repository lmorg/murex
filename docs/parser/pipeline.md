# _murex_ Shell Docs

## Parser Reference: Pipeline

> Overview of what a "pipeline" is

## Description

In the _murex_ docs you'll over hear the term "pipeline". This refers to any
commands sequenced together.

A pipeline can be joined via any pipe token (eg `|`, `->`, `=>`, `?`). But,
for the sake of documentation, a pipeline might even be a solitary command.

## Examples

Typical _murex_ pipeline:

    open: example.json -> [[ /node/0 ]]
    
Example of a single command pipeline:

    top
    
Pipeline you might see in Bash / Zsh (this is also valid in _murex_):

    cat names.txt | sort | uniq
    
Pipeline filtering out a specific error from `example-cmd`

    example-cmd ? grep: "File not found"

## See Also

* [parser/Arrow Pipe (`->`) Token](../parser/pipe-arrow.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [parser/Formatted Pipe (`=>`) Token](../parser/pipe-format.md):
  Pipes a reformatted STDOUT stream from the left hand command to STDIN of the right hand command
* [parser/POSIX Pipe (`|`) Token](../parser/pipe-posix.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [parser/STDERR Pipe (`?`) Token](../parser/pipe-err.md):
  Pipes STDERR from the left hand command to STDIN of the right hand command
* [parser/Schedulers](../parser/schedulers.md):
  Overview of the different schedulers (or run modes) in _murex_