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
* [ansi](../commands/ansi.md):
  
* [brace-quote](../commands/brace-quote.md):
  
* [greater-than](../commands/greater-than.md):
  
* [greater-than-greater-than](../commands/greater-than-greater-than.md):
  
* [namedpipe](../commands/namedpipe.md):
  

<hr/>

This document was generated from [builtins/core/io/echo_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/io/echo_doc.yaml).