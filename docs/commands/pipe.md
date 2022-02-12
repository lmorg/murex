# _murex_ Shell Docs

## Command Reference: `pipe`

> Manage _murex_ named pipes

## Description

`pipe` creates and destroys _murex_ named pipes.

## Usage

Create pipe

    pipe: name [ pipe-type ]
    
Destroy pipe

    !pipe: name

## Examples

Create a standard pipe:

    pipe: example
    
Delete a pipe:

    !pipe: example
    
Create a TCP pipe (deleting a pipe is the same regardless of the type of pipe):

    pipe example --tcp-dial google.com:80
    bg { <example> }
    out: "GET /" -> <example>

## Detail

### What are _murex_ named pipes?

In POSIX, there is a concept of STDIN, STDOUT and STDERR, these are FIFO files
while are "piped" from one executable to another. ie STDOUT for application 'A'
would be the same file as STDIN for application 'B' when A is piped to B:
`A | B`. _murex_ adds a another layer around this to enable support for passing
data types and builtins which are agnostic to the data serialization format
traversing the pipeline. While this does add overhead the advantage is this new
wrapper can be used as a primitive for channelling any data from one point to
another.

_murex_ named pipes are where these pipes are created in a global store,
decoupled from any executing functions, named and can then be used to pass
data along asynchronously.

For example

    pipe: example
    
    bg {
        <example> -> match: Hello
    }
    
    out: "foobar"        -> <example>
    out: "Hello, world!" -> <example>
    out: "foobar"        -> <example>
    
    !pipe: example
    
This returns `Hello, world!` because `out` is writing to the **example** named
pipe and `match` is also reading from it in the background (`bg`).

Named pipes can also be inlined into the command parameters with `<>` tags

    pipe: example
    
    bg {
        <example> -> match: Hello
    }
    
    out: <example> "foobar"
    out: <example> "Hello, world!"
    out: <example> "foobar"
    
    !pipe: example
    
> Please note this is also how `test` works.

_murex_ named pipes can also represent network sockets, files on a disk or any
other read and/or write endpoint. Custom builtins can also be written in Golang
to support different abstractions so your _murex_ code can work with those read
or write endpoints transparently.

To see the different supported types run

    runtime --pipes
    
### Namespacing and used in modules and packages

Pipes created via `pipe` are created in the global namespace. This allows pipes
to be used across different functions easily however it does pose a risk with
name clashes where _murex_ named pipes are used heavily. Thus is it recommended
that pipes created in modules should be prefixed with the name of its package.

## Synonyms

* `pipe`
* `!pipe`


## See Also

* [user-guide/Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [commands/`<>` (murex named pipe)](../commands/namedpipe.md):
  Reads from a _murex_ named pipe
* [commands/`<stdin>` ](../commands/stdin.md):
  Read the STDIN belonging to the parent code block
* [commands/`bg`](../commands/bg.md):
  Run processes in the background
* [commands/`match`](../commands/match.md):
  Match an exact value in an array
* [commands/`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [commands/`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [commands/`test`](../commands/test.md):
  _murex_'s test framework - define tests, run tests and debug shell scripts
* [parser/namedpipe](../parser/namedpipe.md):
  