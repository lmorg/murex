# `err`

> Print a line to the STDERR

## Description

Write parameters to STDERR with a trailing new line character.

## Usage

```
err string to write -> <stderr>
```

## Examples

```
» err Hello, World!
Hello, World!
```

## Detail

`err` outputs as `string` data-type. This can be changed by casting

```
err { "Code": 404, "Message": "Page not found" } ? cast json
```

However passing structured data-types along the STDERR stream is not recommended
as any other function within your code might also pass error messages along the
same stream and thus taint your structured data. This is why Murex does not
supply a `tout` function for STDERR. The recommended solution for passing
messages like these which you want separate from your STDOUT stream is to create
a new Murex named pipe.

```
» pipe --create messages
» bg { <messages> -> pretty }
» tout <messages> json { "Code": 404, "Message": "Page not found" }
» pipe --close messages
{
    "Code": 404,
    "Message": "Page not found"
}
```

### ANSI Constants

`err` supports ANSI constants.

## See Also

* [ANSI Constants](../user-guide/ansi.md):
  Infixed constants that return ANSI escape sequences
* [`(brace quote)`](../parser/brace-quote-func.md):
  Write a string to the STDOUT without new line (deprecated)
* [`<pipe>` Read Named Pipe](../commands/namedpipe.md):
  Reads from a Murex named pipe
* [`>>` Append File](../parser/greater-than-greater-than.md):
  Writes STDIN to disk - appending contents if file already exists
* [`bg`](../commands/bg.md):
  Run processes in the background
* [`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [`pipe`](../commands/pipe.md):
  Manage Murex named pipes
* [`pretty`](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [`pt`](../commands/pt.md):
  Pipe telemetry. Writes data-types and bytes written
* [`tout`](../commands/tout.md):
  Print a string to the STDOUT and set it's data-type
* [`|>` Truncate File](../parser/greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists

<hr/>

This document was generated from [builtins/core/io/echo_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/io/echo_doc.yaml).