# _murex_ Shell Docs

## Parser Reference: Formatted Pipe (`=>`) Token

> Pipes a reformatted STDOUT stream from the left hand command to STDIN of the right hand command

## Description

This token behaves much like the `->` pipe would except it injects `format
generic` into the pipeline. The purpose of a formatted pipe is to support
piping out to external commands which don't support _murex_ data types. For
example they might expect arrays as lists rather than JSON objects).

## Examples

    » ja: [Mon..Wed] => cat
    Mon
    Tue
    Wed
    
The above is literally the same as typing:

    » ja: [Mon..Wed] -> format generic -> cat
    Mon
    Tue
    Wed
    
To demonstrate how the previous pipeline might look without a formatted pipe:

    » ja: [Mon..Wed] -> cat
    ["Mon","Tue","Wed"]
    
    » ja: [Mon..Wed] | cat
    ["Mon","Tue","Wed"]
    
    » ja: [Mon..Wed]
    [
        "Mon",
        "Tue",
        "Wed"
    ]

## See Also

* [parser/Arrow Pipe (`->`) Token](../parser/pipe-arrow.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [parser/POSIX Pipe (`|`) Token](../parser/pipe-posix.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [parser/Pipeline](../parser/pipeline.md):
  Overview of what a "pipeline" is
* [parser/STDERR Pipe (`?`) Token](../parser/pipe-err.md):
  Pipes STDERR from the left hand command to STDIN of the right hand command
* [commands/`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [commands/`ja`](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [commands/cat](../commands/cat.md):
  
* [parser/pipe-named](../parser/pipe-named.md):
  