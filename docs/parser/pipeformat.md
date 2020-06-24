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

* [parser/Arrow Pipe (`->`) Token](../parser/pipearrow.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [parser/POSIX Pipe (`|`) Token](../parser/pipeposix.md):
  Pipes STDOUT from the left hand command to STDIN of the right hand command
* [parser/STDERR Pipe (`?`) Token](../parser/pipeerr.md):
  Pipes STDERR from the left hand command to STDIN of the right hand command
* [commands/`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [commands/`ja`](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [commands/cat](../commands/cat.md):
  
* [parser/pipenamed](../parser/pipenamed.md):
  