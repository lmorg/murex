# _murex_ Shell Docs

## Command Reference: `<>` (read pipe)

> Reads from a _murex_ named pipe

## Description

Sometimes you will need to start a commandline with a _murex_ named pipe:

    » <readpipe> -> match: foobar
    
> See the documentation on `pipe` for more details about _murex_ named pipes.

## Usage

    <example> -> <stdout>

## Examples

The follow two examples function the same

    » pipe: example
    » bg { <example> -> match: 2 }
    » a: <example> [1..3]
    2
    » !pipe: example

## Synonyms

* `<>`


## See Also

* [commands/`<stdin>` ](../commands/stdin.md):
  Read the STDIN belonging to the parent code block
* [commands/`pipe`](../commands/pipe.md):
  Manage _murex_ named pipes