# _murex_ Shell Docs

## Command Reference: `<stdin>` 

> Read the STDIN belonging to the parent code block

## Description

This is used inside functions and other code blocks to pass that block's
STDIN down a pipeline

## Usage

    <stdin> -> <stdout>

## Examples

When writing more complex scripts, you cannot always invoke your read as the
first command in a code block. For example a simple pipeline might be:

    » function: example { -> match: 2 }
    
But this only works if `->` is the very first command. The following would
fail:

    # Incorrect code
    function: example {
        out: "only match 2"
        -> match 2
    }
    
This is where `<stdin>` comes to our rescue:

    function: example {
        out: "only match 2"
        <stdin> -> match 2
    }
    
This could also be written as:

    function: example { out: "only match 2"; <stdin> -> match 2 }

## Synonyms

* `<stdin>`


## See Also

* [commands/`<>` (read pipe)](../commands/readpipe.md):
  Reads from a _murex_ named pipe
* [commands/pipe](../commands/pipe.md):
  